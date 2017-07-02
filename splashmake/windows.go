package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/asticode/go-astilog"
	"github.com/jteeuwen/go-bindata"
	"github.com/pkg/errors"
)

// DataWindows represents windows template data
type DataWindows struct {
	Binary string
	DLL    string
}

// makeWindows makes all proper steps for Windows
func makeWindows() (err error) {
	// Build
	var d DataWindows
	if d, err = buildWindows(); err != nil {
		err = errors.Wrap(err, "building for windows failed")
		return
	}

	// Execute template
	if err = executeTemplate(d, "/windows.tmpl", "./asset_windows.go"); err != nil {
		err = errors.Wrap(err, "executing template failed")
		return
	}
	return
}

// buildWindows builds the windows binary and returns the windows data
func buildWindows() (d DataWindows, err error) {
	// Update args
	var args = []string{"-o", "./splashmake/tmp/windows", "./splashmake/splash.c"}

	// Retrieve pkg-config
	astilog.Debug("Retrieving pkg-config")
	var cmd = exec.Command("pkg-config", "--cflags", "--libs", "gtk+-3.0")
	cmd.Env = append([]string{"PKG_CONFIG_PATH=/mnt/hgfs/shared/mingw32/lib/pkgconfig/"}, os.Environ()...)
	var b []byte
	if b, err = cmd.CombinedOutput(); err != nil {
		err = errors.Wrapf(err, "retrieving pkg-config failed with body %s", b)
		return
	}
	args = append(args, strings.Split(string(bytes.TrimSpace(b)), " ")...)

	// Build
	astilog.Debug("Building")
	cmd = exec.Command("i686-w64-mingw32-gcc", args...)
	if b, err = cmd.CombinedOutput(); err != nil {
		err = errors.Wrapf(err, "executing %s failed with output %s", strings.Join(cmd.Args, " "), b)
		return
	}

	// Read file
	astilog.Debug("Reading linux binary")
	if b, err = ioutil.ReadFile("./splashmake/tmp/windows"); err != nil {
		err = errors.Wrap(err, "reading \"./splashmake/tmp/windows\" failed")
		return
	}

	// Write
	var buf = &bytes.Buffer{}
	var w = &bindata.StringWriter{Writer: buf}
	w.Write(b)
	d.Binary = buf.String()
	return
}
