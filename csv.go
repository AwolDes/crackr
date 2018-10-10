package main

import (
	"encoding/csv"
	"os"
)

func getRows(file string) [][]string {
	f, err := os.Open(file + ".csv")
	checkError("Cannot open file", err)
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	checkError("Cannot get rows file: ", err)
	return rows
}

func addNewRow(rows [][]string, newRow []string) [][]string {
	newRows := append(rows, newRow)
	return newRows
}

func writeChanges(file string, rows [][]string) {
	f, err := os.OpenFile(file+".csv", os.O_WRONLY, os.ModeAppend)
	checkError("Couldn't save changes: ", err)
	err = csv.NewWriter(f).WriteAll(rows)
	checkError("Could not write changes: ", err)
	f.Close()
}

func writeCSV(file string, data []string) {
	rows := getRows(file)
	rows = addNewRow(rows, data)
	writeChanges(file, rows)
}
