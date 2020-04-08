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
	Run() error
}

func root() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Unknown subcommand")
	}

	subcommands := []Command{
		fill.NewCommand(),
		merge.NewCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range subcommands {
		if subcommand == cmd.Name() {
			flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
			cmd.ExportFlags(flagSet)
			return cmd.Run()
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
