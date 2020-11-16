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

func getSiblings(individual person.Person) []person.Person {
	var siblings []person.Person
	if individual != nil {
		siblings = individual.GetSiblings()
	}
	return siblings
}

func getSiblingsByGender(gender string, siblings []person.Person) []string {
	var aunt []string
	for _, sibling := range siblings {
		if sibling.GetGender() == gender {
			aunt = append(aunt, sibling.GetName())
		}
	}
	return aunt
}
