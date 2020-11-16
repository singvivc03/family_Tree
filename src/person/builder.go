package person

type builder struct {
	person *person
}

// Builder interface for building person
type Builder interface {
	SetName(name string) Builder
	SetGender(gender string) Builder
	SetMother(mother Person) Builder
	SetFather(father Person) Builder
	SetSpouse(spouse Person) Builder
	Build() Person
}

// NewPersonBuilder ...
func NewPersonBuilder() Builder {
	return &builder{person: &person{}}
}

func (b *builder) SetName(name string) Builder {
	b.person.name = name
	return b
}

func (b *builder) SetGender(gender string) Builder {
	b.person.gender = gender
	return b
}

func (b *builder) SetMother(mother Person) Builder {
	b.person.mother = mother
	return b
}

func (b *builder) SetFather(father Person) Builder {
	b.person.father = father
	return b
}

func (b *builder) SetSpouse(spouse Person) Builder {
	b.person.spouse = spouse
	return b
}

func (b *builder) Build() Person {
	return b.person
}
