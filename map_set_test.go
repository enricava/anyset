package anyset_test

import (
	"github.com/enricava/anyset"

	"testing"
)

func TestMapSetOneItem(t *testing.T) {
	// Call the generic verifier function
	VerifySetImplOneItem[anyset.MapSet[int]](anyset.NewMapSet[int], t)
}

func TestMapSetManyItems(t *testing.T) {
	// Call the generic verifier function
	VerifySetManyItems[anyset.MapSet[int]](anyset.NewMapSet[int], t)
}
