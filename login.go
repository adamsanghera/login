package login

import (
	"time"
)

// Authenticator is an interface for services providing authentication.
// It is used by Badge as a badge-scanner as a passport database.
// It is used by Login as a identity-affirming database.
type Authenticator interface {
	// Commands
	AddAuth(id interface{}, challenge interface{}) error
	RemoveAuth(id interface{}, challenge interface{}) error

	// Queries
	Auth(id interface{}, challenge interface{}) (bool, error)
	IDExists(id interface{}) (bool error)
}

// Badge is an interface for services providing identity insurance.
// Badges are a passport, certifying a recent check of citizenship.
type Badge interface {
	// Commands
	Revoke() error

	// Queries
	BelongsTo() Category
	Create() (uid interface{}, bid interface{}, badgeScanner Authenticator, timeToLive time.Duration, err error)
	Refresh() (uid interface{}, bid interface{}, badgeScanner Authenticator, timeToLive time.Duration, err error)
	Validate(uid interface{}, bid interface{}, badgeScanner Authenticator) (bool, error)
}

// Category is an interface for forming users into groups.
// Category is the nation, which users are citizens of.
// Membership in a category is maintained by a sorted list.
type Category interface {
	// Commands
	AddUser(uid interface{}) error
	RemoveUser(uid interface{}) error

	Contains(uid interface{}) (bool, error)
}

// Login is a unified interface for managing Categories, Badges, and Authenticators.
// Categories are like nation-states, which have citizens (users).
// Badges are passports, created for and taken from citizens (users).
// Authenticators are like barcode scanners, used to validate badges and citizenship.
type Login interface {
	// Category Commands
	AddCategory(cid interface{}) error
	RemoveCategory(cid interface{}) error
	GrantMembership(cid interface{}, uid interface{}) error
	RevokeMembership(cid interface{}, uid interface{}) error
	IsMember(cid interface{}, uid interface{}) (bool, error)

	// Commands on people, irrespective of categories
	MakeUser(uid interface{}) error
	RemoveUser(uid interface{}, challenge interface{}) error
	AddAuth(uid interface{}, challenge interface{}, auth Authenticator) error
	RemoveAuth(uid interface{}, challenge interface{}, auth Authenticator) error

	// Commands regarding badges
	ReturnBadge(Badge) error

	// Queries regarding badges
	ValidateBadge(bid interface{}) (bool, error)
	MintBadge(uid interface{}, challenge interface{}, cid interface{}) (Badge, error)

	// Queries on categories
	CategoryExists(cid interface{}) (bool, error)
}
