package cobman

import (
	"github.com/wmentor/cobman/envs"
)

var (
	envStorage = envs.NewStorage()
)

func SetEnv(name string, description string) {
	envStorage.Set(name, description)
}
