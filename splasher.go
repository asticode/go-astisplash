package astisplash

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Splasher represents an object capable of displaying splash screens
type Splasher struct {
	binaryPath string
}

// New creates a new splasher
func New() (s *Splasher, err error) {
	// Init
	s = &Splasher{}

	// Get executable dir path
	var p string
	if p, err = os.Executable(); err != nil {
		err = errors.Wrap(err, "os.Executable failed")
		return
	}
	p = filepath.Dir(p)

	// Disembed assets
	astilog.Debugf("Disembedding to %s", p)
	if s.binaryPath, err = disembedAssets(p); err != nil {
		err = errors.Wrap(err, "disembedding assets failed")
		return
	}
	astilog.Debugf("Disembedded to %s", s.binaryPath)
	return
}

// Close closes the Splasher properly
func (s *Splasher) Close() (err error) {
	// Remove binary
	astilog.Debugf("Removing %s", s.binaryPath)
	if err = os.Remove(s.binaryPath); err != nil {
		err = errors.Wrapf(err, "remove of %s failed", s.binaryPath)
		return
	}
	return
}

// Splash displays a splash screen
func (s *Splasher) Splash(src string, width, height int) (*Splash, error) {
	return s.splash(exec.Command(s.binaryPath, "-b"+src, "-w"+strconv.Itoa(width), "-h"+strconv.Itoa(height)))
}

// splash executes a command and returns a *Splash
func (s *Splasher) splash(cmd *exec.Cmd) (sp *Splash, err error) {
	// Exec
	sp = &Splash{cmd: cmd}
	astilog.Debugf("Executing %s", strings.Join(cmd.Args, " "))
	if err = sp.cmd.Start(); err != nil {
		err = errors.Wrapf(err, "starting %s failed", s.binaryPath)
	}

	// Wait
	go cmd.Wait()
	return
}
