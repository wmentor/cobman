package cobman

var (
	envMap = map[string]string{}
)

func SetEnv(name string, description string) {
	envMap[name] = description
}
