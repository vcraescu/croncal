package calendar

import (
	"fmt"
	"time"

	"github.com/vcraescu/croncal/internal/app/croncal/crontab"
)

// Calendar type
type Calendar struct {
	tab *crontab.Tab
}

// New returns a new calendar
func New(tab *crontab.Tab) *Calendar {
	return &Calendar{
		tab: tab,
	}
}

// Hourly returns all crons scheduled for every hour of the day
func (cal Calendar) Hourly(day time.Time) (Grid, error) {
	data := make(map[string]map[string]crontab.Cron)
	for i := 0; i <= 59; i++ {
		data[fmt.Sprintf("%02d", i)] = nil
	}

	for _, c := range cal.tab.Crons {
		ts, err := c.DaySchedule(day)
		if err != nil {
			return nil, err
		}

		for _, t := range ts {
			k := fmt.Sprintf("%02d", t.Minute())
			if data[k] == nil {
				data[k] = make(map[string]crontab.Cron)
			}

			data[k][c.ID] = *c
		}
	}

	return data, nil
}
