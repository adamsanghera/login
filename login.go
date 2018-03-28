package login

import (
	"github.com/adamsanghera/category"
)

// Login is a unified interface for managing Categories, Badges, and Authenticators.
// Categories are like nation-states, which have citizens (users).
// Badges are passports, created for and taken from citizens (users).
// Authenticators are like barcode scanners, used to validate badges and citizenship.
type Login interface {
	// Category Commands
	AddCategory(cid category.Category) error
	RemoveCategory(cid category.Category) error
	GrantMembership(cid category.Category, uid interface{}) error
	RevokeMembership(cid category.Category, uid interface{}) error
	IsMember(cid category.Category, uid interface{}) (bool, error)

	// Commands on people, irrespective of categories
	MakeUser(uid interface{}) error
	RemoveUser(uid interface{}, challenge interface{}) error
	AddAuth(uid interface{}, challenge interface{}) error
	RemoveAuth(uid interface{}, challenge interface{}) error

	// Commands regarding badges
	ReturnBadge(string) error

	// Queries regarding badges
	ValidateBadge(bid interface{}) (bool, error)
	MintBadge(uid interface{}, challenge interface{}, cid category.Category) (string, error)

	// Queries on categories
	CategoryExists(cid category.Category) (bool, error)
}
