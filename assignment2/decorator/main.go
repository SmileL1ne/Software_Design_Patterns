package main

import "fmt"

func main() {
	sampleText := "I'm stored in some txt file here!"

	fileDataSource := &FileDataSource{
		name: "testFile.txt",
	}
	compressionDecorator := &CompressionDecorator{
		dataSource: fileDataSource,
	}
	encryptionDecorator := &EncryptionDecorator{
		dataSource: compressionDecorator,
	}
	encodedAndCompressed := &DataSourceDecorator{
		dataSource: encryptionDecorator,
	}

	encodedAndCompressed.writeData(sampleText)
	plain := &FileDataSource{
		name: "testFile.txt",
	}

	fmt.Println("Intial data: ", sampleText)
	fmt.Println("Encrypted data: ", plain.readData())
	fmt.Println("Data from compressed and encrypted file: ", encodedAndCompressed.readData())
}
