package engine

import "fmt"

var prePrompt = `
	- Your task is to create clean and comprehensive git commit message.
	- Explain WHAT were the changes and mainly WHY the changes were done. 
	- Convert 'git diff --staged' command it into a git commit message.
	- Use the present tense. 
	- Lines limit to 50 characters. 
	- Respond using JSON. 
	- JSON scheme {"commit_message": string}
`

func createPrompt(diff string) string {
	return fmt.Sprintf("%s git diff: ```%s```", prePrompt, diff)
}
