package engine

type Engine interface {
	GetCommitMessage(gitDiff, style string) (string, error)
}
