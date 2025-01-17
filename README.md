# This is golog (Go-Log)
It is a simple log parsing tool.

#### Currently expected log template
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

## Commands to run on system
Go v1.23.4 was used in development
   
```bash
# Clone the repository
$ git clone https://github.com/gitKashish/golog.git
$ cd golog

# Build the executable
$ go build

# Run the executable
$ ./golog <log_file_path>
```
