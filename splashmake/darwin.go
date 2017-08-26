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

// makeDarwin makes all proper steps for Darwin
func makeDarwin() (err error) {
	// Build
	var d TemplateData
	if d, err = buildDarwin(); err != nil {
		err = errors.Wrap(err, "building for darwin failed")
		return
	}

	// Execute template
	if err = executeTemplate(d, "./asset_darwin.go"); err != nil {
		err = errors.Wrap(err, "executing template failed")
		return
	}
	return
}

// buildDarwin builds the darwin binary and returns the darwin data
func buildDarwin() (d TemplateData, err error) {
	// Update args
	var args = []string{"-o", "./splashmake/tmp/darwin", "-framework", "Cocoa", "./splashmake/darwin.m"}

	// Build
	astilog.Debug("Building")
	var cmd = exec.Command("gcc", args...)
	cmd.Env = os.Environ()
	var b []byte
	if b, err = cmd.CombinedOutput(); err != nil {
		err = errors.Wrapf(err, "executing %s failed with output %s", strings.Join(cmd.Args, " "), b)
		return
	}

	// Read file
	astilog.Debug("Reading darwin binary")
	if b, err = ioutil.ReadFile("./splashmake/tmp/darwin"); err != nil {
		err = errors.Wrap(err, "reading \"./splashmake/tmp/darwin\" failed")
		return
	}

	// Write
	var buf = &bytes.Buffer{}
	var w = &bindata.StringWriter{Writer: buf}
	w.Write(b)
	d.Binary = buf.String()
	return
}
