package recorder

const CSV_EXTENSION = "csv"

type Recorder interface {
	Write(data []string) error
	Flush() error
	Close() error
}

func NewRecorder(output *Output) Recorder {
	switch output.Ext {
		case CSV_EXTENSION:
			return newCsvr(output)
		default:
			panic("Wrong extension type")
	}
}