package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"strings"
)

var stringBuilder strings.Builder

type CompressionDecorator struct {
	dataSource IDataSource
	compLevel  int
}

func (c *CompressionDecorator) writeData(data string) {
	c.dataSource.writeData(c.compress(data))
}

func (c *CompressionDecorator) readData() string {
	return c.decompress(c.dataSource.readData())
}

func (c *CompressionDecorator) compress(stringData string) string {
	var compressedBuffer bytes.Buffer
	gzipWriter, errWriter := gzip.NewWriterLevel(&compressedBuffer, c.compLevel)

	if errWriter != nil {
		panic(errWriter)
	}

	_, errWrite := gzipWriter.Write([]byte(stringData))

	if errWrite != nil {
		panic(errWrite)
	}

	if errClose := gzipWriter.Close(); errClose != nil {
		panic(errClose)
	}

	stringBuilder.Write(compressedBuffer.Bytes())

	return stringBuilder.String()
}

func (c *CompressionDecorator) decompress(stringData string) string {
	stringBuilder.Reset()

	compressedData := []byte(stringData)
	compressedReader := bytes.NewReader(compressedData)
	gzipReader, errReader := gzip.NewReader(compressedReader)

	if errReader != nil {
		panic(errReader)
	}

	decompressedData, errRead := io.ReadAll(gzipReader)

	if errRead != nil {
		panic(errRead)
	}

	_, errWrite := stringBuilder.Write(decompressedData)

	if errWrite != nil {
		panic(errWrite)
	}

	return stringBuilder.String()
}

func (c *CompressionDecorator) getCompLevel() int {
	return c.compLevel
}

func (c *CompressionDecorator) setCompLevel(compressionLevel int) {
	c.compLevel = compressionLevel
}
