package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const csvDelimiter = ';'

func readCSVFile(filePath string, delimiter rune) ([][]string, error) {
	// Открываем файл CSV
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Создаем CSV reader
	reader := csv.NewReader(file)
	if delimiter == 0 {
		delimiter = csvDelimiter
	}
	reader.Comma = delimiter // Устанавливаем разделитель как ';'

	// Читаем все строки из CSV файла
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read lines from file '%s': %w", filePath, err)
	}

	return records, nil
}
