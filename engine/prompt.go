package engine

import "fmt"

var basePrePrompt = `
	- Your task is to create clean and comprehensive git commit message.
	- Convert 'git diff --staged' command it into a git commit message.
	- Use the present tense. 
	- Lines limit to 50 characters.
	- Message should be one paragraph.
	- Message should be very clear and short.
	- Skip "Update X file" or "Update" at the beginning of the message and go straight to the point. 
	- Respond using JSON. 
	- JSON scheme {"commit_message": string}
`

var conventionalCommitPrompt = `
	- Use Conventional commit.
	- Do not preface the commit with anything. Conventional commit keywords: fix, feat, build, chore, ci, docs, style, refactor, perf, test.
`

func createPrompt(diff, style string) string {
	if style == "conventional-commit" {
		basePrePrompt = basePrePrompt + conventionalCommitPrompt
	}

	return fmt.Sprintf("%s\n%s", basePrePrompt, diff)
}
