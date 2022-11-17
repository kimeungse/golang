package single

import (
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("emptydd name")
	}

	return name, nil
}
