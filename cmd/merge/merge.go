package merge

import (
	"flag"
	"time"

	"github.com/damirm/nurse-schedule/internal/command/merge"
	"github.com/damirm/nurse-schedule/internal/schedule"
)

func NewCommand() *Command {
	return &Command{}
}

type Config struct {
	employee string
}

var config = Config{}

type Command struct {
}

func (f *Command) Name() string {
	return "merge"
}

func (f *Command) ExportFlags(flagSet *flag.FlagSet) error {
	return nil
}

func (f *Command) Run() error {
	dateFrom, _ := time.Parse(schedule.TimeFormat, "11/03/2020")
	dateTo, _ := time.Parse(schedule.TimeFormat, "11/04/2020")

	config := merge.Config{
		DateFrom:  dateFrom,
		DateTo:    dateTo,
		Paths:     []string{"./test/yc_lb.csv", "./test/yc_network.csv"},
		Employees: []string{"staff:yesworld", "staff:flatline", "staff:raorn", "staff:ovov", "staff:tolmalev", "staff:lavrukov", "staff:wronglink"},
		Out:       "./out/merged.csv",
	}

	return merge.Merge(config)
}
