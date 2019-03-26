# Goshell
> Utility helper for Go (golang) os/exec package

`goshell` is a Go (golang) package that wraps [os/exec](https://golang.org/pkg/os/exec/). It extends cmd.Start() to automatically output the cmd's stderr and stdout pipes to the terminal or to a supplied channel. It allows to sequentially execute shell commands.

## Install

`go get github.com/michaelmass/goshell`

## Usage

With goshell you can sequentially execute commands using the `Exec` function.

```go
shell := goshell.New()

shell.Exec("go", "version")
shell.Exec("go", "vet")
shell.Exec("go", "fmt")
```

There are multiple options you can set when creating your shell.

```go
shell := goshell.New()

shell.ShowCommands = false // the command themselves won't be shown in the terminal
shell.ShowOutput = false // the commands output won't be shown in the terminal
shell.StopOnError = false // the shell will continue executing if a previous command failed (exitcode != 0)
shell.Dir = "C:/dev" // path where to execute the commands
shell.Env = []string{"VAR=value"} // adds environment variables to the command

shell.Exec("go", "version")
shell.Exec("go", "vet")
shell.Exec("go", "fmt")
```

## Test

`go test ./...`

## Contribution

If you are interested in fixing issues and contributing directly to the code base, you can directly submit your requests and I will be happily accept them.

## License
The package is available as open source under the terms of the MIT License.
