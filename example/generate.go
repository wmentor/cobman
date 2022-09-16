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

	var rootParam1 int
	var rootParam2 bool
	var rootCommonParam1 int
	var rootCommonParam2 int

	rootCmd.Flags().IntVarP(&rootParam1, "param1", "b", 0, "first parameter (timer)")
	rootCmd.Flags().BoolVarP(&rootParam2, "interactive", "i", false, "run command in interactive mode")

	rootCmd.PersistentFlags().IntVarP(&rootCommonParam1, "seconds", "s", 0, "seconds")
	cobman.SetPersistentFlagValueName(rootCmd, "seconds", "seconds")

	rootCmd.PersistentFlags().IntVar(&rootCommonParam2, "some-flag", 0, "second common flag")

	rootCmd.CompletionOptions.DisableDescriptions = true

	subCmd1 := &cobra.Command{
		Use:   "subcmd1",
		Short: "my sub command",
		RunE: func(_ *cobra.Command, args []string) error {
			return nil
		},
	}

	cobman.SetCommandDescription(rootCmd, `Use the .Nm macro to refer to your program throughout the man page like
such: Untitled, Underlining is accomplished with the .Ar macro like this:
underlined text. A list of items with descriptions:

1. listitem1 optional  parameters
	1. listitem11
	1. listitem22
		1. 24214 42134 1234 12
		1. 2314 21341 24 12
1. listitem2
1. listitem3

And code block:

`+"```json"+`
{
	"id": 123132,
	"name": "Mikhail K."
}
`+"```\n")

	rootCmd.AddCommand(subCmd1)

	cobman.SetSection(1)
	cobman.SetDate("2016-03-23")
	cobman.SetOS("GNU Linux")
	cobman.SetTitle("My simple util")

	cobman.SetEnv("MY_ETCD1", `optional parameters that are not
specific to the utility. They specify etcd connection settings, cluster
name and a few more settings. By default shardmanctl tries to connect
to the etcd store 127.0.0.1:2379 and use the cluster0 cluster name. The
default log level is info.`)

	cobman.SetEnv("MY_ETCD2", `optional parameters that are not
specific to the utility. They specify etcd connection settings, cluster
name and a few more settings. By default shardmanctl tries to connect
to the etcd store 127.0.0.1:2379 and use the cluster0 cluster name. The
default log level is info.`)

	cobman.AddCommandExample(rootCmd, "Interactive mode", "```\nmycli -i\n```\nThis command runs the **mycli** in interactive mode")
	cobman.AddCommandExample(rootCmd, "Get command help", "```\nmycli -h\n```\nIt will display help on how to use the command")

	cobman.SeeAlso("groff", 1)
	cobman.SeeAlso("man", 7)
	cobman.SeeAlso("man-pages", 7)

	ioutil.WriteFile("./example.1", cobman.MakeMan(rootCmd), 0644)
}
