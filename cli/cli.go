package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	ErrUnknownCLICommand = errors.New("unknown cli command")
)

type Command struct {
	Name        string
	Description string
	Action      func(args []string)
	Flags       *flag.FlagSet
}

type CLI struct {
	commands       map[string]*Command
	defaultCommand *Command
}

func New() *CLI {
	return &CLI{
		commands: make(map[string]*Command),
	}
}

func (cli *CLI) AddCommand(name string, description string, action func(args []string)) *Command {
	cmd := &Command{
		Name:        name,
		Description: description,
		Action:      action,
		Flags:       flag.NewFlagSet(name, flag.ExitOnError),
	}
	cli.commands[name] = cmd
	return cmd
}

func (cli *CLI) SetDefaultCommand(cmd *Command) {
	cli.defaultCommand = cmd
}

func (cli *CLI) Execute() error {
	if len(os.Args) == 1 {
		if cli.defaultCommand != nil {
			cli.defaultCommand.Action(nil)
			return nil
		}
		cli.PrintHelp()
		return nil
	}

	cmdName := os.Args[1]
	cmd, exists := cli.commands[cmdName]
	if !exists {
		if cmdName == "-h" || cmdName == "-help" || cmdName == "--help" {
			cli.PrintHelp()
			return nil
		} else {
			return ErrUnknownCLICommand
		}
	}

	if err := cmd.Flags.Parse(os.Args[2:]); err != nil {
		return err
	}

	cmd.Action(cmd.Flags.Args())
	return nil
}

func (cli *CLI) PrintHelp() {
	fmt.Println("Usage:")
	fmt.Println("  kemit [command]")
	fmt.Println("")
	fmt.Println("Available Commands:")

	maxLen := 0
	for name := range cli.commands {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	for _, cmd := range cli.commands {
		if cmd.Name != "" {
			fmt.Printf("  %-*s  %s\n", maxLen+2, cmd.Name, cmd.Description)
		}
	}

	fmt.Println("")
	fmt.Println("Use 'kemit [command] -h' for more information about a command.")
}
