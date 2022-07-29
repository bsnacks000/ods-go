package ods

// List implements methods Size, Ste, Get, Add and Remove.
type IList[T comparable] interface {
	Size() int
	Set(i int, x T)
	Get(i int) T
	Add(i int, x T)
	Remove(i int) T
}

type IStack[T comparable] interface {
	Size() int
	Add(x T)
	Remove() T
}
