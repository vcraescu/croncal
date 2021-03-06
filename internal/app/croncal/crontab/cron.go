package crontab

import (
	"time"

	"github.com/gorhill/cronexpr"
)

// Cron represents a cron from crontab
type Cron struct {
	ID       string `json:"id"`
	Interval string `json:"interval" validate:"required,min=9,interval"`
	Cmd      string `json:"cmd" validate:"required,min=4"`
	Name     string `json:"name" validate:"required,min=2"`
	Runtime  uint   `json:"runtime" validate:"required,min=0,max=59"`
	Position uint   `json:"position" validate:"min=0"`
}

// CronsByPosition type for sorting crons by Position
type CronsByPosition []Cron

func (c CronsByPosition) Len() int {
	return len(c)
}
func (c CronsByPosition) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c CronsByPosition) Less(i, j int) bool {
	return c[i].Position < c[j].Position
}

// NewCron return a new Cron
func NewCron(interval, cmd string) *Cron {
	return &Cron{
		Interval: interval,
		Cmd:      cmd,
		Name:     cmd,
		Runtime:  1,
	}
}

func (c Cron) String() string {
	return c.Name
}

// IntervalExpr returns expression from interval string
func (c Cron) IntervalExpr() (*cronexpr.Expression, error) {
	return cronexpr.Parse(c.Interval)
}

// DaySchedule returns schedule for a date
func (c Cron) DaySchedule(today time.Time) ([]time.Time, error) {
	t := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)

	expr, err := c.IntervalExpr()
	if err != nil {
		return nil, err
	}

	return expr.NextN(t, 1440), nil
}
