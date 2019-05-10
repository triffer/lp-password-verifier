package internal

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func LoadPasswordRecords() []PasswordRecord {
	csvFile, err := os.Open(getCsvFilePath())
	CheckError(err)

	reader := csv.NewReader(csvFile)
	var records []PasswordRecord
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			CheckError(err)
		}

		records = append(records, PasswordRecord{
			Password: line[2],
			Name:     line[4],
		})
	}
	return records
}

func getCsvFilePath() string {
	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatalln("Missing CSV file path argument. Please provide a valid path e.g. /home/user/export.csv")
	}

	return os.Args[1]
}

type PasswordRecord struct {
	Name, Password string
}
