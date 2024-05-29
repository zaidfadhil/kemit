# kemit
Automate the process of generating commit messages based on the diff of staged files in a Git repository

## Requirements
- Git
- [Ollama](https://ollama.com) (with llama3 or mistral)
- Supported platforms: macOS, Linux (currently Arch only)

## Installation

using [homebrew](https://brew.sh).
```shell
brew install kemit
```

or download the binary directly from github
```shell
sudo curl -sSL https://raw.githubusercontent.com/zaidfadhil/kemit/main/install.sh | sh
```

## Setup
Make sure you have [ollama](https://ollama.com) installed and running in `serve` mode.

To set or update the configuration, use the config command:

```shell
kemit config [options]
```
- `--ollama_host`: Set the Ollama host. Example: http://localhost:11434. (required).
- `--ollama_model`: Set the Ollama model. Example: llama3. (required).

example:
```shell
kemit config --ollama_host http://localhost:11434 --ollama_model llama3
```

## Usage

To generate a commit message:

```shell
kemit
```

If there are no staged changes, the application will output "nothing to commit". Otherwise, it will generate and print a commit message based on the staged diff.

## License
kemit is licensed under the [MIT License](https://github.com/zaidfadhil/kemit/blob/master/LICENSE).
