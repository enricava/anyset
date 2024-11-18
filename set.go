package anyset

// A Set is any container that can Add or Remove items of any kind, and verify
// whether an item belongs to the Set
type Set[T comparable] interface {
	// Add an item to the Set
	Add(T)

	// Remove an item from the Set
	Remove(T)

	// Check if an item belongs to the Set
	Contains(T) bool
}

// A generic function type that returns a Set implementation of type T
type Create[T comparable] func() Set[T]

func AddMany[T comparable](set Set[T], items ...T) {
	for _, item := range items {
		set.Add(item)
	}
}

func RemoveMany[T comparable](set Set[T], items ...T) {
	for _, item := range items {
		set.Remove(item)
	}
}

func ContainsEvery[T comparable](set Set[T], items ...T) bool {
	for _, item := range items {
		if !set.Contains(item) {
			return false
		}
	}
	return true
}

func ContainsAny[T comparable](set Set[T], items ...T) bool {
	for _, item := range items {
		if set.Contains(item) {
			return true
		}
	}
	return false
}
