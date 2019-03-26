package goshell

import (
	"errors"
	"fmt"
	"strings"
)

// Shell struct embeds exec.Cmd
type Shell struct {
	ShowCommands bool
	ShowOutput   bool
	StopOnError  bool
	Err          error
	Dir          string
	Env          []string
	ErrorCommand *Command
	Commands     []*Command
}

// New creates a shell
func New() *Shell {
	return &Shell{
		ShowCommands: true,
		ShowOutput:   true,
		StopOnError:  true,
	}
}

// AddEnv adds an environment variable the the virtual shell executions
func (shell *Shell) AddEnv(key string, value string) {
	shell.Env = append(shell.Env, fmt.Sprintf("%s=%s", key, value))
}

// Exec a command using a virtual shell
func (shell *Shell) Exec(name string, arg ...string) (*Command, error) {
	if shell.Err != nil && shell.StopOnError {
		return nil, errors.New("a previous command failed")
	}

	cmd := &Command{
		Name:       name,
		Arg:        arg,
		Env:        shell.Env,
		Dir:        shell.Dir,
		ShowOutput: shell.ShowOutput,
	}

	if shell.ShowCommands {
		fmt.Printf("\033[0;36m%s %s\033[0m\n", name, strings.Join(arg, " "))
	}

	err := cmd.Exec()

	shell.Commands = append(shell.Commands, cmd)

	if err != nil {
		shell.Err = err
		shell.ErrorCommand = cmd
	}

	return cmd, err
}
