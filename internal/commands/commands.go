package commands

import (
	"fmt"

	"github.com/arkkis27/gator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Functions map[string]func(*state.State, Command) error
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.Functions[name] = f
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	if f, ok := c.Functions[cmd.Name]; !ok {
		return fmt.Errorf("command not found: %v", cmd.Name)
	} else {
		return f(s, cmd)
	}
}
