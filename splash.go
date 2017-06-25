package astisplash

import (
	"os/exec"

	"github.com/pkg/errors"
)

// Splash represents a splash screen
type Splash struct {
	cmd *exec.Cmd
}

// Close closes the splash screen properly
func (s *Splash) Close() (err error) {
	if s.cmd != nil && s.cmd.Process != nil {
		if err = s.cmd.Process.Kill(); err != nil {
			err = errors.Wrapf(err, "killing process %d failed", s.cmd.Process.Pid)
			return
		}
	}
	return
}
