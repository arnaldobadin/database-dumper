package recorder

import (
	"os"
	"bufio"
)

type pipe struct {
	Output *Output
	
	file *os.File
	writer *bufio.Writer
}

func newPipe(output *Output) *pipe {
	path := output.Path + "/" + output.Name + "." + output.Ext
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	return &pipe{output, file, writer}
}

func (c *pipe) Write(data []string) error {
	var output string

	for i, val := range data {
		output += val
		if i != (len(data) - 1) {
			output += "|"
		} else {
			output += "\n"
		}
	}

	c.writer.Write([]byte(output))
	return nil
}

func (c *pipe) Flush() error {
	c.writer.Flush()
	return nil
}

func (c *pipe) Close() error {
	err := c.file.Close()
	if err != nil {
		return err
	}
	return nil
}