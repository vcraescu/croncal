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
	"github.com/gobuffalo/packr"
	"fmt"
	"github.com/CrowdSurge/banner"
	"github.com/labstack/gommon/color"
)

func initStatics(e *echo.Echo) {
	box := packr.NewBox("../../web/public")
	fs := http.FileServer(box)

	e.GET("/*", echo.WrapHandler(fs))
}

func newCLIApp() *cli.App {
	app := cli.NewApp()
	app.Name = "CronCal"
	app.Usage = "Linux crontab calendar"
	app.Author = "Viorel Craescu <viorel@craescu.com>"
	app.Version = "1.0"
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
			log.SetLevel(log.DebugLevel)
		}

		filename := c.Args().First()
		tab, err := crontab.FromCronTabFile(filename)
		if err != nil {
			return err
		}

		if tab.Empty() {
			return errors.New("crontab is empty")
		}

		jsonFilename := crontabJSONFilename(filename)
		err = tab.ExportToJSON(jsonFilename, true)
		if err != nil {
			return err
		}

		e := echo.New()
		e.HideBanner = true
		e.Debug = c.Bool("debug")
		initStatics(e)

		app := api.New(api.Config{
			Prefix:	"/api/v1",
			CronTabJSONFilename: jsonFilename,
			CronTabFilename: filename,
			Address: c.String("bind"),
			Debug: c.Bool("debug"),
			Logger: log.New(),
		})

		printBanner()

		return app.Start(e)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
	}
}

func printBanner() {
	color.Println(color.Blue(banner.PrintS("croncal")))
	fmt.Println()
}

func crontabJSONFilename(filename string) string {
	return fmt.Sprintf("%s.json", filename)
}

