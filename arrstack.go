package ods

// ArrayStack implements the List interface using a "resizable" backing array. For the
// go impl we just use a slice with the same semantics we would use if this were written in C.
type ArrayStack[T comparable] struct {
	arr []T // we are pretending this is like a heap alloced pointer in C even though we know its a slice...
	n   int
}

// NewArrayStack initializes and returns an ArrayStack.
func NewArrayStack[T comparable]() ArrayStack[T] {
	arr := make([]T, 1) // using the [] len as if it was capacity
	n := 0              // size of the ArrayStack; should be <  len(arr)
	return ArrayStack[T]{arr, n}

}

// Get adds an item to the index i or returns an error if out of bounds.
func (a *ArrayStack[T]) Get(i int) T {
	return a.arr[i]
}

// Set replaces an item at position i and returns the item or an error if out of bounds.
func (a *ArrayStack[T]) Set(i int, x T) T {

	out := a.arr[i]
	a.arr[i] = x
	return out
}

// Size returns the Size of the ArrayStack
func (a *ArrayStack[T]) Size() int {
	return a.n
}

// Add inserts a new item x in the ArrayStack at position i, resizing if neccessary.
func (a *ArrayStack[T]) Add(i int, x T) {

	if a.n+1 > len(a.arr) { // resize if we hit capacity
		a.resize()
	}
	for j := a.n; j > i; j-- { // start from the end and shift right up to i (reverse iter)
		a.arr[j] = a.arr[j-1]
	}

	a.arr[i] = x
	a.n++

}

// Remove removes the item at position i.
func (a *ArrayStack[T]) Remove(i int) T {

	retVal := a.arr[i]
	for j := i; j < a.n-1; j++ { // start from middle and shift left to a.n (forw iter)
		a.arr[j] = a.arr[j+1]
	}
	a.n--
	if len(a.arr) >= 3*a.n { // check if we can shrink
		a.resize()
	}
	//fmt.Printf("%v\n", a.arr)
	return retVal
}

func (a *ArrayStack[T]) resize() {
	newArr := make([]T, odsMax(2*a.n, 1)) // this will both shrink or grow depending on n
	for i := 0; i < a.n; i++ {
		newArr[i] = a.arr[i]
	}
	a.arr = newArr
}

// for debugging
func (a *ArrayStack[T]) getBackingArr() []T {
	return a.arr
}
