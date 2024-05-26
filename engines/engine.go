package engines

import "github.com/zaidfadhil/kemit.git/git"

type Engine interface {
	GetCommit(files []git.GitFile) (string, error)
}
