package cobman

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func SetFlagValueName(cmd *cobra.Command, flag string, name string) {
	cmd.Flags().SetAnnotation(flag, keyFlagValueName, []string{name})
}

func getFlagValueName(f *pflag.Flag) string {
	if f.Annotations != nil {
		if val, has := f.Annotations[keyFlagValueName]; has && len(val) == 1 && val[0] != "" {
			return val[0]
		}
	}

	if val := f.Value.Type(); val != "bool" {
		return val
	}

	return ""
}
