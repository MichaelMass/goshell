package goshell

import (
	"testing"

	"github.com/kami-zh/go-capturer"
)

func TestExec(t *testing.T) {
	tests := []struct {
		command string
		input   []string
		output  string
	}{
		{"go", []string{"help", "version"}, "usage: go version\n\nVersion prints the Go version, as reported by runtime.Version.\n"},
	}

	var stdoutOutput string
	var stdoutOutputExpected string

	shell := New()
	shell.ShowCommands = false

	for _, test := range tests {
		test := test // pin scope for function literal
		stdoutOutput += capturer.CaptureStdout(func() {
			cmd, _ := shell.Exec(test.command, test.input...)

			output := string(cmd.Stdout)

			stdoutOutputExpected += output

			if output != test.output {
				t.Errorf("exec output was different from expected %s", output)
			}
		})
	}

	if shell.Err != nil {
		t.Error("shell error should be nil")
		t.Error(shell.Err)
	}

	if stdoutOutput != stdoutOutputExpected {
		t.Errorf("stdout output was different from expected %s", stdoutOutput)
	}
}

func TestExecOnError(t *testing.T) {
	tests := []struct {
		command string
		input   []string
		output  string
	}{
		{"go", []string{"help", "help"}, "go help help: unknown help topic. Run 'go help'.\n"},
	}

	shell := New()
	shell.ShowCommands = false
	shell.ShowOutput = false
	shell.StopOnError = false

	for _, test := range tests {
		cmd, _ := shell.Exec(test.command, test.input...)

		output := string(cmd.Stderr)

		if output != test.output {
			t.Errorf("exec output was different from expected %s", output)
		}
	}
}
