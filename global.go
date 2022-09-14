package cobman

import (
	"github.com/spf13/cobra"
)

var (
	globalGen = New()
)

func SetRootCommand(cmd *cobra.Command) {
	globalGen.SetRootCommand(cmd)
}

func MakeMan() []byte {
	return globalGen.MakeMan()
}

func SetProgram(name string) {
	globalGen.SetProgram(name)
}

func SetSection(num int) error {
	return globalGen.SetSection(num)
}

func SetDate(dt string) {
	globalGen.SetDate(dt)
}

func SetOS(osName string) {
	globalGen.SetOS(osName)
}

func SetMachine(name string) {
	globalGen.SetMachine(name)
}

func SetCommandDescription(cmd *cobra.Command, text string) {
	globalGen.SetCommandDescription(cmd, text)
}

func SetCommandExample(cmd *cobra.Command, text string) {
	globalGen.SetCommandExample(cmd, text)
}
