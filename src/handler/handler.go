package handler

import (
	"familyTree/src/family"
	"familyTree/src/person"
)

var allHandlers map[string]Handle

func init() {
	allHandlers = map[string]Handle{
		"Maternal-Aunt": maternalAuntHandler,
		"Paternal-Aunt": paternalAuntHandler,
	}
}

// Handle the relation type
type Handle func(person person.Person, tree family.Tree) string

// GetHandler returns the handler type
func GetHandler(relationType string) Handle {
	return allHandlers[relationType]
}
