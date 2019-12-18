package main

import (
	"fmt"
	"flag"
	"github.com/arnaldobadin/dumper/pkg/config"
	"github.com/arnaldobadin/dumper/pkg/acquirer"
)

const EXCS_DEFAULT_PATH = "executions.json"

func main() {
	var excsFlagPath string
	flag.StringVar(&excsFlagPath, "excs", EXCS_DEFAULT_PATH, "Put the executions configuration (json file) here")
	flag.Parse()

	excs := []acquirer.Execution{}
	err := config.Load(excsFlagPath, &excs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, exc := range excs {
		err := exc.Execute()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Execution '" + exc.Name + "' dumped with success")
	}

	return
}