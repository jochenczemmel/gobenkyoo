package csvimport

import "errors"

// ErrNoKey represents missing keys in a format object.
var ErrNoKey = errors.New("no keys defined")

// InvalidKeyError represents the error case that a key
// is not valid.
type InvalidKeyError string

func (e InvalidKeyError) Error() string {
	return "invalid key: " + string(e)
}
