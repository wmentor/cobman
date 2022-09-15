package cobman

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/spf13/cobra"
)

type exampleObject struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func SetCommandDescription(cmd *cobra.Command, descriptionMarkdown string) {
	makeCommandAnnotationsIfNotExists(cmd)
	cmd.Annotations[keyCmdDescription] = descriptionMarkdown
}

func AddCommandExample(cmd *cobra.Command, title string, exampleMarkDown string) {
	list := append(getCommandExample(cmd), exampleObject{Title: title, Text: exampleMarkDown})
	buf := bytes.NewBuffer(nil)
	json.NewEncoder(buf).Encode(list)
	cmd.Annotations[keyCmdExample] = buf.String()
}

func getCommandExample(cmd *cobra.Command) []exampleObject {
	makeCommandAnnotationsIfNotExists(cmd)
	list := []exampleObject{}
	if val, has := cmd.Annotations[keyCmdExample]; has {
		if err := json.NewDecoder(strings.NewReader(val)).Decode(&list); err != nil {
			list = []exampleObject{}
		}
	}
	return list
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
