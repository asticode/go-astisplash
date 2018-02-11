package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/asticode/go-astilog"
	"github.com/asticode/go-bindata"
	"github.com/pkg/errors"
)

// makeWindows makes all proper steps for Windows
func makeWindows() (err error) {
	// Build
	var d TemplateData
	if d, err = buildWindows(); err != nil {
		err = errors.Wrap(err, "building for windows failed")
		return
	}

	// Execute template
	if err = executeTemplate(d, "./asset_windows.go"); err != nil {
		err = errors.Wrap(err, "executing template failed")
		return
	}
	return
}

// buildWindows builds the windows binary and returns the template data
func buildWindows() (d TemplateData, err error) {
	// Build
	astilog.Debug("Building")
	var cmd = exec.Command("i686-w64-mingw32-gcc", "-mwindows", "-o", "./splashmake/tmp/windows", "./splashmake/windows.c")
	cmd.Env = os.Environ()
	var b []byte
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
