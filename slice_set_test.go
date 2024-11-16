package anyset_test

import (
	"github.com/enricava/anyset"

	"testing"
)

func TestSliceSetOneItem(t *testing.T) {
	// Call the generic verifier function
	VerifySetImplOneItem[anyset.SliceSet[int]](anyset.NewSliceSet[int], t)
}

func TestSliceSetManyItems(t *testing.T) {
	// Call the generic verifier function
	VerifySetManyItems[anyset.SliceSet[int]](anyset.NewSliceSet[int], t)
}
