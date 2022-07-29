package ods

// List implements methods Size, Ste, Get, Add and Remove.
type List[T comparable] interface {
	Size() int
	Set(i int, x T)
	Get(i int) T
	Add(x T) T
	Remove(i int) T
}
