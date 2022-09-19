// The envs package defines the logic for working with environment variable declarations.
package envs

import (
	"sort"
	"strings"
)

// Env - environment variable declarations.
type Env struct {
	Name        string
	Description string
}

// Storage for environment variable declarations.
type Storage struct {
	data map[string]Env
}

// NewStorage creates new storage for environment variable declarations.
func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]Env),
	}
}

// Set new environment variable declaration.
func (st *Storage) Set(name string, description string) {
	st.data[name] = Env{
		Name:        strings.ToUpper(name),
		Description: description,
	}
}

// Return list of environment variable declarations.
func (st *Storage) List() []Env {
	list := make([]Env, 0, len(st.data))

	for _, v := range st.data {
		list = append(list, v)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})

	return list
}
