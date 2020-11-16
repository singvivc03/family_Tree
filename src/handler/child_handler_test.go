package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnSonOfShan(t *testing.T) {
	shan, _ := tree.FindMemberByName("King Shan")
	sons := sonHandler(shan, tree)
	assert.Equal(t, sons, "Chit,Ish,Vich,Aras")
}

func TestShouldReturnEmptyWhenParentHasNoSon(t *testing.T) {
	ish, _ := tree.FindMemberByName("Ish")
	sons := sonHandler(ish, tree)
	assert.Equal(t, sons, "")
}

func TestShouldReturnDaughterOfShan(t *testing.T) {
	shan, _ := tree.FindMemberByName("King Shan")
	daughter := daughterHandler(shan, tree)
	assert.Equal(t, daughter, "Satya")
}

func TestShouldReturnEmptyWhenParentHasNoDaughter(t *testing.T) {
	dritha, _ := tree.FindMemberByName("Dritha")
	daughter := daughterHandler(dritha, tree)
	assert.Equal(t, daughter, "")
}
