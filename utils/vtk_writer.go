package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func writefile(filename string, cell_data []float64) error {
	destinationDir := "output_directory" // Новая директория
	destinationFile := filepath.Join(destinationDir, filename)

	err := os.MkdirAll(destinationDir, 0755) // 0755 - права доступа
	if err != nil {
		fmt.Println("Ошибка при создании директории:", err)
		return err
	}

	err = copyFile(filename, destinationFile)
	if err != nil {
		fmt.Println("Ошибка при копировании файла:", err)
		return err
	}

	file, err := os.OpenFile(destinationFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer file.Close()

	// Записываем первую строку "cell_data"
	_, err = fmt.Fprintf(file, "CELL_DATA %d\nSCALARS numerical double 1\nLOOKUP_TABLE default\n", len(cell_data)) // Используем Fprintf для записи строки
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}

	// Записываем каждое число из массива cellData в отдельной строке
	for _, value := range cell_data {
		_, err = fmt.Fprintf(file, "%f\n", value) // Используем Fprintf для записи чисел
		if err != nil {
			return fmt.Errorf("ошибка при записи в файл: %v", err)
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
