package env

import (
	"errors"
	"os"
)

type EnvVar struct {
	CompilerPath string
}

func Config() (EnvVar, error) {
	compilerPath := "./bin/sml"

	if compilerPath == "" {
		return EnvVar{}, errors.New("Error getting compiler path")
	}

	if _, err := os.Stat(compilerPath); os.IsNotExist(err) {
		return EnvVar{}, errors.New("Path '%s' does not exist.")
	}

	return EnvVar{
		CompilerPath: compilerPath,
	}, nil
}
