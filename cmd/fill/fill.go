package fill

import "flag"

func NewCommand() *Command {
	return &Command{}
}

type Config struct {
	employees string
	dates     string
}

var config = Config{}

type Command struct {
}

func (f *Command) Name() string {
	return "fill"
}

func (f *Command) ExportFlags(subset *flag.FlagSet) error {
	subset.StringVar(&config.employees, "employees", "", "comma-separateed list of employees names")
	subset.StringVar(&config.dates, "dates", "", "comma-separateed list of start and end dates")
	return nil
}

func (f *Command) Run(args []string) error {
	return nil
}
