package api

import (
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo"
	"github.com/vcraescu/croncal/internal/app/croncal/calendar"
	"github.com/vcraescu/croncal/internal/app/croncal/crontab"
	"gopkg.in/go-playground/validator.v9"
)

type endpoints struct {
	app application
}

type cronResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Cmd      string `json:"cmd"`
	Runtime  uint   `json:"runtime"`
	Position uint   `json:"position"`
	Interval string `json:"interval"`
}

type gridRowResponse struct {
	Label string   `json:"label"`
	Crons []string `json:"crons"`
}

type gridResponse struct {
	Rows []gridRowResponse `json:"rows"`
}

func (e endpoints) crons(ctx echo.Context) error {
	ct, err := crontab.FromJSONFile(e.app.ctx.CronTabJSONFilename)
	if err != nil {
		return err
	}

	crons := make([]crontab.Cron, len(ct.Crons))
	r := make([]cronResponse, len(ct.Crons))

	i := 0
	for _, c := range ct.Crons {
		crons[i] = *c
		i++
	}

	sort.Sort(crontab.CronsByPosition(crons))

	for i, c := range crons {
		r[i] = newCronResponse(c)
	}

	return ctx.JSON(http.StatusOK, r)
}

func (e endpoints) updateCron(ctx echo.Context) error {
	ct, err := crontab.FromJSONFile(e.app.ctx.CronTabJSONFilename)
	if err != nil {
		return err
	}

	id := ctx.Param("id")
	c, ok := ct.Crons[id]
	if !ok {
		return echo.ErrNotFound
	}

	if err = ctx.Bind(c); err != nil {
		return err
	}

	if err = ctx.Validate(c); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, newInvalidEntityResponse(err))
	}

	if c.Runtime <= 0 {
		c.Runtime = 0
	}

	err = ct.ExportToJSON(e.app.ctx.CronTabJSONFilename)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, newCronResponse(*c))
}

func (e endpoints) hourly(ctx echo.Context) error {
	ct, err := crontab.FromJSONFile(e.app.ctx.CronTabJSONFilename)
	if err != nil {
		return err
	}

	r := gridResponse{}

	cal := calendar.New(ct)

	view, err := cal.Hourly(time.Now())
	if err != nil {
		return err
	}

	labels := view.RowsLabels()
	r.Rows = make([]gridRowResponse, len(labels))

	for i, l := range labels {
		ids, _ := view.CronsIDsFromRow(l)
		r.Rows[i] = gridRowResponse{
			Label: l,
			Crons: ids,
		}
	}

	return ctx.JSON(http.StatusOK, r)
}

func newCronResponse(c crontab.Cron) cronResponse {
	return cronResponse{
		ID:       c.ID,
		Name:     c.Name,
		Cmd:      c.Cmd,
		Runtime:  c.Runtime,
		Position: c.Position,
		Interval: c.Interval,
	}
}

func newInvalidEntityResponse(err error) map[string][]string {
	errs := make(map[string][]string)

	vErrs, _ := err.(validator.ValidationErrors)
	for _, vErr := range vErrs {
		name := vErr.Field()
		if _, ok := errs[name]; !ok {
			errs[name] = make([]string, 0)
		}

		errs[name] = append(errs[name], vErr.Translate(validatorTranslator))
	}

	return errs
}
