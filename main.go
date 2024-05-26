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

	fmt.Println(diff)
}
