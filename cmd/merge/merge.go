package merge

import (
	"flag"
	"strings"
	"time"

	"github.com/damirm/nurse-schedule/internal/command/merge"
	"github.com/damirm/nurse-schedule/internal/schedule"
)

func NewCommand() *Command {
	return &Command{}
}

type Config struct {
	paths     string
	employees string
	out       string
	dates     string
}

var config = Config{}

type Command struct {
}

func (f *Command) Name() string {
	return "merge"
}

func (f *Command) ExportFlags(flagSet *flag.FlagSet) error {
	flagSet.StringVar(&config.paths, "paths", "", "list of file paths: ./path/to/a.csv,./path/to/b.csv")
	flagSet.StringVar(&config.employees, "employees", "", "list of employees names: a, b")
	flagSet.StringVar(&config.dates, "dates", "", "01/01/2020,01/02/2020")
	flagSet.StringVar(&config.out, "out", "", "out file path")
	return nil
}

func (f *Command) Run(args []string) error {
	dates := strings.Split(config.dates, ",")

	dateFrom, err := time.Parse(schedule.TimeFormat, dates[0])
	if err != nil {
		return err
	}

	dateTo, err := time.Parse(schedule.TimeFormat, dates[1])
	if err != nil {
		return nil
	}

	config := merge.Config{
		DateFrom:  dateFrom,
		DateTo:    dateTo,
		Paths:     strings.Split(config.paths, ","),
		Employees: strings.Split(config.employees, ","),
		Out:       config.out,
	}

	return merge.Merge(config)
}
