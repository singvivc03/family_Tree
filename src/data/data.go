package data

import (
	"familyTree/src/family"
	"familyTree/src/person"
)

type peopleStub struct {
	name   string
	gender string
	father string
	mother string
	spouse string
}

var people = []peopleStub{
	{"King Shan", "Male", "", "", "Queen Anga"},
	{"Chit", "Male", "King Shan", "Queen Anga", "Amba"},
	{"Ish", "Male", "King Shan", "Queen Anga", ""},
	{"Vich", "Male", "King Shan", "Queen Anga", "Lika"},
	{"Aras", "Male", "King Shan", "Queen Anga", "Chitra"},
	{"Satya", "Female", "King Shan", "Queen Anga", "Vyan"},
	{"Dritha", "Female", "Chit", "Amba", "Jaya"},
	{"Tritha", "Female", "Chit", "Amba", ""},
	{"Vritha", "Male", "Chit", "Amba", ""},
	{"Yodhan", "Male", "Jaya", "Dritha", ""},
	{"Vila", "Female", "Vich", "Lika", ""},
	{"Chika", "Female", "Vich", "Lika", ""},
	{"Jnki", "Female", "Aras", "Chitra", "Arit"},
	{"Ahit", "Male", "Aras", "Chitra", ""},
	{"Laki", "Male", "Arit", "Jnki", ""},
	{"Lavnya", "Male", "Arit", "Jnki", ""},
	{"Asva", "Male", "Vyan", "Satya", "Satvy"},
	{"Vyas", "Male", "Vyan", "Satya", "Krpi"},
	{"Atya", "Female", "Vyan", "Satya", ""},
	{"Vasa", "Male", "Asva", "Satvy", ""},
	{"Kriya", "Male", "Vyas", "Krpi", ""},
	{"Krithi", "Female", "Vyas", "Krpi", ""},
}

// Build ...
func Build() family.Tree {
	dummyTree := family.New()
	for _, individual := range people {
		member := person.NewPersonBuilder().SetName(individual.name).SetGender(individual.gender).Build()
		if individual.mother != "" {
			mother, err := dummyTree.FindMemberByName(individual.mother)
			if err != nil {
				panic(err.Error())
			}
			member = member.ToBuilder().SetMother(mother).Build()
		}
		if individual.father != "" {
			father, err := dummyTree.FindMemberByName(individual.father)
			if err != nil {
				panic(err.Error())
			}
			member = member.ToBuilder().SetFather(father).Build()
		}
		if individual.spouse != "" {
			spouseGender := person.Male
			if individual.gender == person.Male {
				spouseGender = person.Female
			}
			spouse := person.NewPerson(individual.spouse, spouseGender)
			member = member.ToBuilder().SetSpouse(spouse).Build()
		}
		dummyTree.Add(member)
	}
	return dummyTree
}
