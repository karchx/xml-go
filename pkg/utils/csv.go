package utils

import (
	"encoding/csv"
	"os"

	log "github.com/gothew/l-og"
)

type CSV struct {
	NumeroInterno string
}

func ReadCsvFile(filePath string) []CSV {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var data []CSV

	for _, line := range records {
		d := CSV{
			NumeroInterno: line[0],
		}

		data = append(data, d)
	}
	return data
}
