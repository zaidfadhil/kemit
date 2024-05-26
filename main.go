package main

import (
	"fmt"

	"github.com/zaidfadhil/kemit/engine"
	"github.com/zaidfadhil/kemit/git"
)

func main() {
	diff, err := git.Diff()
	if err != nil {
		fmt.Println("error:", err)
	}

	if len(diff) == 0 {
		fmt.Println("nothing to commit")
	} else {
		ollama := engine.NewOllama(diff)
		message, err := ollama.GetCommit()
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println(message)
		}
	}
}
