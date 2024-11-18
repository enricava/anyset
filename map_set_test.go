package anyset_test

import (
	"github.com/enricava/anyset"

	"testing"
)

func TestMapSetOneItem(t *testing.T) {
	// Call the generic verifier function
	VerifySetImplOneItem(anyset.NewMapSet, t)
}

func TestMapSetManyItems(t *testing.T) {
	// Call the generic verifier function
	VerifySetManyItems(anyset.NewMapSet, t)
}
