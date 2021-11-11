package database

import (
	"bytes"
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var (
	db               database
	ComputerDatabase embed.FS
)

type database struct {
	reader *csv.Reader
}

func DB() database {
	if db == (database{}) {
		file, err := ComputerDatabase.ReadFile("assets/computers.csv")
		if err != nil {
			panic("")
		}
		db = database{
			reader: csv.NewReader(bytes.NewBuffer(file)),
		}
	}
	return db
}

func (database) GetComputerMacAddress(ipAddress string) string {
	for {
		record, err := db.reader.Read()
		if err != io.EOF {
			if strings.Contains(record[3], ipAddress) {
				return record[4]
			}
			continue
		}
		break
	}
	return ""
}

func (database) ToString() {
	for {
		record, err := db.reader.Read()
		if err != io.EOF {
			fmt.Println(record)
			continue
		}
		return
	}
}
