package cobman

import (
	"bytes"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/wmentor/cobman/man"
)

func MakeMan(rootCmd *cobra.Command) []byte {
	buffer := bytes.NewBuffer(nil)

	manWriteTH(buffer, rootCmd)
	manWriteName(buffer, rootCmd)
	manWriteSynopsis(buffer, rootCmd)
	manWriteDescription(buffer, rootCmd)
	manWriteCommonFlags(buffer, rootCmd)
	manWriteEnvs(buffer)
	manWriteSeeAlso(buffer)

	return buffer.Bytes()
}

func manWriteCommonFlags(buffer *bytes.Buffer, rootCmd *cobra.Command) {
	buffer.WriteString(".SH \"COMMON OPTIONS\"\n")

	buffer.WriteString(".PP\n")
	buffer.WriteString("\\fB")
	buffer.WriteString(man.Escape("--help, -h"))
	buffer.WriteString("\\fR\n")
	buffer.WriteString(".RS 4\n")
	buffer.WriteString(man.Escape("Show brief usage information."))
	buffer.WriteString("\n.RE\n")

	if rootCmd.PersistentFlags().HasFlags() {
		rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
			if f.Hidden || f.Deprecated != "" {
				return
			}

			buffer.WriteString(".PP\n")
			buffer.WriteString("\\fB")
			buffer.WriteString("\\-\\-")
			buffer.WriteString(man.Escape(f.Name))
			if f.ShorthandDeprecated == "" && f.Shorthand != "" {
				buffer.WriteString(", ")
				buffer.WriteString("\\-")
				buffer.WriteString(man.Escape(f.Shorthand))
			}
			buffer.WriteString("\\fR")

			if v := getFlagValueName(f); v != "" {
				buffer.WriteRune(' ')
				buffer.WriteString("\\fI")
				buffer.WriteString(man.Escape(v))
				buffer.WriteString("\\fR")
			}

			buffer.WriteString("\n.RS 4\n")
			buffer.WriteString(man.Escape(f.Usage))
			buffer.WriteString("\n.RE\n")
		})
	}

	if rootCmd.Version != "" {
		buffer.WriteString(".PP\n")
		buffer.WriteString("\\fB")
		buffer.WriteString(man.Escape("--version, -v"))
		buffer.WriteString("\\fR\n")
		buffer.WriteString(".RS 4\n")
		buffer.WriteString(man.Escape("Show shardman-utils version information."))
		buffer.WriteString("\n.RE\n")
	}
}

func manWriteTH(buffer *bytes.Buffer, rootCmd *cobra.Command) {
	buffer.WriteString(".TH \"")
	buffer.WriteString(man.QuoteEscape(strings.ToUpper(rootCmd.Name())))
	buffer.WriteString("\" \"")
	buffer.WriteString(strconv.Itoa(thSection))
	buffer.WriteString("\" \"")
	buffer.WriteString(man.QuoteEscape(thDate))
	buffer.WriteString("\" \"")
	buffer.WriteString(man.QuoteEscape(thOS))
	buffer.WriteString("\" \"")
	buffer.WriteString(man.QuoteEscape(thTitle))
	buffer.WriteString("\"\n")
}

func manWriteName(buffer *bytes.Buffer, rootCmd *cobra.Command) {
	buffer.WriteString(".SH \"NAME\"\n.PP\n")
	buffer.WriteString(man.Escape(rootCmd.Name()))
	buffer.WriteString(" \\- ")
	buffer.WriteString(man.Escape(rootCmd.Short))
	buffer.WriteRune('\n')
}

func manWriteDescription(buffer *bytes.Buffer, rootCmd *cobra.Command) {
	if val := manWriteSingleDescription(rootCmd, true); val != "" {
		buffer.WriteString(".SH \"DESCRIPTION\"\n")
		buffer.WriteString(val)
	}
}

func manWriteSingleDescription(cmd *cobra.Command, isRoot bool) string {
	buffer := strings.Builder{}

	if cmd.Hidden {
		return ""
	}

	if val, has := getCommandAnnotationKey(cmd, keyCmdDescription); has {
		if !isRoot {
			path := cmd.CommandPath()
			buffer.WriteString(".SS \"")
			buffer.WriteString(man.QuoteEscape(path))
			buffer.WriteString("\"\n")
		}
		buffer.WriteString(man.Md2Man(val))
		buffer.WriteString("\n")
	}

	for _, childCmd := range cmd.Commands() {
		if val := manWriteSingleDescription(childCmd, false); val != "" {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

func manWriteSynopsis(buffer *bytes.Buffer, cmd *cobra.Command) {
	buffer.WriteString(".SH \"SYNOPSIS\"\n.PP\n")

	manEachCommandSynopsis(buffer, cmd, "")
}

func manWriteSeeAlso(buffer *bytes.Buffer) {
	if len(seeAlso) == 0 {
		return
	}

	buffer.WriteString(".SH \"SEE ALSO\"\n")
	buffer.WriteString(".PP\n")

	names := make([]string, 0, len(seeAlso))
	for name := range seeAlso {
		names = append(names, name)
	}

	sort.Strings(names)

	for i, name := range names {
		if i > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString("\\fB")
		buffer.WriteString(man.Escape(name))
		buffer.WriteString("\\fR")
		buffer.WriteRune('(')
		buffer.WriteString(man.Escape(strconv.Itoa(seeAlso[name])))
		buffer.WriteRune(')')
	}

	buffer.WriteRune('\n')
}

func manWriteEnvs(buffer *bytes.Buffer) {
	if len(envMap) == 0 {
		return
	}

	buffer.WriteString(".SH \"ENVIRONMENT\"\n")

	names := make([]string, 0, len(envMap))

	for name := range envMap {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, env := range names {
		buffer.WriteString(".PP\n")
		buffer.WriteString("\\fB")
		buffer.WriteString(man.Escape(env))
		buffer.WriteString("\\fR\n")
		buffer.WriteString(".RS 4\n")
		buffer.WriteString(man.Md2Man(envMap[env]))
		buffer.WriteString("\n.RE\n")
	}
}

func manEachCommandSynopsis(buf *bytes.Buffer, curCmd *cobra.Command, prefix string) {
	if !curCmd.IsAvailableCommand() {
		return
	}

	commonFlags := ""

	if prefix != "" {
		prefix = prefix + " \\fB" + curCmd.Name() + "\\fR"
	} else {
		prefix = "\\fB" + curCmd.Name() + "\\fR"
		if curCmd.PersistentFlags().HasFlags() {
			prefix = prefix + " [\\fIcommon_options\\fR]"
			commonFlags, _ = manMakeSynopsisFlagList(curCmd.PersistentFlags())
		}
	}

	if curCmd.Runnable() {
		cmdStr := prefix
		if flagStr, has := manMakeSynopsisFlagList(curCmd.Flags()); has {
			cmdStr = cmdStr + " " + flagStr
		}
		buf.WriteString(".sp\n")
		buf.WriteString(".RS 0\n")
		buf.WriteString(cmdStr)
		buf.WriteString("\n.RE\n")
	}

	for _, cmd := range curCmd.Commands() {
		manEachCommandSynopsis(buf, cmd, prefix)
	}

	if commonFlags != "" {
		buf.WriteString(".PP\n.RS 12\nHere \\fIcommon_options\\fR are:\n.RE\n.PP\n")
		buf.WriteString(commonFlags)
		buf.WriteString("\n")
	}
}

func manMakeSynopsisFlagList(fl *pflag.FlagSet) (string, bool) {
	if !fl.HasFlags() {
		return "", false
	}

	maker := strings.Builder{}

	fl.VisitAll(func(f *pflag.Flag) {
		if f.Hidden || f.Deprecated != "" {
			return
		}

		if maker.Len() > 0 {
			maker.WriteRune(' ')
		}
		maker.WriteRune('[')
		maker.WriteString("\\fB\\-\\-")
		maker.WriteString(man.Escape(f.Name))
		if f.ShorthandDeprecated == "" && f.Shorthand != "" {
			maker.WriteRune('|')
			maker.WriteString("\\-")
			maker.WriteString(man.Escape(f.Shorthand))
		}
		maker.WriteString("\\fR")

		if v := getFlagValueName(f); v != "" {
			maker.WriteRune(' ')
			maker.WriteString("\\fI")
			maker.WriteString(man.Escape(v))
			maker.WriteString("\\fR")
		}
		maker.WriteRune(']')
	})

	if maker.Len() > 0 {
		maker.WriteRune(' ')
	}

	maker.WriteString("[\\fB\\-\\-help|\\-h\\fR]")

	return maker.String(), true
}
