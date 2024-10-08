package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/zaidfadhil/kemit/cli"
	"github.com/zaidfadhil/kemit/config"
	"github.com/zaidfadhil/kemit/engine"
	"github.com/zaidfadhil/kemit/git"
)

//go:embed VERSION
var version string

func main() {
	app := cli.New()

	cfg := &config.Config{}
	err := cfg.Load()
	if err != nil {
		end(err)
	}

	// config
	var configCmd *cli.Command
	configCmd = app.AddCommand("config", "Set the configuration", func(_ []string) {
		cfg.Provider = configCmd.Flags.Lookup("provider").Value.String()
		cfg.OllamaHost = configCmd.Flags.Lookup("ollama_host").Value.String()
		cfg.OllamaModel = configCmd.Flags.Lookup("ollama_model").Value.String()
		cfg.CommitStyle = configCmd.Flags.Lookup("commit_style").Value.String()
		err := cfg.Save()
		if err != nil {
			end(err)
		}
	})
	configCmd.Flags.String("provider", cfg.Provider, "LLM models provider. ex: ollama")
	configCmd.Flags.String("ollama_host", cfg.OllamaHost, "Ollama host")
	configCmd.Flags.String("ollama_model", cfg.OllamaModel, "Ollama model")
	configCmd.Flags.String("commit_style", cfg.CommitStyle, "Commit style. ex: conventional-commit")

	// version
	app.AddCommand("version", "Show the version", func(_ []string) {
		fmt.Println("kemit version", version)
	})

	// help
	app.AddCommand("help", "Help about any command", func(_ []string) {
		app.PrintHelp()
	})

	app.SetDefaultCommand(app.AddCommand("", "", func(args []string) {
		err := cfg.Validate()
		if err != nil {
			end(err)
		}

		run(cfg)
	}))

	err = app.Run()
	if err != nil {
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
