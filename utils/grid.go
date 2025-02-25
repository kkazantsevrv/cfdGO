package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Point представляет точку в 3D пространстве
type Point struct {
	X, Y, Z float64
}

// Cell представляет ячейку, состоящую из индексов точек
type Cell struct {
	Indices []int
}

// Face представляет грань, по разные стороны от которых расположены ячейки с номерами left и right
type Face struct {
	Left, Right int
}

// VTKGrid представляет сетку в формате VTK
type VTKGrid struct {
	Title         string
	Format        string
	DatasetType   string
	Points        []Point
	Cells         []Cell
	CellTypes     []int // Добавлено поле для типов ячеек
	Faces         []Face
	Faces_bnd_cel [][2]Face
	Faces_in_cel  [][2]Face
	Cell_centers  []Point
	Cell_volumes  []float64
}

// Grid загружает VTK-файл
func Grid(filename string) (VTKGrid, error) {
	// Открываем файл
	file, err := os.Open(filename)
	if err != nil {
		return VTKGrid{}, fmt.Errorf("ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Создаем экземпляр VTKGrid
	grid := VTKGrid{}

	// Читаем файл построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Заголовок
		if strings.HasPrefix(line, "# vtk DataFile Version") {
			grid.Title = strings.TrimSpace(strings.TrimPrefix(line, "# vtk DataFile Version"))
		}

		// Формат
		if line == "ASCII" {
			grid.Format = "ASCII"
		}

		// Тип данных
		if strings.HasPrefix(line, "DATASET") {
			grid.DatasetType = strings.TrimSpace(strings.TrimPrefix(line, "DATASET"))
		}

		// Точки
		if strings.HasPrefix(line, "POINTS") {
			parts := strings.Fields(line)
			numPoints, err := strconv.Atoi(parts[1])
			if err != nil {
				return VTKGrid{}, fmt.Errorf("ошибка при парсинге количества точек: %v", err)
			}
			grid.Points = make([]Point, numPoints)

			for i := 0; i < numPoints; i++ {
				scanner.Scan()
				coords := strings.Fields(scanner.Text())
				if len(coords) < 3 {
					return VTKGrid{}, fmt.Errorf("недостаточно координат для точки")
				}
				x, err := strconv.ParseFloat(coords[0], 64)
				if err != nil {
					return VTKGrid{}, fmt.Errorf("ошибка при парсинге координаты X: %v", err)
				}
				y, err := strconv.ParseFloat(coords[1], 64)
				if err != nil {
					return VTKGrid{}, fmt.Errorf("ошибка при парсинге координаты Y: %v", err)
				}
				z, err := strconv.ParseFloat(coords[2], 64)
				if err != nil {
					return VTKGrid{}, fmt.Errorf("ошибка при парсинге координаты Z: %v", err)
				}
				grid.Points[i] = Point{X: x, Y: y, Z: z}
			}
		}

		// Ячейки
		if strings.HasPrefix(line, "CELLS") {
			parts := strings.Fields(line)
			numCells, err := strconv.Atoi(parts[1])
			if err != nil {
				return VTKGrid{}, fmt.Errorf("ошибка при парсинге количества ячеек: %v", err)
			}
			grid.Cells = make([]Cell, numCells)

			for i := 0; i < numCells; i++ {
				scanner.Scan()
				cellData := strings.Fields(scanner.Text())
				if len(cellData) == 0 {
					return VTKGrid{}, fmt.Errorf("пустая строка для ячейки")
				}
				numIndices, err := strconv.Atoi(cellData[0])
				if err != nil {
					return VTKGrid{}, fmt.Errorf("ошибка при парсинге количества индексов: %v", err)
				}
				indices := make([]int, numIndices)
				for j := 0; j < numIndices; j++ {
					indices[j], err = strconv.Atoi(cellData[j+1])
					if err != nil {
						return VTKGrid{}, fmt.Errorf("ошибка при парсинге индекса: %v", err)
					}
				}
				grid.Cells[i] = Cell{Indices: indices}
			}
		}

		// Типы ячеек
		if strings.HasPrefix(line, "CELL_TYPES") {
			parts := strings.Fields(line)
			numCellTypes, err := strconv.Atoi(parts[1])
			if err != nil {
				return VTKGrid{}, fmt.Errorf("ошибка при парсинге количества типов ячеек: %v", err)
			}
			grid.CellTypes = make([]int, numCellTypes)

			for i := 0; i < numCellTypes; i++ {
				scanner.Scan()
				cellType, err := strconv.Atoi(scanner.Text())
				if err != nil {
					return VTKGrid{}, fmt.Errorf("ошибка при парсинге типа ячейки: %v", err)
				}
				grid.CellTypes[i] = cellType
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return VTKGrid{}, fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	// === Добавляем грани ===

	// Карта для хранения граней и связанных с ними ячеек
	faceMap := make(map[[2]int][]int)
	// Перебираем все ячейки и их рёбра
	for cellIndex, cell := range grid.Cells {
		numIndices := len(cell.Indices)

		for i := 0; i < numIndices; i++ {
			// Грань = два соседних индекса (петляем обратно на первый)
			a, b := cell.Indices[i], cell.Indices[(i+1)%numIndices]

			// Сортируем индексы, чтобы избежать дублирования (a,b) и (b,a)
			if a > b {
				a, b = b, a
			}
			key := [2]int{a, b}
			faceMap[key] = append(faceMap[key], cellIndex)
		}
	}
	// Формируем список граней
	for points, cells := range faceMap {
		if len(cells) > 2 {
			return VTKGrid{}, fmt.Errorf("грань принадлежит более чем двум ячейкам: %v", cells)
		}
		left := cells[0]
		right := -1
		if len(cells) > 1 {
			right = cells[1] // Вторая ячейка (если есть)
			grid.Faces_in_cel = append(grid.Faces_in_cel, [2]Face{Face{Left: points[0], Right: points[1]}, Face{Left: left, Right: right}})
		} else {
			grid.Faces_bnd_cel = append(grid.Faces_bnd_cel, [2]Face{Face{Left: points[0], Right: points[1]}, Face{Left: left, Right: right}})
		}
		grid.Faces = append(grid.Faces, Face{Left: left, Right: right})
	}

	return grid, nil
}

func (grid *VTKGrid) Need_cell_centers() {
	grid.Cell_volumes = make([]float64, len(grid.Cells))
	grid.Cell_centers = make([]Point, len(grid.Cells))
	for i := 0; i < len(grid.Cells); i++ {
		sum := 0.0
		sum0 := 0.0
		center := Point{0.0, 0.0, 0.0}
		cel := grid.Cells[i]
		numIndices := len(grid.Cells[i].Indices)
		for j := 0; j < numIndices-1; j++ {
			x0 := grid.Points[cel.Indices[0]]
			xi := grid.Points[cel.Indices[j]]
			xip := grid.Points[cel.Indices[j+1]]
			sum = ((xi.X-x0.X)*(xip.Y-x0.Y) - (xi.Y-x0.Y)*(xip.X-x0.X)) / 2.0
			sum0 += sum
			center.X += (x0.X + xi.X + xip.X) * sum / 3.0
			center.Y += (x0.Y + xi.Y + xip.Y) * sum / 3.0
			center.Z += (x0.Z + xi.Z + xip.Z) * sum / 3.0
		}
		grid.Cell_volumes[i] = sum0
		grid.Cell_centers[i] = Point{
			X: center.X / sum0,
			Y: center.Y / sum0,
			Z: center.Z / sum0,
		}
	}
}
