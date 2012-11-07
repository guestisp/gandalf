package repository

import (
	"fmt"
	"github.com/globocom/config"
	"os/exec"
	"path"
)

var bare string

func bareLocation() string {
	if bare != "" {
		return bare
	}
	var err error
	bare, err = config.GetString("bare-location")
	if err != nil {
		panic("You should configure a bare-location for gandalf.")
	}
	return bare
}

func newBare(name string) error {
	args := []string{"init", path.Join(bareLocation(), formatName(name)), "--bare"}
	if bareTempl, err := config.GetString("bare-template"); err == nil {
		args = append(args, "--template="+bareTempl)
	}
	cmd := exec.Command("git", args...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Could not create git bare repository: %s", err)
	}
	return nil
}

func removeBare(name string) error {
	err := filesystem().RemoveAll(path.Join(bareLocation(), formatName(name)))
	if err != nil {
		return fmt.Errorf("Could not remove git bare repository: %s", err)
	}
	return nil
}

func formatName(name string) string {
	return fmt.Sprintf("%s.git", name)
}
