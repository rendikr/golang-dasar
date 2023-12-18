package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "john,doe,jane\n" +
		"peter,steven\n" +
		"lorem,ipsum"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF { // EOF = End of File
			break
		}

		fmt.Println(record)
	}
}
