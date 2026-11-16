package greeting

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyName          = errors.New("name is empty")
	ErrNameNotInFirstHalf = errors.New("name is not in the first half of the alphabet")
)

func ValidateName(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	first := name[0]
	if !(first >= 'a' && first <= 'm') && !(first >= 'A' && first <= 'M') {
		return fmt.Errorf("%w: %q", ErrNameNotInFirstHalf, name)
	}
	return nil
}
