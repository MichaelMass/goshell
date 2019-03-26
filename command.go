package goshell

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"sync"
)

// Command is a wrapper around exec.Cmd
type Command struct {
	Stdout     []byte
	Stderr     []byte
	Name       string
	Arg        []string
	Env        []string
	Dir        string
	ShowOutput bool
}

// Exec executes the command
func (command *Command) Exec() error {
	cmd := exec.Command(command.Name, command.Arg...)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var stdoutBuf, stderrBuf bytes.Buffer
	stdout, stderr := command.createIoWriter(&stdoutBuf, &stderrBuf)

	if command.Dir != "" {
		cmd.Dir = command.Dir
	}

	cmd.Env = append(os.Environ(), command.Env...)

	err := cmd.Start()

	if err != nil {
		return err
	}

	var errStdout, errStderr error
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()

	command.Stdout = stdoutBuf.Bytes()
	command.Stderr = stderrBuf.Bytes()

	if err != nil {
		return err
	}

	if errStdout != nil {
		return errors.New("error while capturing stdout")
	}

	if errStderr != nil {
		return errors.New("error while capturing stderr")
	}

	return nil
}

func (command *Command) createIoWriter(stdoutBuf io.Writer, stderrBuf io.Writer) (io.Writer, io.Writer) {
	if command.ShowOutput {
		stdout := io.MultiWriter(os.Stdout, stdoutBuf)
		stderr := io.MultiWriter(os.Stderr, stderrBuf)

		return stdout, stderr
	}

	return stdoutBuf, stderrBuf
}
