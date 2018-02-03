package crontab

import (
	"time"

	"github.com/gorhill/cronexpr"
)

// Cron represents a cron from crontab
type Cron struct {
	ID       string `json:"id"`
	Interval string `json:"interval"`
	Cmd      string `json:"cmd"`
	Name     string `json:"name"`
	Runtime  uint   `json:"runtime"`
}

// CronsByID type for sorting crons by ID
type CronsByID []Cron

func (s CronsByID) Len() int {
	return len(s)
}
func (s CronsByID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CronsByID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
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
