package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonBuilder(t *testing.T) {
	personBuilder := NewPersonBuilder()
	person := personBuilder.SetName("Test").SetGender("Male").
		Build()
	assert.Equal(t, person.GetName(), "Test")
	assert.Equal(t, person.GetGender(), "Male")
}

func TestShouldVerifyExistingPersonIsModified(t *testing.T) {
	personBuilder := NewPersonBuilder()
	person := personBuilder.SetName("Test").SetGender("Male").
		Build()
	spouseTest := NewPerson("Test Spouse", "Female")
	person = person.ToBuilder().SetSpouse(spouseTest).Build()
	assert.Equal(t, person.GetName(), "Test")
	assert.Equal(t, person.GetGender(), "Male")
	assert.Equal(t, person.GetSpouse().GetName(), "Test Spouse")
	assert.Equal(t, person.GetSpouse().GetGender(), "Female")
}

func TestShouldVerifyMotherHasAChild(t *testing.T) {
	personBuilder := NewPersonBuilder()
	mother := NewPerson("Test Mother", "Female")
	person := personBuilder.SetName("Test").SetGender("Male").SetMother(mother).Build()
	assert.NotNil(t, person.GetMother())
	assert.Equal(t, person.GetMother().GetName(), "Test Mother")
	assert.Equal(t, len(person.GetMother().GetChildren()), 1)
}

func TestShouldVerifyFatherHasAChild(t *testing.T) {
	personBuilder := NewPersonBuilder()
	father := NewPerson("Test Father", "Male")
	person := personBuilder.SetName("Test").SetGender("Male").SetFather(father).Build()
	assert.NotNil(t, person.GetFather())
	assert.Equal(t, person.GetFather().GetName(), "Test Father")
	assert.Equal(t, len(person.GetFather().GetChildren()), 1)
}
