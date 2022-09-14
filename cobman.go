package cobman

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/wmentor/cobman/man"
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

func New() *DocGen {
	return &DocGen{
		commands: make(map[*cobra.Command]*cmdInfo),
		date:     time.Now().Format(timeFormat),
		os:       "Application",
		section:  1,
		program:  "UTIL",
		machine:  "Application",
	}
}

func (gen *DocGen) SetRootCommand(cmd *cobra.Command) {
	gen.rootCmd = cmd
}

func (gen *DocGen) SetProgram(name string) {
	gen.program = name
}

func (gen *DocGen) SetSection(secNum int) error {
	if secNum >= 1 && secNum <= 9 {
		gen.section = secNum
		return nil
	}
	return ErrInvalidSection
}

func (gen *DocGen) SetOS(osName string) {
	gen.os = osName
}

func (gen *DocGen) SetDate(dt string) {
	gen.date = dt
}

func (gen *DocGen) SetMachine(name string) {
	gen.machine = name
}

func (gen *DocGen) SetCommandDescription(cmd *cobra.Command, text string) {
	if rec, has := gen.commands[cmd]; has {
		rec.Description = text
	} else {
		gen.commands[cmd] = &cmdInfo{Description: text}
	}
}

func (gen *DocGen) SetCommandExample(cmd *cobra.Command, text string) {
	if rec, has := gen.commands[cmd]; has {
		rec.Example = text
	} else {
		gen.commands[cmd] = &cmdInfo{Example: text}
	}
}

func (gen *DocGen) MakeMan() []byte {
	buffer := bytes.NewBuffer(nil)

	gen.writeTH(buffer)
	gen.writeName(buffer)
	gen.writeSynopsis(buffer)
	gen.writeDescription(buffer)

	return buffer.Bytes()
}

func (gen *DocGen) writeTH(buffer *bytes.Buffer) {
	buffer.WriteString(".TH \"")
	buffer.Write(man.QuoteEscape(bytes.ToUpper([]byte(gen.program))))
	buffer.WriteString("\" \"")
	buffer.WriteString(strconv.Itoa(gen.section))
	buffer.WriteString("\" \"")
	buffer.Write(man.QuoteEscape([]byte(gen.date)))
	buffer.WriteString("\" \"")
	buffer.Write(man.QuoteEscape([]byte(gen.os)))
	buffer.WriteString("\" \"")
	buffer.Write(man.QuoteEscape([]byte(gen.machine)))
	buffer.WriteString("\"\n")
}

func (gen *DocGen) writeName(buffer *bytes.Buffer) {
	buffer.WriteString(".SH \"NAME\"\n")
	buffer.Write(man.Escape([]byte(gen.rootCmd.Name())))
	buffer.WriteString(" \\- ")
	buffer.Write(man.Escape([]byte(gen.rootCmd.Short)))
	buffer.WriteRune('\n')
}

func (gen *DocGen) writeDescription(buffer *bytes.Buffer) {
	info, has := gen.commands[gen.rootCmd]
	if !has || info.Description == "" {
		return
	}

	buffer.WriteString(".SH \"DESCRIPTION\"\n.sp\n")
	buffer.WriteString(Md2Man(info.Description))
	buffer.WriteString("\n")
}

func (gen *DocGen) writeSynopsis(buffer *bytes.Buffer) {
	buffer.WriteString(".SH \"SYNOPSIS\"\n")

	gen.eachCommandSynopsis(buffer, gen.rootCmd, "")

}

func (gen *DocGen) eachCommandSynopsis(buf *bytes.Buffer, curCmd *cobra.Command, prefix string) {
	if !curCmd.IsAvailableCommand() {
		return
	}

	if prefix != "" {
		prefix = prefix + " " + curCmd.Name()
	} else {
		prefix = curCmd.Name()
		if gen.hasCommonFlags() {
			prefix = prefix + " [common_flags]"
		}
	}

	if curCmd.Runnable() {
		cmdStr := prefix
		if flagStr, has := gen.makeFlagsList(curCmd.Flags()); has {
			cmdStr = cmdStr + " " + flagStr
		}
		buf.WriteString(".sp\n")
		buf.WriteString(".RS 0\n")
		buf.Write(man.Escape([]byte(cmdStr)))
		buf.WriteString("\n.RE\n")
	}

	for _, cmd := range curCmd.Commands() {
		gen.eachCommandSynopsis(buf, cmd, prefix)
	}
}

func (gen *DocGen) hasCommonFlags() bool {
	return gen.rootCmd.PersistentFlags().HasFlags()
}

func (gen *DocGen) makeFlagsList(fl *pflag.FlagSet) (string, bool) {
	if !fl.HasFlags() {
		return "", false
	}

	maker := strings.Builder{}

	fl.VisitAll(func(f *pflag.Flag) {
		if maker.Len() > 0 {
			maker.WriteRune(' ')
		}
		maker.WriteRune('[')
		maker.WriteString("--")
		maker.WriteString(f.Name)
		maker.WriteRune('|')
		maker.WriteRune('-')
		maker.WriteString(f.Shorthand)

		if tp := f.Value.Type(); tp != "bool" {
			maker.WriteRune(' ')
			maker.WriteString(tp)
		}
		maker.WriteRune(']')
	})

	return maker.String(), true
}
