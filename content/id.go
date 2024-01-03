// Package content provides functions common to several content types.
package content

// Identifier provides access to something that has a unique identity.
type Identifier interface {
	ID() string
}
