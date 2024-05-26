package engines

type Engine interface {
	GetCommit() (string, error)
}
