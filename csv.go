package main

import (
	"encoding/csv"
	"os"
)

func writeChanges(file string, data []string) {
	f, err := os.OpenFile(file+".csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	checkError("Couldn't save changes: ", err)
	w := csv.NewWriter(f)
	w.Write(data)
	w.Flush()
	f.Close()
}

func writeCSV(file string, data []string) {
	writeChanges(file, data)
}
