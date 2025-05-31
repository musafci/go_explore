package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	// "strconv"

	"github.com/xuri/excelize/v2"
)

// SplitCSV function to split a CSV file
func SplitCSV(inputFile string, rowsPerFile int) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("empty file")
	}

	header := records[0]
	data := records[1:]

	return splitAndWriteCSV(header, data, rowsPerFile, "file")
}

// SplitXLSX function to split an XLSX file
func SplitXLSX(inputFile string, rowsPerFile int) error {
	f, err := excelize.OpenFile(inputFile)
	if err != nil {
		return err
	}
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}

	if len(rows) == 0 {
		return fmt.Errorf("empty file")
	}

	header := rows[0]
	data := rows[1:]

	return splitAndWriteCSV(header, data, rowsPerFile, "file")
}

// splitAndWriteCSV function to split data and write to CSV files
func splitAndWriteCSV(header []string, data [][]string, rowsPerFile int, prefix string) error {
	totalFiles := (len(data) + rowsPerFile - 1) / rowsPerFile

	for i := 0; i < totalFiles; i++ {
		start := i * rowsPerFile
		end := start + rowsPerFile
		if end > len(data) {
			end = len(data)
		}

		outputFile := fmt.Sprintf("%s%d.csv", prefix, i+1)
		file, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write(header)
		writer.WriteAll(data[start:end])

		fmt.Println("Created:", outputFile)
	}

	return nil
}

func main() {
	inputFile := "./teest.csv" // Change to "input.xlsx" for XLSX processing
	rowsPerFile := 100         // Set the desired number of rows per file

	var err error
	if inputFile[len(inputFile)-4:] == ".csv" {
		err = SplitCSV(inputFile, rowsPerFile)
	} else if inputFile[len(inputFile)-5:] == ".xlsx" {
		err = SplitXLSX(inputFile, rowsPerFile)
	} else {
		log.Fatal("Unsupported file format")
	}

	if err != nil {
		log.Fatal(err)
	}
}
