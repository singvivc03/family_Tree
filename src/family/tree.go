package family

import (
	"errors"
	"familyTree/src/person"
)

// errors used when operations on tree has any failure
var (
	errPersonNotFound     = errors.New("PERSON_NOT_FOUND")
	errChildAddtionFailed = errors.New("CHILD_ADDITION_FAILED")
	errDuplicateName      = errors.New("Person with the same name already exists")
)

type tree struct {
	people map[string]person.Person
}

// Tree ...
type Tree interface {
	Add(person person.Person) error
	AddChild(motherName string, child person.Person) error
	FindMemberByName(name string) (person.Person, error)
}

// New creates a new family tree
func New() Tree {
	return &tree{
		people: make(map[string]person.Person),
	}
}

func (t *tree) Add(person person.Person) error {
	if _, ok := t.people[person.GetName()]; ok {
		return errDuplicateName
	}
	if person.GetSpouse() != nil {
		spouse := person.GetSpouse().ToBuilder().SetSpouse(person).Build()
		t.people[spouse.GetName()] = spouse
		person = person.ToBuilder().SetSpouse(spouse).Build()
	}
	t.people[person.GetName()] = person
	return nil
}

func (t *tree) AddChild(motherName string, child person.Person) error {
	mother, err := t.FindMemberByName(motherName)
	if err != nil {
		return err
	}
	if mother.GetGender() == person.Male {
		return errChildAddtionFailed
	}
	child = child.ToBuilder().SetMother(mother).SetFather(mother.GetSpouse()).Build()
	mother.AddChild(child)
	mother.GetSpouse().AddChild(child)
	return t.Add(child)
}

func (t *tree) FindMemberByName(name string) (person.Person, error) {
	if _, ok := t.people[name]; !ok {
		return nil, errPersonNotFound
	}
	return t.people[name], nil
}
