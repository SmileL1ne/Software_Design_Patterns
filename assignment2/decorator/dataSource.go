package main

type IDataSource interface {
	writeData(data string)
	readData() string
}
