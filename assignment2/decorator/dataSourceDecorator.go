package main

type DataSourceDecorator struct {
	dataSource IDataSource
}

func (d *DataSourceDecorator) writeData(data string) {
	d.dataSource.writeData(data)
}

func (d *DataSourceDecorator) readData() string {
	return d.dataSource.readData()
}
