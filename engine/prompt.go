package engine

import "fmt"

var basePrePrompt = `
*** WHAT TO DO ***
1. Your task is to create clean and comprehensive git commit message.
2. Convert 'git diff --staged' command it into a git commit message.
3. Use the present tense.
4. Lines limit to 50 characters.
5. Message should be one paragraph.
6. Message should be very clear and short.
7. Skip "Update X file" or "Update" at the beginning of the message and go straight to the point.
8. Respond using JSON.
9. JSON scheme {"commit_message": string}
`

var conventionalCommitPrompt = `
*** Conventional Commit ***
1. Use Conventional commit.
2. Conventional commit keywords: fix, feat, build, chore, ci, docs, style, refactor, perf, test.
	- feat: new feature
	- fix: bug fix
	- build: changes that affect the build system or external dependencies
	- ci: changes to our CI configuration files and scripts
	- chore: other changes that don't modify src or test files
	- docs: documentation only changes
	- style: changes that do not affect the meaning of the code (white-space, formatting, etc)
	- refactor: code change that neither fixes a bug nor adds feature
	- perf: code change that improves performance
	- test: adding missing tests or correcting existing tests
3. Do not preface the commit with anything. 
`

func createPrompt(diff, style string) string {
	if style == "conventional-commit" {
		basePrePrompt = basePrePrompt + conventionalCommitPrompt
	}

	return fmt.Sprintf("git diff: ```%s``` \n %s", diff, basePrePrompt)
}
