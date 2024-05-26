package engines

import "github.com/zaidfadhil/kemit.git/git"

var prePrompt = `
	- Your task is to create clean and comprehensive git commit message.
	- Explain WHAT were the changes and mainly WHY the changes were done. 
	- Convert 'git diff --staged' command it into a git commit message.
	- Use the present tense. 
	- Line limit to 100 chars. 
	- Respond using JSON. 
	- JSON scheme {"commit_message": string}
	
	git diff: 
`

func createPrompt(files []git.GitFile) string {
	return prePrompt
}
