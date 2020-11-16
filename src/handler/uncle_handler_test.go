package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnPaternalUncle(t *testing.T) {
	dritha, _ := tree.FindMemberByName("Dritha")
	paternalUncle := paternalUncleHandler(dritha, tree)
	assert.Equal(t, paternalUncle, "Ish,Vich,Aras")
}

func TestShouldReturnEmptyPaternalUncle(t *testing.T) {
	yodhan, _ := tree.FindMemberByName("Yodhan")
	paternalUncle := paternalUncleHandler(yodhan, tree)
	assert.Equal(t, paternalUncle, "")
}

func TestShouldReturnMaternalUncle(t *testing.T) {
	yodhan, _ := tree.FindMemberByName("Yodhan")
	maternalUncle := maternalUncleHandler(yodhan, tree)
	assert.Equal(t, maternalUncle, "Vritha")
}

func TestShouldReturnEmptyMaternalUncle(t *testing.T) {
	shan, _ := tree.FindMemberByName("King Shan")
	maternalUncle := maternalUncleHandler(shan, tree)
	assert.Equal(t, maternalUncle, "")
}
