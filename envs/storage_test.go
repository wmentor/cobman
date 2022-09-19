package envs_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wmentor/cobman/envs"
)

func TestEnvs(t *testing.T) {
	t.Parallel()

	limit := 25
	storage := envs.NewStorage()
	list := make([]envs.Env, 0, limit)

	for i := 1; i <= limit; i++ {
		name := fmt.Sprintf("name %02d", i)
		description := fmt.Sprintf("desc %d", i)

		storage.Set(name, description)

		list = append(list, envs.Env{Name: strings.ToUpper(name), Description: description})
	}

	resultList := storage.List()
	require.Equal(t, list, resultList)
}
