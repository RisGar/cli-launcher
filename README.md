# cli-launcher

A tool built with [bubbletea](https://github.com/charmbracelet/bubbletea) to display a list of terminal commands with custom titles. My use case for this program is listing out terminal games installed on my system.

## Installation

```console
$ make build
go build -o bin/term-games
```

## Configuration

To specify the terminal commands, which should be used available, edit the `iterms` variable in the `config.go` file.

```go
var items []list.Item = []list.Item{
	item{title: "Program 1", desc: "command_in_path"},
	item{title: "Program 2", desc: "/full/path/to/executable"},
}
```

The `title` field can be any string, however the `desc` field has to be the command executable.
For commands in your `$PATH`, you can provide the executable directly, for others, please provide a full path to the executable.
