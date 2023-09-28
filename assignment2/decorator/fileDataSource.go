package main

import (
	"os"
)

type FileDataSource struct {
	dataSource IDataSource
	name       string
}

func (f *FileDataSource) writeData(dataString string) {
	file, errFile := os.Create(f.name)

	if errFile != nil {
		panic(errFile)
	}
	defer file.Close()

	_, errWrite := file.WriteString(dataString)
	if errWrite != nil {
		panic(errWrite)
	}
}

func (f *FileDataSource) readData() string {
	byteData, err := os.ReadFile(f.name)

	if err != nil {
		panic(err)
	}

	return string(byteData)
}
