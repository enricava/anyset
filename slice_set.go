package anyset

// An implementation of Set with an underlying slice
type SliceSet[T comparable] struct {
	container *[]T
}

func (o SliceSet[T]) Add(item T) {
	*o.container = append(*o.container, item)
}

func (o SliceSet[T]) Remove(item T) {
	if len(*o.container) < 1 {
		return
	}

	newContainer := make([]T, len(*o.container)-1)
	for _, cur := range *o.container {
		if cur != item {
			newContainer = append(newContainer, cur)
		}
	}

	*o.container = newContainer
}

func (o SliceSet[T]) Contains(item T) (ok bool) {
	for _, cur := range *o.container {
		if cur == item {
			return true
		}
	}

	return false
}

func NewSliceSet[T comparable]() Set[T] {
	newContainer := make([]T, 0)
	return SliceSet[T]{
		container: &newContainer,
	}

}
