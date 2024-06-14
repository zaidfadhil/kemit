package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/zaidfadhil/kemit/config"
	"github.com/zaidfadhil/kemit/engine"
	"github.com/zaidfadhil/kemit/git"
)

//go:embed VERSION
var version string

func main() {
	cfg := &config.Config{}
	err := cfg.Load()
	if err != nil {
		end(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "config" {
			err := setConfig(os.Args[2:], cfg)
			if err != nil {
				end(err)
			}
		} else if os.Args[1] == "version" {
			fmt.Println("kemit version: " + version)
		}
	} else {
		err = cfg.Validate()
		if err != nil {
			end(err)
		}

		run(cfg)
	}
}

func run(cfg *config.Config) {
	diff, err := git.Diff()
	if err != nil {
		end(err)
	}

	if diff == "" {
		fmt.Println("nothing to commit")
	} else {
		ollama := engine.NewOllama(cfg.OllamaHost, cfg.OllamaModel)
		message, err := ollama.GetCommit(diff)
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

var flags = flag.NewFlagSet("config", flag.ExitOnError)

var configUsage = `Usage: kemit config command [options]

Options:

Commands:
	--provider			Set LLM Provider. default Ollama
	--ollama_host			Set ollama host. ex: http://localhost:11434
	--ollama_model			Set ollama host. ex: llama3`

func setConfig(args []string, cfg *config.Config) error {
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, configUsage)
	}

	flags.StringVar(&cfg.Provider, "provider", cfg.Provider, "llm model provider. ex: ollama")
	flags.StringVar(&cfg.OllamaHost, "ollama_host", cfg.OllamaHost, "ollama host")
	flags.StringVar(&cfg.OllamaModel, "ollama_model", cfg.OllamaModel, "ollama model")

	err := flags.Parse(args)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		flags.Usage()
	} else {
		err = cfg.Save()
		if err != nil {
			return err
		}
	}

	return nil
}
