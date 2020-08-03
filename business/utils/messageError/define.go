package ferror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "room") {
		return errors.New("Room already taken")
	}

	return errors.New("Incorrect Details")
}
