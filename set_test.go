package anyset_test

import (
	"github.com/enricava/anyset"

	"testing"
)

// Any comparable type will do for verification, so we use int for the sake of
// simplicity

// Verifies that the generic Set can Add and Remove one item.
//
// Users of this function must provide an initializer for their Set
// implementation.
//
// Tests the set for Contains after every operation.
func VerifySetImplOneItem[S anyset.Set[int]](createSet anyset.Create[int, S], t *testing.T) {
	var set S
	set = createSet()

	set.Add(1)

	if !set.Contains(1) {
		t.Fail()
	}

	set.Remove(1)

	if set.Contains(1) {
		t.Fail()
	}
}

// Verifies that the generic Set can Add and Remove more than one item.
//
// Users of this function must provide an initializer for their Set
// implementation.
//
// Tests the set for Contains after every operation.
func VerifySetManyItems[S anyset.Set[int]](createSet anyset.Create[int, S], t *testing.T) {
	var set S
	set = createSet()

	anyset.AddMany(set, 1, 2, 3, 4)

	if !anyset.ContainsEvery(set, 1, 2, 3, 4) {
		t.Fail()
	}

	anyset.RemoveMany(set, 1, 4)

	if anyset.ContainsAny(set, 1, 4) {
		t.Fail()
	}

	anyset.RemoveMany(set, 2, 3)

	if anyset.ContainsAny(set, 1, 2, 3, 4) {
		t.Fail()
	}
}
