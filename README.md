# golog

`golog` is a simple and efficient log parsing and formatting tool.  It allows you to define templates to extract and present log data in a structured and readable way.

[![Go Report Card](https://goreportcard.com/badge/github.com/gitKashish/golog)](https://goreportcard.com/report/github.com/gitKashish/golog) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## 🚀 Getting Started

### Installation

#### Binary Installation (Recommended)

```bash
# Install the latest version
go install github.com/gitKashish/golog@latest
```

This will install the `golog` binary to your `$GOPATH/bin` directory (or `$GOBIN` if set). Make sure this directory is in your `PATH`.

#### Building from Source

```bash
# Clone repository
git clone https://github.com/gitKashish/golog.git

# Build binary
cd golog
go build
```

This will create a `golog` executable in your current directory.

### Creating `template.yaml`

`golog` relies on a `template.yaml` file to define how logs should be parsed and formatted. Create this file in the same directory as the `golog` binary.

## 📄 Template Format

The `template.yaml` file defines two key templates: `sourceTemplate` and `targetTemplate`.

### 1. `sourceTemplate`

The `sourceTemplate` describes the structure of your *incoming* log lines. It uses a specific format for defining fields:

```
@fieldName-fieldType@
```

*   **`fieldName`:** The name of the field (must be unique, alphanumeric characters and underscores only).
*   **`fieldType`:** The data type of the field.

Supported Field Types:

| Type      | Symbol | Description                                                              |
| --------- | :----: | ------------------------------------------------------------------------ |
| Raw       | `raw`  | Value is returned as is (no formatting).                               |
| Number    | `number` | Value is treated as a number.                                           |
| String    | `string` | Value is treated as a string.                                          |
| JSON      | `json`  | Value is parsed as a JSON string and pretty-printed.                   |
| Timestamp | `timestamp` | Value is parsed as a timestamp and formatted into RFC822Z format. |
| Default   | N/A    | Used internally when the log doesn't match the `sourceTemplate`.        |

### 2. `targetTemplate`

The `targetTemplate` defines how the *output* should be formatted. It uses the field names defined in the `sourceTemplate`:

```
@fieldName@
```

A field name can be used multiple times in the `targetTemplate`.

## ✨ Example

**Log Input:**

```
5|3022  | -->2025-01-24 07:29:52.954 :----: users :=: updateJobStatus :=: {"EVENT":"deleteNotificationTaskFromScheduler","ERROR":{"errno":-110,"code":"ETIMEDOUT","syscall":"connect","address":"52.41.75.101","port":3013}}
```

**`template.yaml`:**

```yaml
sourceTemplate: "@server-number@|@instance-number@  | -->@time-timestamp@ :----: @module-string@ :=: @api-string@ :=: @details-json@"
targetTemplate: |
  ---------------------------
  Server: @server@
  Instance: @instance@
  Timestamp: @time@
  API: @api@
  Module : @module@
  Details: @details@
  ---------------------------
```

## 🔎 Usage

### `golog write`

Formats log and writes to a specified file.

```bash
golog write -i <input_file> -o <output_file> [-s]
```

*   `-i, --input string`: Path to the input log file (required).
*   `-o, --output string`: Path to the output file (required).
*   `-s, --show`: Also print the output to the console.

### `golog show`

Formats log and prints the output to the console.

```bash
golog show -i <input_file>
```

*   `-i, --input string`: Path to the input log file (required).

## ⚙️ Development

Go v1.20 or later is recommended.

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## 📄 License
This project is licensed under the GNU Lesser General Public License v3.0. See the [LICENSE](LICENSE) file for details.
