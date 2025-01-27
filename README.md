`golog` (go-log) is a simple log parsing tool.

## Usage
```yaml
Usage:
  golog [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  show        A brief description of your command
  write       A brief description of your command

Flags:
  -h, --help   help for golog

Use "golog [command] --help" for more information about a command.
```

#### golog `show`
```yaml
Usage:
  golog show [flags]

Flags:
  -h, --help           help for show
  -i, --input string   Path to input file (required)
```

#### golog `write`
```yaml
Usage:
  golog write [flags]

Flags:
  -h, --help            help for write
  -i, --input string    Path to input file (required)
  -o, --output string   Path to output file (required)
  -s, --show            Show output on console
```

## Getting started
Go v1.23.4 was used in development

Install binary and run
```bash
$ go install github.com/gitkashish/golog
$ golog --help
```

Compile yourself
```bash
# Clone the repository
$ git clone https://github.com/gitKashish/golog.git
$ cd golog

# Build the executable
$ go build

# Run the executable
$ ./golog -source="./source/file/path" -target="./target/file/path" -show
```

## I/O Format 
> TODO : make it general purpose
#### Source Log Template
```bash
12|SERVERNAME | ---->2025-01-18 02:40:20.111 :----: module_name :=: api_name :=: {"EVENT" : "LOG DATA",..., "ARR_DATA": [23,43,19]}
```

#### Output format
```bash
Timestamp: 2025-01-18 02:40:20.111
PM ID: 12
Server: SERVERNAME
Module: module_name
API: api_name
Details:
{
  "EVENT": "LOG DATA",
  ...,
  "ARR_DATA": [
    23,
    43,
    19
  ]  
}
--------------------------------------------------------------------------------
```
