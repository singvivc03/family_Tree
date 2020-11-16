package handler

import (
	"familyTree/src/family"
	"familyTree/src/person"
	"strings"
)

func sisterInLawHandler(individual person.Person, tree family.Tree) string {
	siblings := individual.GetSiblings()
	inLaws := getSpouseOfSiblings(siblings, person.Male)
	spouseSiblings := getSpouseSibling(individual, person.Female)
	for _, sibling := range spouseSiblings {
		inLaws = append(inLaws, sibling)
	}
	return strings.Join(inLaws, ",")
}

func brotherInLawHandler(individual person.Person, tree family.Tree) string {
	siblings := individual.GetSiblings()
	inLaws := getSpouseOfSiblings(siblings, person.Female)
	spouseSiblings := getSpouseSibling(individual, person.Male)
	for _, sibling := range spouseSiblings {
		inLaws = append(inLaws, sibling)
	}
	return strings.Join(inLaws, ",")
}

func getSpouseSibling(individual person.Person, gender string) []string {
	var inLaws []string
	if individual.GetSpouse() != nil {
		spouseSiblings := individual.GetSpouse().GetSiblings()
		for _, spouseSibling := range spouseSiblings {
			if spouseSibling.GetGender() == gender {
				inLaws = append(inLaws, spouseSibling.GetName())
			}
		}
	}
	return inLaws
}

func getSpouseOfSiblings(siblings []person.Person, gender string) []string {
	var inLaws []string
	// husband of siblings
	for _, sibling := range siblings {
		if sibling.GetGender() == gender && sibling.GetSpouse() != nil {
			inLaws = append(inLaws, sibling.GetSpouse().GetName())
		}
	}
	return inLaws
}
