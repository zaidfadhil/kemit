package engine

type Engine interface {
	GetCommit(diff string) (string, error)
}
