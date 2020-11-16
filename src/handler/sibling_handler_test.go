package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSiblingsOfChit(t *testing.T) {
	chit, _ := tree.FindMemberByName("Chit")
	siblings := siblingHandler(chit, tree)
	assert.Equal(t, siblings, "Ish,Vich,Aras,Satya")
}
