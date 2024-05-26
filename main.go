package main

import (
	"fmt"

	"github.com/zaidfadhil/kemit.git/git"
)

func main() {
	diff, err := git.DiffFiles()
	if err != nil {
		fmt.Println(err)
	}

	if len(diff) == 0 {
		fmt.Println("nothing to commit")
	} else {
		fmt.Println(diff)
	}
}
