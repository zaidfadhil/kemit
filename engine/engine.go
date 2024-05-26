package engine

type Engine interface {
	GetCommit() (string, error)
}
