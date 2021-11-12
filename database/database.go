package database

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var computerListBytes []byte

type database struct {
	reader *csv.Reader
}

func SetComputersDatabase(computersListFile []byte) {
	computerListBytes = computersListFile
}

func DB() database {
	if computerListBytes == nil {
		panic("Computer list is haven't been set")
	}
	db := database{
		reader: csv.NewReader(bytes.NewBuffer(computerListBytes)),
	}
	return db
}

func (db database) GetComputerMacAddress(ipAddress string) string {
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

func (db database) ToString() {
	for {
		record, err := db.reader.Read()
		if err != io.EOF {
			fmt.Println(record)
			continue
		}
		return
	}
}
