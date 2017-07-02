package astisplash

import (
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
)

// Splash represents a splash screen
type Splash struct {
	cmd *exec.Cmd
}

// Close closes the splash screen properly
func (s *Splash) Close() (err error) {
	if s.cmd != nil && s.cmd.Process != nil {
		if err = s.cmd.Process.Signal(syscall.SIGINT); err != nil {
			err = errors.Wrapf(err, "sending SIGINT to process %d failed", s.cmd.Process.Pid)
			return
		}
	}
	return
}
