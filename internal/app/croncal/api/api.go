package api

import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/labstack/gommon/color"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

// Application struct
type Application struct {
	cfg Config
}

// Config application
type Config struct {
	Prefix              string
	CronTabJSONFilename string
	CronTabFilename     string
	Address             string
	Debug               bool
	Logger              *log.Logger
}

// New returns a new application object
func New(cfg Config) Application {
	if cfg.Logger == nil {
		logger := log.New()
		logger.Out = ioutil.Discard
		cfg.Logger = logger
	}

	app := Application{
		cfg: cfg,
	}

	return app
}

// Start the api
func (app Application) Start(e *echo.Echo) error {
	g := e.Group(app.cfg.Prefix)

	g.GET("/calendar/hourly", app.hourly)
	g.GET("/crons", app.crons)
	g.PATCH("/crons/:id", app.updateCron)
	g.PATCH("/crontab/save", app.saveCrontab)
	g.GET("/crontab/download", app.downloadCrontab)

	e.Validator = newEntityValidator()

	fmt.Printf("Started on ")
	fmt.Println(color.Bold(fmt.Sprintf("http://%s\n", app.cfg.Address)))

	if app.cfg.Debug {
		log.Warnln("Debug Mode: ON")
	}

	return e.Start(app.cfg.Address)
}
