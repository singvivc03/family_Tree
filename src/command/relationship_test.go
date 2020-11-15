package command

import (
	"familyTree/src/data"
	"familyTree/src/family"
	"familyTree/src/person"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree family.Tree

func TestMain(m *testing.M) {
	tree = data.Build()
	m.Run()
}

func TestShouldReturnTheRelatedMembers(t *testing.T) {
	executeFunc := GetCommandFunc(AddChildCommand)
	message := executeFunc([]string{"Chitra", "Aria", person.Female}, tree)
	assert.Equal(t, message, "CHILD_ADDITION_SUCCEEDED")
	executeFunc = GetCommandFunc(GetRelationShipCommand)
	message = executeFunc([]string{"Lavnya", "Maternal-Aunt"}, tree)
	assert.Equal(t, message, "Aria")
}
