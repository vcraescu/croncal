package crontab

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/theckman/go-flock"
)

// Tab represents the entire cron tab file
type Tab struct {
	Crons map[string]*Cron
}

// New returns a new Tab which represents the crontab
func New() *Tab {
	return &Tab{
		Crons: make(map[string]*Cron),
	}
}

// FromJSONFile create a crontab from JSON file
func FromJSONFile(filename string) (*Tab, error) {
	content, err := safeReadFile(filename)
	if err != nil {
		return nil, err
	}

	crons := make([]Cron, 0)
	json.Unmarshal(content, &crons)

	tab := New()

	for _, c := range crons {
		tab.Add(c)
	}

	return tab, nil
}

// FromCronTabFile create a crontab from crontab file
func FromCronTabFile(filename string) (*Tab, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	t := New()
	r := bufio.NewReader(f)
	i := 0
	for {
		i++
		l, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		l = trimWhites(l)
		if l == "" {
			continue
		}

		interval, cmd, err := t.ParseLine(l)
		if err != nil {
			log.Errorln(err)
			continue
		}

		err = t.AddLine(i, interval, cmd)
		if err != nil {
			log.Errorln(err)
		}
	}

	return t, nil
}

// ExportToJSON exports cron tab content to json file
func (t Tab) ExportToJSON(filename string) error {
	crons := make([]Cron, len(t.Crons))

	i := 0
	for _, c := range t.Crons {
		crons[i] = *c
		i++
	}

	content, err := json.Marshal(crons)
	if err != nil {
		return err
	}

	return safeWriteFile(filename, content)
}

// Add adds a new Cron to Tab
func (t *Tab) Add(c Cron) {
	t.Crons[c.ID] = &c
}

// AddLine creates a new Cron from interval and command parts and adds it Tab
func (t *Tab) AddLine(position int, interval, cmd string) (error) {
	c := NewCron(interval, cmd)

	c.ID = uuid.NewV3(uuid.NamespaceOID, strconv.Itoa(position)).String()

	t.Add(*c)

	return nil
}

// ParseLine returns the interval and command parts from a linux crontab line
func (t Tab) ParseLine(l string) (interval string, cmd string, err error) {
	if l == "" {
		err = errors.New("empty line")
		return
	}

	li := stripComments(l)
	if li == "" {
		err = fmt.Errorf("commented line: '%s'", l)
		return
	}

	ts := strings.Split(li, " ")
	if len(ts) < 6 {
		err = fmt.Errorf("invalid line: '%s'", l)
		return
	}

	interval = trimWhites(strings.Join(ts[:5], " "))
	cmd = trimWhites(strings.Join(ts[5:], " "))

	return
}

func stripComments(line string) string {
	line = trimWhites(line)
	if line == "" || line[0] == '#' {
		return ""
	}

	i := strings.Index(line, "#")
	if i == -1 {
		return trimWhites(line)
	}

	return trimWhites(line[0:i])
}

func trimWhites(t string) string {
	return strings.Trim(t, " \n\t\r")
}

func waitForFileLock(filename string) (*flock.Flock, error) {
	fl := flock.NewFlock(fmt.Sprintf("%s.lock", filename))
	locked := false
	var err error

	for !locked {
		locked, err = fl.TryLock()
		if err != nil {
			return nil, err
		}

		time.Sleep(time.Millisecond * 10)
	}

	return fl, nil
}

func safeReadFile(filename string) ([]byte, error) {
	fl, err := waitForFileLock(filename)
	if err != nil {
		return nil, err
	}

	defer fl.Unlock()

	return ioutil.ReadFile(filename)
}

func safeWriteFile(filename string, data []byte) error {
	fl, err := waitForFileLock(filename)
	if err != nil {
		return err
	}

	defer fl.Unlock()

	return ioutil.WriteFile(filename, data, 0744)
}
