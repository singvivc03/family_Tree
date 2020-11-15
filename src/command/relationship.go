package command

import (
	"familyTree/src/family"
	"familyTree/src/handler"
)

func getRelationShip(args []string, tree family.Tree) string {
	name, relationType := args[0], args[1]
	person, err := tree.FindMemberByName(name)
	if err != nil {
		return err.Error()
	}
	handleFunc := handler.GetHandler(relationType)
	return handleFunc(person, tree)
}
