package calendar

import (
	"errors"
	"sort"

	"github.com/vcraescu/croncal/internal/app/croncal/crontab"
)

// Grid calendar view
type Grid map[string]map[string]crontab.Cron

// Crons returns all crons from the view
func (g Grid) Crons() []crontab.Cron {
	data := make(map[string]crontab.Cron)
	for _, crons := range g {
		for _, c := range crons {
			data[c.ID] = c
		}
	}

	crons := make([]crontab.Cron, len(data))

	i := 0
	for _, c := range data {
		crons[i] = c
		i++
	}

	sort.Sort(crontab.CronsByID(crons))

	return crons
}

// Headings return grid headings
func (g Grid) Headings() []string {
	var headings []string
	for _, cols := range g {
		headings = make([]string, len(cols))

		i := 0
		for h := range cols {
			headings[i] = h
			i++
		}
		break
	}

	return headings
}

// RowsLabels returns all row labels from view
func (g Grid) RowsLabels() []string {
	keys := make([]string, len(g))

	i := 0
	for k := range g {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}

// CronsFromRow returns all crons that have a slot occupied on the row
func (g Grid) CronsFromRow(label string) ([]crontab.Cron, error) {
	cs, ok := g[label]
	if !ok {
		return nil, errors.New("label does not exists")
	}

	crons := make([]crontab.Cron, len(cs))

	i := 0
	for _, c := range g[label] {
		crons[i] = c
		i++
	}

	return crons, nil
}

// CronsIDsFromRow returns all crons ids from row
func (g Grid) CronsIDsFromRow(label string) ([]string, error) {
	crons, err := g.CronsFromRow(label)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(crons))
	for i, c := range crons {
		ids[i] = c.ID
	}

	return ids, nil
}
