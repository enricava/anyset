package anyset

// An implementation of Set using an underlying map
//
// For symmetry with SliceSet, it is a struct enveloping its underlying
// container
type MapSet[T comparable] struct {
	container map[T]bool
}

func (o MapSet[T]) Add(item T) {
	o.container[item] = true
}

func (o MapSet[T]) Remove(item T) {
	delete(o.container, item)
}

func (o MapSet[T]) Contains(item T) (ok bool) {
	_, ok = o.container[item]
	return
}

func NewMapSet[T comparable]() Set[T] {
	return MapSet[T]{
		container: make(map[T]bool),
	}
}
