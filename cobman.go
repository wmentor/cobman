package cobman

import (
	"errors"
	"time"

	"github.com/spf13/cobra"
)

const (
	timeFormat = "02 January, 2006"
)

var (
	ErrInvalidSection = errors.New("invalid section number")
)

type DocGen struct {
	rootCmd  *cobra.Command
	commands map[*cobra.Command]*cmdInfo
	date     string
	os       string
	program  string
	section  int
	machine  string
}

type cmdInfo struct {
	Example     string
	Description string
}

func New(rootCmd *cobra.Command) *DocGen {
	return &DocGen{
		rootCmd:  rootCmd,
		commands: make(map[*cobra.Command]*cmdInfo),
		date:     time.Now().Format(timeFormat),
		os:       "Application",
	}
}

func (gen *DocGen) SetProgram(name string) {
	gen.program = name
}

func (gen *DocGen) SetSection(secNum int) error {
	if secNum >= 1 && secNum <= 6 {
		gen.section = secNum
		return nil
	}
	return ErrInvalidSection
}

func (gen *DocGen) SetOS(osName string) {
	gen.os = osName
}

func (gen *DocGen) SetDate(tm time.Time) {
	gen.date = tm.UTC().Format(timeFormat)
}

func (gen *DocGen) SetMachine(name string) {
	gen.machine = name
}
