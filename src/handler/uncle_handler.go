package handler

import (
	"familyTree/src/family"
	"familyTree/src/person"
	"strings"
)

func paternalUncleHandler(individual person.Person, tree family.Tree) string {
	siblings := getSiblings(individual.GetFather())
	paternalUncle := getSiblingsByGender(person.Male, siblings)
	return strings.Join(paternalUncle, ",")
}

func maternalUncleHandler(individual person.Person, tree family.Tree) string {
	siblings := getSiblings(individual.GetMother())
	maternalUncle := getSiblingsByGender(person.Male, siblings)
	return strings.Join(maternalUncle, ",")
}
