package main

import (
	"fmt"

	"github.com/zaidfadhil/kemit.git/engines"
	"github.com/zaidfadhil/kemit.git/git"
)

func main() {
	diff, err := git.DiffFiles()
	if err != nil {
		fmt.Println("error:", err)
	}

	if len(diff) == 0 {
		fmt.Println("nothing to commit")
	} else {
		ollama := engines.NewOllama()
		message, err := ollama.GetCommit(diff)
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println(message)
		}
	}

}
