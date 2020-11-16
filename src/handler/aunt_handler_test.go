package handler

import (
	"familyTree/src/data"
	"familyTree/src/family"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree family.Tree

func TestMain(m *testing.M) {
	tree = data.Build()
	m.Run()
}

func TestShouldReturnMaternalAunt(t *testing.T) {
	yodhan, _ := tree.FindMemberByName("Yodhan")
	maternalAunt := maternalAuntHandler(yodhan, tree)
	assert.Equal(t, maternalAunt, "Tritha")
}

func TestShouldReturnEmptyMaternalAunt(t *testing.T) {
	vila, _ := tree.FindMemberByName("Vila")
	maternalAunt := maternalAuntHandler(vila, tree)
	assert.Equal(t, maternalAunt, "")
}

func TestShouldReturnPaternalAunt(t *testing.T) {
	dritha, _ := tree.FindMemberByName("Dritha")
	paternalAunt := paternalAuntHandler(dritha, tree)
	assert.Equal(t, paternalAunt, "Satya")
}

func TestShouldReturnEmptyPaternalAunt(t *testing.T) {
	shan, _ := tree.FindMemberByName("King Shan")
	paternalAunt := paternalAuntHandler(shan, tree)
	assert.Equal(t, paternalAunt, "")
}
