package handler

import (
	"familyTree/src/family"
	"familyTree/src/person"
	"strings"
)

func sonHandler(individual person.Person, tree family.Tree) string {
	sons := getChild(individual.GetChildren(), person.Male)
	return strings.Join(sons, ",")
}

func daughterHandler(individual person.Person, tree family.Tree) string {
	daughters := getChild(individual.GetChildren(), person.Female)
	return strings.Join(daughters, ",")
}

func getChild(members []person.Person, gender string) []string {
	var child []string
	for _, member := range members {
		if member.GetGender() == gender {
			child = append(child, member.GetName())
		}
	}
	return child
}
