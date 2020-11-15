package person

// Gender constants
const (
	Male   string = "Male"
	Female string = "Female"
)

type person struct {
	name     string
	gender   string
	spouse   Person
	mother   Person
	father   Person
	children []Person
}

// Person ...
type Person interface {
	GetName() string
	GetGender() string
	GetMother() Person
	GetFather() Person
	GetSpouse() Person
	AddChild(child Person)
	GetChildren() []Person
	GetSiblings() []Person
	ToBuilder() Builder
}

// NewPerson creates new person
func NewPerson(name, gender string) Person {
	return &person{
		name:     name,
		gender:   gender,
		spouse:   nil,
		mother:   nil,
		father:   nil,
		children: []Person{},
	}
}

func toBuilder(existingPerson Person) Builder {
	p := &person{
		name:     existingPerson.GetName(),
		gender:   existingPerson.GetGender(),
		spouse:   existingPerson.GetSpouse(),
		mother:   existingPerson.GetMother(),
		father:   existingPerson.GetFather(),
		children: existingPerson.GetChildren(),
	}
	return &builder{person: p}
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) GetGender() string {
	return p.gender
}

func (p *person) GetMother() Person {
	return p.mother
}

func (p *person) GetFather() Person {
	return p.father
}

func (p *person) GetSpouse() Person {
	return p.spouse
}

func (p *person) GetChildren() []Person {
	return p.children
}

func (p *person) AddChild(person Person) {
	for _, child := range p.children {
		if child.GetName() == person.GetName() {
			return
		}
	}
	p.children = append(p.children, person)
}

func (p *person) GetSiblings() []Person {
	var children, siblings []Person
	if p.father != nil {
		children = p.father.GetChildren()
	}
	if p.mother != nil {
		children = p.mother.GetChildren()
	}
	for _, child := range children {
		if child.GetName() == p.name {
			continue
		}
		siblings = append(siblings, child)
	}
	return siblings
}

func (p *person) ToBuilder() Builder {
	return toBuilder(p)
}
