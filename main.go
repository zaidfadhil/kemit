package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zaidfadhil/kemit/config"
	"github.com/zaidfadhil/kemit/engine"
	"github.com/zaidfadhil/kemit/git"
)

func main() {
	cfg := &config.Config{}
	cfg.LoadConfig() //nolint:errcheck

	if len(os.Args) > 1 {
		if os.Args[1] == "config" {
			err := setConfig(os.Args[2:], cfg)
			if err != nil {
				fmt.Println("error:", err)
				os.Exit(1)
			}
		}
	} else {
		run(cfg)
	}
}

func run(cfg *config.Config) {
	diff, err := git.Diff()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if diff == "" {
		fmt.Println("nothing to commit")
	} else {
		ollama := engine.NewOllama(cfg.OllamaHost, cfg.OllamaModel)
		message, err := ollama.GetCommit(diff)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		} else {
			fmt.Println(message)
		}
	}
}

var flags = flag.NewFlagSet("config", flag.ExitOnError)

var configUsage = `Usage: kemit config command [options]

Options:

Commands:
	--ollama_host			Set ollama host. ex: http://localhost:11434
	--ollama_model			Set ollama host. ex: llama3`

func setConfig(args []string, cfg *config.Config) error {
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, configUsage)
	}

	flags.StringVar(&cfg.OllamaHost, "ollama_host", cfg.OllamaHost, "ollama host")
	flags.StringVar(&cfg.OllamaModel, "ollama_model", cfg.OllamaModel, "ollama model")

	err := flags.Parse(args)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		flags.Usage()
	}

	err = cfg.SaveConfig()
	if err != nil {
		return err
	}

	return nil
}
