package commands

import (
	"os"
)

func fileExistsAndNotTty(name string) bool {
	fInfo, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return (fInfo.Mode() & os.ModeCharDevice) == 0 // ignore tty
}
