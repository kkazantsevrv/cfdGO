package utils

import (
	"fmt"
	"math"

	"test.com/mat/matrix"
	"test.com/solvers/amgcl"
)

func exact_solution(p Point) float64 {
	x := p.X
	y := p.Y
	return math.Cos(10.0*x*x)*math.Sin(10.0*y) + math.Sin(10.0*x*x)*math.Cos(10.0*x)
}

func exact_rhs(p Point) float64 {
	x := p.X
	y := p.Y
	return (20.0*math.Sin(10.0*x*x)+(400.0*x*x+100.0)*math.Cos(10.0*x*x))*math.Sin(10.0*y) + (400.0*x*x+100.0)*math.Cos(10.0*x)*math.Sin(10.0*x*x) + (400.0*x*math.Sin(10*x)-20.0*math.Cos(10.0*x))*math.Cos(10.0*x*x)
}

func exact_dudx(p Point) float64 {
	x := p.X
	y := p.Y
	return 20.0*x*math.Cos(10.0*x)*math.Cos(10.0*x*x) - 10.0*math.Sin(10.0*x*x)*(math.Sin(10.0*x)+2.0*x*math.Sin(10.0*y))
}

func exact_dudy(p Point) float64 {
	x := p.X
	y := p.Y
	return 10.0 * math.Cos(10.0*x*x) * math.Cos(10.0*y)
}

func exact_dudn(p Point) float64 {
	x := p.X
	y := p.Y
	if math.Abs(x) < 1e-6 {
		return -exact_dudx(p)
	} else if math.Abs(x-1.0) < 1e-6 {
		return exact_dudx(p)
	}
	if math.Abs(y) < 1e-6 {
		return -exact_dudy(p)
	} else if math.Abs(y-1.0) < 1e-6 {
		return exact_dudy(p)
	}
	return 0.0
}

type Solver struct {
	x    []float64
	grid VTKGrid
	bnd  int
}

func (solver *Solver) Set_bnd_type(bnd int) error {
	if bnd >= 1 && bnd <= 3 {
		solver.bnd = bnd
	} else {
		return fmt.Errorf("this type is not exist")
	}
	return nil
}

func (solver *Solver) Set_grid(grid VTKGrid) {
	solver.grid = grid
	grid.Need_cell_centers()
}

func find_normal(pi Point, pj Point) Point {
	x := pj.Y - pi.Y
	y := -(pj.X - pi.X)
	return Point{x / math.Sqrt(x*x+y*y), y / math.Sqrt(x*x+y*y), 0.0}
}

func set_unit_row(m *matrix.DOKMatrix, row int) {
	cols := m.Cols
	for i := 0; i < cols; i++ {
		m.Set(row, i, 0.0)
	}
	m.Set(row, row, 1.0)
}

func (solver *Solver) Approximate_parts() {
	nn := len(solver.grid.Cells)
	// lhs := sparse.NewDOK(nn, nn)
	lhs, _ := matrix.NewDOKMatrix(nn, nn)
	rhs0 := make([]float64, nn)
	// fmt.Println(solver.grid.Faces_in_cel)
	for i := 0; i < len(solver.grid.Faces_in_cel); i++ {
		face := solver.grid.Faces_in_cel[i]
		left := face[1].Left
		right := face[1].Right
		pi := solver.grid.Points[face[0].Left]
		pj := solver.grid.Points[face[0].Right]
		normal := find_normal(pi, pj)
		ci := solver.grid.Cell_centers[left]
		cj := solver.grid.Cell_centers[right]
		hij := math.Abs((cj.X-ci.X)*normal.X + (cj.Y-ci.Y)*normal.Y + (cj.Z-ci.Z)*normal.Z)
		gij := math.Sqrt((pi.X-pj.X)*(pi.X-pj.X) + (pi.Y-pj.Y)*(pi.Y-pj.Y) + (pi.Z-pj.Z)*(pi.Z-pj.Z))
		v := gij / hij

		lhs.Set(left, left, lhs.Get(left, left)+v)
		lhs.Set(right, right, lhs.Get(right, right)+v)
		lhs.Set(left, right, lhs.Get(left, right)-v)
		lhs.Set(right, left, lhs.Get(right, left)-v)

	}
	set_unit_row(lhs, 0)
	for i := 0; i < len(solver.grid.Cells); i++ {
		rhs0[i] = exact_rhs(solver.grid.Cell_centers[i]) * solver.grid.Cell_volumes[i]
	}
	for i := 0; i < len(solver.grid.Faces_bnd_cel); i++ {
		face := solver.grid.Faces_bnd_cel[i]
		left := face[1].Left
		pi := solver.grid.Points[face[0].Left]
		pj := solver.grid.Points[face[0].Right]
		gij := math.Sqrt((pi.X-pj.X)*(pi.X-pj.X) + (pi.Y-pj.Y)*(pi.Y-pj.Y) + (pi.Z-pj.Z)*(pi.Z-pj.Z))
		center := Point{X: pi.X/2.0 + pj.X/2.0, Y: pi.Y/2.0 + pj.Y/2.0, Z: pi.Z/2.0 + pj.Z/2.0}
		rhs0[left] += gij * exact_dudn(center)
	}
	rhs0[0] = exact_solution(solver.grid.Cell_centers[0])

	csr, _ := lhs.ToCSR()
	slv, err := amgcl.NewSolver(csr)
	if err != nil {
		fmt.Println("error", err)
	}
	solver.x, err = slv.Solve(rhs0)
	if err != nil {
		fmt.Println("error", err)
	}
	slv.Free()
}

func (solver *Solver) Write_to_file(filename string) error {
	var cell_data []float64
	for i := 0; i < len(solver.grid.Cells); i++ {
		cell_data = append(cell_data, solver.x[i])
	}
	writefile(filename, cell_data)
	err := writefile(filename, cell_data)
	if err != nil {
		return err
	}
	return nil
}
