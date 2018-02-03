package api

import (
	"github.com/labstack/echo"
)

type application struct {
	ctx Context
}

// Context api application context
type Context struct {
	Prefix              string
	CronTabJSONFilename string
}

// Start the api
func Start(e *echo.Echo, ctx Context, address string) error {
	g := e.Group(ctx.Prefix)

	app := application{
		ctx: ctx,
	}

	es := &endpoints{
		app: app,
	}

	g.GET("/calendar/hourly", es.hourly)
	g.GET("/crons", es.crons)
	g.PUT("/crons/:id", es.updateCron)

	return e.Start(address)
}
