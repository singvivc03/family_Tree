package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnSisterInLaw(t *testing.T) {
	satya, _ := tree.FindMemberByName("Satya")
	inLaws := sisterInLawHandler(satya, tree)
	assert.Equal(t, inLaws, "Amba,Lika,Chitra")
}

func TestShouldReturnEmptySisterInLaw(t *testing.T) {
	vyan, _ := tree.FindMemberByName("Vyan")
	inLaws := sisterInLawHandler(vyan, tree)
	assert.Equal(t, inLaws, "")
}

func TestShouldReturnBrotherInLaw(t *testing.T) {
	vyan, _ := tree.FindMemberByName("Vyan")
	inLaws := brotherInLawHandler(vyan, tree)
	assert.Equal(t, inLaws, "Chit,Ish,Vich,Aras")
}

func TestShouldReturnEmptyBrotherInLaw(t *testing.T) {
	satya, _ := tree.FindMemberByName("Satya")
	inLaws := brotherInLawHandler(satya, tree)
	assert.Equal(t, inLaws, "")
}
