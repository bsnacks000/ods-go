package ods

// An ArrayQueue FIFO queue using a backing array.
type ArrayQueue[T comparable] struct {
	arr []T
	n   int
	j   int
}

// NewArrayQueue initializes a new ArrayQueue
func NewArrayQueue[T comparable]() ArrayQueue[T] {
	arr := make([]T, 1)
	n := 0
	j := 0
	return ArrayQueue[T]{arr, n, j}
}

// Size returns the Size of the ArrayQueue.
func (a *ArrayQueue[T]) Size() int {
	return a.n
}

// Add adds the item to the top of the stack
func (a *ArrayQueue[T]) Add(x T) {
	a.maybeGrow()
	a.arr[(a.j+a.n)%len(a.arr)] = x
	a.n++
	//fmt.Printf("%v\n", a.arr)
}

// Remove removes the item at the top of the stack
func (a *ArrayQueue[T]) Remove() T {
	x := a.arr[a.j]
	a.j = (a.j + 1) % len(a.arr)
	a.n--
	a.maybeShrink()
	return x
}

func (a *ArrayQueue[T]) maybeGrow() {
	if a.n+1 > len(a.arr) {
		a.resize()
	}
}

func (a *ArrayQueue[T]) maybeShrink() {
	if a.n >= 3*a.n {
		a.resize()
	}
}

func (a *ArrayQueue[T]) resize() {
	newArr := make([]T, odsMax(2*a.n, 1)) // this will both shrink or grow depending on n

	for k := 0; k < a.n; k++ {
		newArr[k] = a.arr[(a.j+k)%len(a.arr)]
	}
	a.arr = newArr
	a.j = 0
}

// for debugging
func (a *ArrayQueue[T]) getBackingArr() []T {
	return a.arr
}
