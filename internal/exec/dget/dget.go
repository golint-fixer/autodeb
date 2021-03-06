package dget

import (
	"os/exec"

	"salsa.debian.org/autodeb-team/autodeb/internal/errors"
)

//Dget the URL in the directory
func Dget(url, directory string) error {
	command := exec.Command(
		"dget",
		url,
	)
	command.Dir = directory

	if combinedOutput, err := command.CombinedOutput(); err != nil {
		return errors.WithMessagef(err, "dget error: %s", combinedOutput)
	}

	return nil
}
