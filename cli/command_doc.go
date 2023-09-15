// This is just to generate documetnation for command line
// This file can be safely removed from the build

package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	commandDoc = cli.Command{
		Name:   "doc",
		Hidden: true,
		Action: func(ctx *cli.Context) error {
			for _, c := range commands {
				if c.Hidden {
					continue
				}

				// Command
				fmt.Printf("- `%s`", c.Name)
				for _, a := range c.Aliases {
					fmt.Printf(" , `%s`", a)
				}
				fmt.Printf(" - %s\n", c.Usage)

				// Flags
				for _, f := range c.Flags {
					fmt.Printf("\t- %s\n", f.String())
				}

			}
			return nil
		},
	}
)

func init() {
	commands = append(commands, &commandDoc)
}
