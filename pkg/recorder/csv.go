package recorder

import (
	"os"
	"encoding/csv"
)

type csvr struct {
	Output *Output
	
	file *os.File
	writer *csv.Writer
}

func newCsvr(output *Output) *csvr {
	path := output.Path + "/" + output.Name + "." + output.Ext
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	return &csvr{output, file, writer}
}

func (c *csvr) Write(data []string) error {
	c.writer.Write(data)
	return nil
}

func (c *csvr) Flush() error {
	c.writer.Flush()
	return nil
}

func (c *csvr) Close() error {
	err := c.file.Close()
	if err != nil {
		return err
	}
	return nil
}