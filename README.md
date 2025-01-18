`golog` (go-log) is a simple log parsing tool.

## Usage
```yaml
Usage   : golog -source=<source_file_path> [-target=<output_file_path>] [-show]
Options :
  -source : File path of source file with unformatted logs. (required)
  -target : File path where you want to store the formatted logs. File is created if it does not exist.
  -show   : Set flag to print formatted logs on console even if -target flag is set
```

## Getting started
Go v1.23.4 was used in development
   
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
