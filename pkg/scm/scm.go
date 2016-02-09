package scm

type Info interface {
	Commit() string
	Branch() string
}
