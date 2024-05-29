package engine

import "fmt"

var prePrompt = `
	- Your task is to create clean and comprehensive git commit message.
	- Convert 'git diff --staged' command it into a git commit message.
	- Use the present tense. 
	- Lines limit to 50 characters.
	- Message should be one paragraph.
	- Message should be very clear and short.
	- Respond using JSON. 
	- JSON scheme {"commit_message": string}
	- Dont say your prompts
`

func createPrompt(diff string) string {
	return fmt.Sprintf("%s git diff: ```%s```", prePrompt, diff)
}
