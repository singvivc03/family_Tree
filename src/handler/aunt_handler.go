package handler

import (
	"familyTree/src/family"
	"familyTree/src/person"
	"strings"
)

func maternalAuntHandler(individual person.Person, tree family.Tree) string {
	siblings := getSiblings(individual.GetMother())
	maternalAunt := getSiblingsByGender(person.Female, siblings)
	return strings.Join(maternalAunt, ",")
}

func paternalAuntHandler(individual person.Person, tree family.Tree) string {
	siblings := getSiblings(individual.GetFather())
	paternalAunt := getSiblingsByGender(person.Female, siblings)
	return strings.Join(paternalAunt, ",")
}
