package command

import (
	"familyTree/src/family"
	"familyTree/src/person"
)

func addChild(args []string, tree family.Tree) string {
	motherName, childName, childGender := args[0], args[1], args[2]
	err := tree.AddChild(motherName, person.NewPerson(childName, childGender))
	if err != nil {
		return err.Error()
	}
	return "CHILD_ADDITION_SUCCEEDED"
}
