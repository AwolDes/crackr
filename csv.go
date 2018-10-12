package main

import (
	"encoding/csv"
	"os"
	"sync"
)

// A useful resource
// https://markhneedham.com/blog/2017/01/31/go-multi-threaded-writing-csv-file/
type ThreadSafeCsvWriter struct {
	mutex  *sync.Mutex
	writer *csv.Writer
	file   *os.File
}

func newThreadSafeCsvWriter() ThreadSafeCsvWriter {
	f, err := os.Create(resultsFile + ".csv")
	checkError("Couldn't save changes: ", err)
	w := csv.NewWriter(f)
	return ThreadSafeCsvWriter{
		writer: w,
		mutex:  &sync.Mutex{},
		file:   f,
	}
}

func (w *ThreadSafeCsvWriter) writeChanges(data []string) {
	w.mutex.Lock()
	w.writer.Write(data)
	w.mutex.Unlock()
}

func (w *ThreadSafeCsvWriter) flush() {
	w.mutex.Lock()
	w.writer.Flush()
	w.mutex.Unlock()
}
