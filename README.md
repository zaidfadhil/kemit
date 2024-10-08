# kemit
[![Go Reference](https://pkg.go.dev/badge/github.com/zaidfadhil/kemit.svg)](https://pkg.go.dev/github.com/zaidfadhil/kemit)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaidfadhil/kemit)](https://goreportcard.com/report/github.com/zaidfadhil/kemit)

Automate the process of generating commit messages based on the diff of staged files in a Git repository

## Requirements
- Git
- [Ollama](https://ollama.com) (with llama3.1 or mistral ... or any code model)
- Supported platforms: macOS, Linux

## Installation

### Shell script

Install Kemit on macOS or Linux using the following command:

#### Using Curl

```shell
sudo curl -fsSL https://raw.githubusercontent.com/zaidfadhil/kemit/main/scripts/install.sh | sh
```

#### Or using Wget

```shell
sudo wget -qO- https://raw.githubusercontent.com/zaidfadhil/kemit/main/scripts/install.sh | sh
```

### From Source

1. Clone the repository:
```shell
git clone https://github.com/zaidfadhil/kemit.git
cd kemit
```

2. Build and Install:
```shell
make install
```

## Setup
Make sure [ollama](https://ollama.com) installed and running in `serve` mode.

To set or update the configuration, use the config command:

```shell
kemit config [options]
```
- `-provider`: Set the LLM Provider. (default: `ollama`).
- `-ollama_host`: Set the Ollama Host. Example: http://localhost:11434. (required).
- `-ollama_model`: Set the Ollama Model. Example: llama3. (required).
- `-commit_style`: Set the Commit Style. Example: normal, conventional-commit (default: `conventional-commit`)

example:
```shell
kemit config -ollama_host http://localhost:11434 -ollama_model llama3.1
```

## Usage

To generate a commit message:

```shell
kemit
```

If there are no staged changes, the application will output "Nothing to commit". Otherwise, it will generate and print a commit message based on the staged diff.

## Uninstall

To uninstall Kemit, you can use the uninstall script which removes the installed binary:

```shell
sudo curl -fsSL https://raw.githubusercontent.com/zaidfadhil/kemit/main/scripts/uninstall.sh | sh
// or
sudo wget -qO- https://raw.githubusercontent.com/zaidfadhil/kemit/main/scripts/uninstall.sh | sh
```

## License
kemit is licensed under the [MIT License](https://github.com/zaidfadhil/kemit/blob/master/LICENSE).
