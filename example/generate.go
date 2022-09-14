package main

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/wmentor/cobman"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "mycli",
		Short: "my little command line tool",
		RunE: func(_ *cobra.Command, args []string) error {
			return nil
		},
	}

	cobman.SetRootCommand(rootCmd)
	cobman.SetProgram("mycli")
	cobman.SetSection(1)
	cobman.SetDate("23 March, 2016")
	cobman.SetOS("GNU Linux")
	cobman.SetMachine("My simple util")

	cobman.SetCommandDescription(rootCmd, `Use the .Nm macro to refer to your program throughout the man page like
such: Untitled, Underlining is accomplished with the .Ar macro like this:
underlined text.

A list of items with descriptions:

item a   Description of item a

item b   Description of item b`)

	ioutil.WriteFile("./example.1", cobman.MakeMan(), 0644)
}
