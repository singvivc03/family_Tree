package handler

import (
	"familyTree/src/family"
	"familyTree/src/person"
)

var allHandlers map[string]Handle

func init() {
	allHandlers = map[string]Handle{
		"Maternal-Aunt":  maternalAuntHandler,
		"Paternal-Aunt":  paternalAuntHandler,
		"Paternal-Uncle": paternalUncleHandler,
		"Maternal-Uncle": maternalUncleHandler,
	}
}

// Handle the relation type
type Handle func(person person.Person, tree family.Tree) string

// GetHandler returns the handler type
func GetHandler(relationType string) Handle {
	return allHandlers[relationType]
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
