package anyset_test

import (
	"github.com/enricava/anyset"

	"testing"
)

func TestSliceSetOneItem(t *testing.T) {
	// Call the generic verifier function
	VerifySetImplOneItem(anyset.NewSliceSet, t)
}

func TestSliceSetManyItems(t *testing.T) {
	// Call the generic verifier function
	VerifySetManyItems(anyset.NewSliceSet, t)
}
