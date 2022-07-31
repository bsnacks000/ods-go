package ods

// An ArrayDeque implements a Deque with a resizable backing array
type ArrayDeque[T comparable] struct {
	arr []T
	n   int
	j   int
}

// NewArrayDeque initializes a new ArrayDeque
func NewArrayDeque[T comparable]() ArrayDeque[T] {
	arr := make([]T, 1)
	n := 0
	j := 0
	return ArrayDeque[T]{arr, n, j}
}

// Size returns the Size of the ArrayDeque.
func (a *ArrayDeque[T]) Size() int {
	return a.n
}

func (a *ArrayDeque[T]) Get(i int) T {
	return a.arr[(a.j+i)%len(a.arr)]
}

func (a *ArrayDeque[T]) Set(i int, x T) T {
	y := a.arr[(a.j+i)%len(a.arr)]
	a.arr[(a.j+i)%len(a.arr)] = x
	return y

}

func (a *ArrayDeque[T]) Add(i int, x T) {
	a.maybeGrow()
	if i < a.n/2 { // left shift
		if a.j == 0 {
			a.j = len(a.arr) - 1
		} else {
			a.j--
		}

		for k := 0; k <= i-1; k++ {
			a.arr[(a.j+k)%len(a.arr)] = a.arr[(a.j+k+1)%len(a.arr)]
		}

	} else { // right shift
		for k := a.n; k > i; k-- {
			a.arr[(a.j+k)%len(a.arr)] = a.arr[(a.j+k-1)%len(a.arr)]
		}
	}
	a.arr[(a.j+i)%len(a.arr)] = x
	a.n++
}

func (a *ArrayDeque[T]) Remove(i int) T {
	x := a.arr[(a.j+i)%len(a.arr)]

	if i < a.n/2 { // shift right
		for k := i; k > 0; k-- {
			a.arr[(a.j+k)%len(a.arr)] = a.arr[(a.j+k-1)%len(a.arr)]
		}
		a.j = (a.j + 1) % len(a.arr)
	} else { // shift left
		for k := i; k < a.n-1; k++ {
			a.arr[(a.j+k)%len(a.arr)] = a.arr[(a.j+k+1)%len(a.arr)]
		}
	}
	a.n--
	a.maybeShrink()
	return x
}

func (a *ArrayDeque[T]) maybeGrow() {
	if a.n+1 > len(a.arr) {
		a.resize()
	}
}

func (a *ArrayDeque[T]) maybeShrink() {
	if a.n >= 3*a.n {
		a.resize()
	}
}

func (a *ArrayDeque[T]) resize() {
	newArr := make([]T, odsMax(2*a.n, 1)) // this will both shrink or grow depending on n

	for k := 0; k < a.n; k++ {
		newArr[k] = a.arr[(a.j+k)%len(a.arr)]
	}
	a.arr = newArr
	a.j = 0
}
