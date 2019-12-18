package acquirer

import (
	"github.com/arnaldobadin/dumper/pkg/storage"
	"github.com/arnaldobadin/dumper/pkg/recorder"
)

type Execution struct {
	Name string `json:"name"`
	Query string `json:"query"`
	Conn storage.Connection `json:"connection"`
	Output recorder.Output `json:"output"`
}

func (e *Execution) Execute() error {
	ctx := storage.NewConnector(&e.Conn)
	ctx.Open()
	defer ctx.Close()

	rec := recorder.NewRecorder(&e.Output)
	defer rec.Close()

	err := ctx.Query(e.Query, func(data []string) {
		rec.Write(data)
	})
	if err != nil {
		return err
	}

	err = rec.Flush()
	return err
}
