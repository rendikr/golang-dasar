package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"john", "doe", "jane"})
	_ = writer.Write([]string{"peter", "steven"})
	_ = writer.Write([]string{"lorem", "ipsum"})

	writer.Flush()
}
