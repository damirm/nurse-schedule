package fill

import (
	"flag"
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
	return "fill"
}

func (f *Command) ExportFlags(flagSet *flag.FlagSet) error {
	return nil
}

func (f *Command) Run() error {
	return nil
}
