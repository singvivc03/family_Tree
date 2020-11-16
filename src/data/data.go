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
		member := person.NewPersonBuilder().SetName(individual.name).SetGender(individual.gender)
		member = updateMemberWithSpouse(individual.spouse, individual.gender, member)
		person := updateMemberWithParent(individual.father, individual.mother, member, dummyTree)
		dummyTree.Add(person)
	}
	return dummyTree
}

func updateMemberWithSpouse(spouseName string, spouseGender string, memberBuilder person.Builder) person.Builder {
	if len(spouseName) > 0 {
		gender := person.Male
		if spouseGender == person.Male {
			gender = person.Female
		}
		memberBuilder = memberBuilder.SetSpouse(person.NewPerson(spouseName, gender))
	}
	return memberBuilder
}

func updateMemberWithParent(fName string, mName string, memberBuilder person.Builder, dummyTree family.Tree) person.Person {
	var father, mother person.Person
	var err error
	if len(fName) > 0 {
		father, err = dummyTree.FindMemberByName(fName)
		if err != nil {
			panic(err.Error())
		}
	}
	if len(mName) > 0 {
		mother, err = dummyTree.FindMemberByName(mName)
		if err != nil {
			panic(err.Error())
		}
	}
	person := memberBuilder.SetMother(mother).SetFather(father).Build()
	if father != nil && mother != nil {
		father.AddChild(person)
		mother.AddChild(person)
	}
	return person
}
