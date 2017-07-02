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

// DataLinux represents linux template data
type DataLinux struct {
	Binary string
}

// makeLinux makes all proper steps for Linux
func makeLinux() (err error) {
	// Build
	var d DataLinux
	if d, err = buildLinux(); err != nil {
		err = errors.Wrap(err, "building for linux failed")
		return
	}

	// Execute template
	if err = executeTemplate(d, "/linux.tmpl", "./asset_linux.go"); err != nil {
		err = errors.Wrap(err, "executing template failed")
		return
	}
	return
}

// buildLinux builds the linux binary and returns the linux data
func buildLinux() (d DataLinux, err error) {
	// Update args
	var args = []string{"-o", "./splashmake/tmp/linux", "./splashmake/splash.c"}

	// Retrieve pkg-config
	astilog.Debug("Retrieving pkg-config")
	var cmd = exec.Command("pkg-config", "--cflags", "--libs", "gtk+-3.0")
	var b []byte
	if b, err = cmd.CombinedOutput(); err != nil {
		err = errors.Wrap(err, "retrieving pkg-config failed")
		return
	}
	args = append(args, strings.Split(string(bytes.TrimSpace(b)), " ")...)

	// Build
	astilog.Debug("Building")
	cmd = exec.Command("gcc", args...)
	cmd.Env = os.Environ()
	if b, err = cmd.CombinedOutput(); err != nil {
		err = errors.Wrapf(err, "executing %s failed with output %s", strings.Join(cmd.Args, " "), b)
		return
	}

	// Read file
	astilog.Debug("Reading linux binary")
	if b, err = ioutil.ReadFile("./splashmake/tmp/linux"); err != nil {
		err = errors.Wrap(err, "reading \"./splashmake/tmp/linux\" failed")
		return
	}

	// Write
	var buf = &bytes.Buffer{}
	var w = &bindata.StringWriter{Writer: buf}
	w.Write(b)
	d.Binary = buf.String()
	return
}
