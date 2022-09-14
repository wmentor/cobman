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

	rootCmd.Flags().IntVarP(&rootParam1, "param1", "i", 0, "first parameter (timer)")
	rootCmd.Flags().BoolVarP(&rootParam2, "param2", "b", false, "my second parameter")

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

item a   Description of item a

item b   Description of item b`)

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

	ioutil.WriteFile("./example.1", cobman.MakeMan(rootCmd), 0644)
}
