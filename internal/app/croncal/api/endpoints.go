package api

import (
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo"
	"github.com/vcraescu/croncal/internal/app/croncal/calendar"
	"github.com/vcraescu/croncal/internal/app/croncal/crontab"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"os"
)

type gridRowResponse struct {
	Label string   `json:"label"`
	Crons []string `json:"crons"`
}

type gridResponse struct {
	Rows []gridRowResponse `json:"rows"`
}

type unprocessableEntityResponse struct {
	Message string `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

func (app Application) crons(ctx echo.Context) error {
	ct, err := crontab.FromJSONFile(app.cfg.CronTabJSONFilename)
	if err != nil {
		return err
	}

	crons := ct.All()
	sort.Sort(crontab.CronsByPosition(crons))

	return ctx.JSON(http.StatusOK, crons)
}

func (app Application) updateCron(ctx echo.Context) error {
	ct, err := crontab.FromJSONFile(app.cfg.CronTabJSONFilename)
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
		return ctx.JSON(http.StatusUnprocessableEntity, newUnprocessableEntityResponse(err))
	}

	if c.Runtime <= 0 {
		c.Runtime = 0
	}

	err = ct.ExportToJSON(app.cfg.CronTabJSONFilename, false)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, c)
}

func (app Application) saveCrontab(ctx echo.Context) error {
	ct, err := crontab.FromJSONFile(app.cfg.CronTabJSONFilename)
	if err != nil {
		return err
	}

	err = ct.ExportToCrontab(app.cfg.CronTabFilename)
	if err != nil {
		return err
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]string{
			"message": "Crontab successfully saved",
		},
	)
}

func (app Application) downloadCrontab(ctx echo.Context) error {
	ct, err := crontab.FromJSONFile(app.cfg.CronTabJSONFilename)
	if err != nil {
		return err
	}

	f, err := ioutil.TempFile("/tmp", "crontab")
	if err != nil {
		return err
	}

	filename := f.Name()
	err = ct.ExportToCrontab(filename)
	if err != nil {
		return err
	}

	defer os.Remove(filename)

	return ctx.Attachment(filename, "crontab")
}

func (app Application) hourly(ctx echo.Context) error {
	ct, err := crontab.FromJSONFile(app.cfg.CronTabJSONFilename)
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

func newUnprocessableEntityResponse(err error) unprocessableEntityResponse {
	r := unprocessableEntityResponse{
		Message: "Please correct the errors",
		Errors: make(map[string][]string),
	}

	vErrs, _ := err.(validator.ValidationErrors)
	for _, vErr := range vErrs {
		name := vErr.Field()
		if _, ok := r.Errors[name]; !ok {
			r.Errors[name] = make([]string, 0)
		}

		r.Errors[name] = append(r.Errors[name], vErr.Translate(validatorTranslator))
	}

	return r
}
