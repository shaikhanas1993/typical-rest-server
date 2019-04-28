package appcli

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var Version = "0.0.1"

func Run() {
	cliApp := cli.NewApp()
	cliApp.Version = Version
	cliApp.Commands = Commands()

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}