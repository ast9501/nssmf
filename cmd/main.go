package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ast9501/nssmf/docs"
	"github.com/ast9501/nssmf/pkg/logger"
	nssmf_service "github.com/ast9501/nssmf/pkg/service"
	"github.com/urfave/cli/v2"
)

var NSSMF = &nssmf_service.NSSMF{}

//	@title			O-RAN NSSMF api doc
//	@version		1.0
//	@description	winlab O-RAN NSSMF

//	@contact.name	ast9501
//	@contact.email	ast9501.cs10@nycu.edu.tw

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// schemes http
func main() {
	app := cli.NewApp()
	app.Name = "nssmf"
	app.Usage = "3GPP NSSMF function for O-RAN"
	app.Action = action
	app.Flags = NSSMF.GetCliCmd()

	if err := app.Run(os.Args); err != nil {
		// TODO: Add logger printer
		logger.InitLog.Errorf("NSSMF Run Error: %v\n", err)
	}
	// generate host ip dynamicly for api doc
	docs.SwaggerInfo.Host = NSSMF.Config.Addr + NSSMF.Config.Port
	logger.InitLog.Debugln("Generate swagger api doc target server location: ", docs.SwaggerInfo.Host)

}

func action(c *cli.Context) error {
	if c.String("c") == "" {
		fmt.Println("config is null!")
		return nil
	}
	// TODO: Add log: print config file path
	NSSMF.Initialize(c.String("c"))

	NSSMF.Start(NSSMF.Config.Cert, NSSMF.Config.Key)

	return nil
}

// generate server outbound IP for api doc
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
