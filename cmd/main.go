package main

import (
	"os"

	//"github.com/ast9501/nssmf/internal/service/management"
	//nssmf_service "github.com/ast9501/nssmf/pkg/service"
	"github.com/urfave/cli/v2"
)

var NSSMF = &nssmf_service.NSSMF{}

func main() {
	app := cli.NewApp()
	app.Name = "nssmf"
	app.Usage = "3GPP NSSMF function for O-RAN"
	app.Action = action
	//app.Flags = NSSF.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		logger.AppLog.Errorf("NSSMF Run Error: %v\n", err)
	}
}

func action(c *cli.Context) error {
	NSSF.Start()

	return nil
}
