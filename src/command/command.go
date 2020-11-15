package command

import (
	"familyTree/src/family"
)

// Command defines the type of the given command
type Command string

// Following commands are available.
const (
	AddChildCommand        Command = "ADD_CHILD"
	GetRelationShipCommand Command = "GET_RELATIONSHIP"
)

var commandVsCommandFunc map[Command]Execute

func init() {
	commandVsCommandFunc = map[Command]Execute{
		AddChildCommand:        addChild,
		GetRelationShipCommand: getRelationShip,
	}
}

// Execute a function that returns the result depends on the operation it performs
type Execute func(args []string, tree family.Tree) string

// GetCommandFunc ...
func GetCommandFunc(command Command) Execute {
	return commandVsCommandFunc[command]
}
