package astisplash

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Splasher represents an object capable of displaying splash screens
type Splasher struct {
	binaryPath    string
	directoryPath string
}

// New creates a new splasher
func New() (s *Splasher, err error) {
	// Init
	s = &Splasher{directoryPath: filepath.Join(os.TempDir(), "astisplash")}

	// Remove directory
	if err = os.RemoveAll(s.directoryPath); err != nil {
		err = errors.Wrapf(err, "removall of %s failed", s.directoryPath)
		return
	}

	// Create directory
	if err = os.MkdirAll(s.directoryPath, 0755); err != nil {
		err = errors.Wrapf(err, "mkdirall of %s failed", s.directoryPath)
		return
	}

	// Disembed assets
	if s.binaryPath, err = disembedAssets(s.directoryPath); err != nil {
		err = errors.Wrap(err, "disembedding assets failed")
		return
	}
	return
}

// Close closes the Splasher properly
func (s *Splasher) Close() (err error) {
	// Remove temp directory
	if err = os.RemoveAll(s.directoryPath); err != nil {
		err = errors.Wrapf(err, "removeall of %s failed", s.directoryPath)
		return
	}
	return
}

// Splash displays a splash screen at a specific position of the screen
func (s *Splasher) Splash(src string) (*Splash, error) {
	return s.splash(exec.Command(s.binaryPath, "-i"+src, "-c"))
}

// Splash displays a splash screen at a specific position of the screen
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
