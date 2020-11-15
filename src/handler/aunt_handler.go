package handler

import (
	"familyTree/src/family"
	"familyTree/src/person"
	"strings"
)

func maternalAuntHandler(individual person.Person, tree family.Tree) string {
	siblings := individual.GetMother().GetSiblings()
	var maternalAunt []string
	for _, sibling := range siblings {
		if sibling.GetGender() == person.Female {
			maternalAunt = append(maternalAunt, sibling.GetName())
		}
	}
	return strings.Join(maternalAunt, ",")
}
