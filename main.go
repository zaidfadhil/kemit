package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/zaidfadhil/kemit/cli"
	"github.com/zaidfadhil/kemit/config"
	"github.com/zaidfadhil/kemit/engine"
	"github.com/zaidfadhil/kemit/git"
)

//go:embed VERSION
var version string

func main() {
	cmd := cli.New()

	cfg := &config.Config{}
	if err := cfg.Load(); err != nil {
		end(err)
	}

	// config
	var configCmd *cli.Command
	configCmd = cmd.AddCommand("config", "Set the configuration", func(_ []string) {
		cfg.Provider = configCmd.Flags.Lookup("provider").Value.String()
		cfg.OllamaHost = configCmd.Flags.Lookup("ollama_host").Value.String()
		cfg.OllamaModel = configCmd.Flags.Lookup("ollama_model").Value.String()
		cfg.CommitStyle = configCmd.Flags.Lookup("commit_style").Value.String()
		if err := cfg.Save(); err != nil {
			end(err)
		}
	})
	configCmd.Flags.String("provider", cfg.Provider, "LLM models provider. ex: ollama")
	configCmd.Flags.String("ollama_host", cfg.OllamaHost, "Ollama host. ex: localhost:11434")
	configCmd.Flags.String("ollama_model", cfg.OllamaModel, "Ollama model. ex: llama3.1")
	configCmd.Flags.String("commit_style", cfg.CommitStyle, "Commit style. ex: conventional-commit")

	// version
	cmd.AddCommand("version", "Show the version", func(_ []string) {
		fmt.Println("version:", strings.TrimSpace(version))
	})

	// help
	cmd.AddCommand("help", "Help about any command", func(_ []string) {
		cmd.PrintHelp()
	})

	cmd.SetDefaultCommand(cmd.AddCommand("", "", func(_ []string) {
		if err := cfg.Validate(); err != nil {
			end(err)
		}

		run(cfg)
	}))

	if err := cmd.Execute(); err != nil {
		end(err)
	}
}

func run(cfg *config.Config) {
	diff, err := git.Diff()
	if err != nil {
		end(err)
	}

	if diff == "" {
		fmt.Println("Nothing to commit")
	} else {
		// TODO: move this to the enigne pkg
		ollama := engine.NewOllama(cfg.OllamaHost, cfg.OllamaModel)
		message, err := ollama.GetCommitMessage(diff, cfg.CommitStyle)
		if err != nil {
			end(err)
		} else {
			fmt.Println(message)
		}
	}
}

func end(err error) {
	fmt.Println("error:", err)
	os.Exit(1)
}
