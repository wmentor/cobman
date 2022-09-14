package cobman

import (
	"github.com/spf13/cobra"
)

func SetCommandDescription(cmd *cobra.Command, descriptionMarkdown string) {
	makeCommandAnnotationsIfNotExists(cmd)
	cmd.Annotations[keyCmdDescription] = descriptionMarkdown
}

func SetCommandExample(cmd *cobra.Command, exampleMarkDown string) {
	makeCommandAnnotationsIfNotExists(cmd)
	cmd.Annotations[keyCmdExample] = exampleMarkDown
}

func makeCommandAnnotationsIfNotExists(cmd *cobra.Command) {
	if cmd.Annotations == nil {
		cmd.Annotations = make(map[string]string)
	}
}

func getCommandAnnotationKey(cmd *cobra.Command, key string) (string, bool) {
	if cmd.Annotations != nil {
		if val, has := cmd.Annotations[key]; has && val != "" {
			return val, has
		}
	}
	return "", false
}
