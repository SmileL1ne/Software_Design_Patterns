package main

import (
	"encoding/base64"
	"fmt"
)

type EncryptionDecorator struct {
	dataSource IDataSource
}

func (e *EncryptionDecorator) writeData(data string) {
	e.dataSource.writeData(e.encode(data))
}

func (e *EncryptionDecorator) readData() string {
	return e.decode(e.dataSource.readData())
}

func (e *EncryptionDecorator) encode(stringData string) string {
	return base64.StdEncoding.EncodeToString([]byte(stringData))
}

func (e *EncryptionDecorator) decode(stringData string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(stringData)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(decodedBytes)
}
