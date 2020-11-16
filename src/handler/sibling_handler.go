package handler

import (
	"familyTree/src/family"
	"familyTree/src/person"
	"strings"
)

func siblingHandler(individual person.Person, tree family.Tree) string {
	siblings := individual.GetSiblings()
	var siblingNames []string
	for _, sibling := range siblings {
		siblingNames = append(siblingNames, sibling.GetName())
	}
	return strings.Join(siblingNames, ",")
}
