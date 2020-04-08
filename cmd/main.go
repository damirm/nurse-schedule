package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/damirm/nurse-schedule/cmd/fill"
	"github.com/damirm/nurse-schedule/cmd/merge"
)

// commands: fill, merge

type Command interface {
	Name() string
	ExportFlags(*flag.FlagSet) error
	Run(args []string) error
}

var subcommands = []Command{
	fill.NewCommand(),
	merge.NewCommand(),
}

func root() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Unknown subcommand")
	}

	subcommand := os.Args[1]

	for _, cmd := range subcommands {
		if subcommand == cmd.Name() {
			subset := flag.NewFlagSet(cmd.Name(), flag.PanicOnError)
			cmd.ExportFlags(subset)

			if err := subset.Parse(os.Args[2:]); err != nil {
				flag.PrintDefaults()
				return err
			}

			return cmd.Run(os.Args[2:])
		}
	}

	return fmt.Errorf("Unknown submcommand %s", subcommand)
}

func main() {
	if err := root(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
