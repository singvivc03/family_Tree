package family

import (
	"familyTree/src/person"
	"testing"

	"github.com/stretchr/testify/assert"
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
	{"Satya", "Femal", "King Shan", "Queen Anga", "Vyan"},
}

var dummyTree = New()

func TestMain(m *testing.M) {
	for _, individual := range people {
		member := person.NewPersonBuilder().SetName(individual.name).SetGender(individual.gender).Build()
		if individual.mother != "" {
			mother := person.NewPerson(individual.mother, person.Female)
			member = member.ToBuilder().SetMother(mother).Build()
		}
		if individual.father != "" {
			father := person.NewPerson(individual.father, person.Male)
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
	m.Run()
}

func TestShouldFindMemberFromTree(t *testing.T) {
	person, _ := dummyTree.FindMemberByName("Vich")
	assert.NotNil(t, person)
	assert.Equal(t, person.GetName(), "Vich")
	assert.Equal(t, person.GetGender(), "Male")
	assert.Equal(t, person.GetFather().GetName(), "King Shan")
	assert.Equal(t, person.GetMother().GetName(), "Queen Anga")
}

func TestShouldAddChildToTheTree(t *testing.T) {
	member := person.NewPersonBuilder().SetName("Vila").SetGender("Female").Build()
	dummyTree.AddChild("Lika", member)
	person, _ := dummyTree.FindMemberByName("Vila")
	assert.Equal(t, person.GetName(), "Vila")
	assert.Equal(t, person.GetGender(), "Female")
	assert.Nil(t, person.GetSpouse())
	assert.Equal(t, person.GetFather().GetName(), "Vich")
	assert.Equal(t, person.GetMother().GetName(), "Lika")
}

func TestShouldVerifySameNameIsNotAddedMulitpleTimes(t *testing.T) {
	member := person.NewPersonBuilder().SetName("Chit").SetGender("Male").Build()
	err := dummyTree.Add(member)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Person with the same name already exists")
}

func TestShouldReturnErrorWhenPersonIsNotFound(t *testing.T) {
	member := person.NewPersonBuilder().SetName("Test").SetGender("Male").Build()
	err := dummyTree.AddChild("dummy", member)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "PERSON_NOT_FOUND")
}

func TestShouldReturnChildAdditionFailed(t *testing.T) {
	member := person.NewPersonBuilder().SetName("Test").SetGender("Male").Build()
	err := dummyTree.AddChild("King Shan", member)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "CHILD_ADDITION_FAILED")
}
