package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func writefile(filename string, cell_data []float64) error {
	// Определяем целевую директорию
	destinationDir := "output_directory"
	destinationFile := filepath.Join(destinationDir, filename)

	// Создаем целевую директорию, если она не существует
	err := os.MkdirAll(destinationDir, 0755) // 0755 - права доступа
	if err != nil {
		return fmt.Errorf("ошибка при создании директории: %v", err)
	}

	// Копируем файл из test_data в output_directory
	sourceFile := filepath.Join("test_data", filename)
	err = copyFile(sourceFile, destinationFile)
	if err != nil {
		return fmt.Errorf("ошибка при копировании файла: %v", err)
	}

	// Открываем файл для записи (перезаписываем, если файл уже существует)
	file, err := os.OpenFile(destinationFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer file.Close()

	// Записываем первую строку "cell_data"
	_, err = fmt.Fprintf(file, "CELL_DATA %d\nSCALARS numerical double 1\nLOOKUP_TABLE default\n", len(cell_data))
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}

	// Записываем каждое число из массива cell_data в отдельной строке
	for _, value := range cell_data {
		_, err = fmt.Fprintf(file, "%f\n", value)
		if err != nil {
			return fmt.Errorf("ошибка при записи в файл: %v", err)
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	// Открываем исходный файл
	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("не удалось открыть исходный файл: %v", err)
	}
	defer source.Close()

	// Создаем целевой файл
	destination, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("не удалось создать целевой файл: %v", err)
	}
	defer destination.Close()

	// Копируем данные из исходного файла в целевой
	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("ошибка при копировании данных: %v", err)
	}

	return nil
}
