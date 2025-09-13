# file2kv-yaml

A command-line tool that converts file contents into key-value YAML format. It reads file paths from stdin and generates individual YAML files with a specific key-value structure.

## Features

- Reads file paths from standard input (stdin)
- Converts each file's content into a separate YAML file
- Generates YAML files with key-value structure
- Supports only UTF-8 text files

## Installation

```bash
$ go install github.com/zinrai/file2kv-yaml@latest
```

## Usage

```bash
$ find ./source_directory -type f ! -path '*/\.*' | file2kv-yaml ./output_directory
```

## File Path Conversion Rules

Input file paths are converted to YAML filenames using these rules:

- `/` -> `_` (directory separator to underscore)
- `-` -> `_` (hyphen to underscore)
- `.` -> `_` (dot to underscore)
- Leading `./` is removed
- `.yaml` extension is appended

### Conversion Examples

| Input Path                   | Output Filename                 | Key                        |
|------------------------------|---------------------------------|----------------------------|
| `./hcnfrlw/zorrzzwpgmr.jbk`  | `hcnfrlw_zorrzzwpgmr_jbk.yaml`  | `hcnfrlw_zorrzzwpgmr_jbk`  |
| `./data/log-file.txt`        | `data_log_file_txt.yaml`        | `data_log_file_txt`        |
| `./config/app-settings.json` | `config_app_settings_json.yaml` | `config_app_settings_json` |

## YAML Output Structure

Each generated YAML file follows this structure:

```yaml
key: <converted_filename_without_extension>
value: |
  <original_file_content>
```

### Example

**Input file:** `./example/hello.txt`

```
Hello, World!
This is a test file.
```

**Output file:** `example_hello_txt.yaml`

```yaml
key: example_hello_txt
value: |
  Hello, World!
  This is a test file.
```

## Error Handling

The tool follows a **fail-fast** approach for data integrity. It will stop immediately if:

- **File not found**: Input file doesn't exist
- **Permission denied**: Cannot read input file
- **Binary file detected**: File is not valid UTF-8 text
- **Write error**: Cannot write to output directory

The only exception is files in the output directory, which are automatically skipped (not an error):

## License

This project is licensed under the [MIT License](./LICENSE).
