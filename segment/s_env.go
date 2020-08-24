package segment

import (
	"errors"
	"github.com/RomainGiraud/gompt/format"
	"os"
	"path/filepath"
)

// Create a Text segment.
// If environment variable env exists, a Text segment is loaded with the value returned from fn.
func NewCheckEnv(env string, fn func(string) string) (*Text, error) {
	return NewCheckEnvStylized(env, fn, format.NewStyleStandard(format.UniBrush{format.Default}, format.UniBrush{format.Default}))
}

// Create a Text segment with a style.
// If environment variable env exists, a Text segment is loaded with the value returned from fn.
func NewCheckEnvStylized(env string, fn func(string) string, style format.Style) (*Text, error) {
	path, ok := os.LookupEnv(env)
	if !ok {
		return nil, errors.New("Variable does not exist.")
	}
	return &Text{" " + fn(path) + " ", style}, nil
}

// Create a Text segment displaying direnv name.
func NewDirEnv() (*Text, error) {
	return NewDirEnvStylized(format.NewStyleStandard(format.UniBrush{format.Default}, format.UniBrush{format.Default}))
}

// Create a Text segment displaying direnv name with a style.
func NewDirEnvStylized(style format.Style) (*Text, error) {
	return NewCheckEnvStylized("DIRENV_DIR", filepath.Base, style)
}

// Create a Text segment displaying virtualenv name.
func NewVirtualEnv() (*Text, error) {
	return NewVirtualEnvStylized(format.NewStyleStandard(format.UniBrush{format.Default}, format.UniBrush{format.Default}))
}

// Create a Text segment displaying virtualenv name with a style.
func NewVirtualEnvStylized(style format.Style) (*Text, error) {
	return NewCheckEnvStylized("VIRTUAL_ENV", filepath.Base, style)
}
