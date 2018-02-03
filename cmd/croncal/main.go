package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/vcraescu/croncal/internal/app/croncal/api"
	"github.com/vcraescu/croncal/internal/app/croncal/crontab"
	"gopkg.in/urfave/cli.v1"
)

const cronTabJSONFilename = "crontab.json"

func initStatics(e *echo.Echo) {
	fs := http.FileServer(http.Dir("./web/public"))

	e.GET("/", echo.WrapHandler(fs))
	e.GET("/compiled/*", echo.WrapHandler(fs))
}

func newCLIApp() *cli.App {
	app := cli.NewApp()
	app.Name = "CronCal"
	app.Usage = "Linux crontab calendar"
	app.Author = "Viorel Craescu <viorel@craescu.com>"
	app.Version = "0.1"
	app.UsageText = "croncal <options> [crontab file]"
	app.ArgsUsage = "[crontab file]"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind",
			Value: "0.0.0.0:8080",
			Usage: "Binding address",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Debug mode",
		},
	}

	return app
}

func main() {
	app := newCLIApp()

	app.Action = func(c *cli.Context) error {
		if c.NArg() <= 0 {
			return errors.New("crontab path is missing")
		}

		if c.Bool("debug") {
			log.Warnln("Debug Mode: ON")
			log.SetLevel(log.DebugLevel)
		}

		filename := c.Args().First()
		tab, err := crontab.FromCronTabFile(filename)
		if err != nil {
			return err
		}

		if _, err := os.Stat(cronTabJSONFilename); os.IsNotExist(err) {
			err = tab.ExportToJSON(cronTabJSONFilename)
			if err != nil {
				return err
			}
		}

		e := echo.New()
		initStatics(e)

		ctx := api.Context{}
		ctx.Prefix = "/api/v1"
		ctx.CronTabJSONFilename = cronTabJSONFilename

		return api.Start(e, ctx, c.String("bind"))
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
