package resolveString

type StringResolver interface {
	Resolve() (string, error)
}
