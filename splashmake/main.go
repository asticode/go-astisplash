package main

import (
	"flag"
	"text/template"

	"os"

	"github.com/asticode/go-astilog"
	"github.com/asticode/go-astitools/template"
	"github.com/pkg/errors"
)

// Vars
var t *template.Template

func main() {
	// Parse flags
	flag.Parse()
	astilog.SetLogger(astilog.New(astilog.FlagConfig()))

	// Load templates
	astilog.Debug("Loading templates")
	var err error
	if t, err = astitemplate.ParseDirectory("./splashmake", ".tmpl"); err != nil {
		astilog.Fatal(errors.Wrap(err, "loading templates failed"))
	}

	// Make
	astilog.Debug("Making")
	if err = makeAll(); err != nil {
		astilog.Fatal(errors.Wrap(err, "building failed"))
	}
}

// makeAll makes all proper steps for all OSes
func makeAll() (err error) {
	// Linux
	astilog.Debug("Making for Linux")
	if err = makeLinux(); err != nil {
		err = errors.Wrap(err, "making for linux failed")
		return
	}
	return

	// Windows
	astilog.Debug("Making for Windows")
	if err = makeWindows(); err != nil {
		err = errors.Wrap(err, "making for windows failed")
		return
	}
	return
}

// executeTemplate executes a template
func executeTemplate(data interface{}, templateName, outputPath string) (err error) {
	// Create output
	astilog.Debugf("Creating %s", outputPath)
	var f *os.File
	if f, err = os.Create(outputPath); err != nil {
		err = errors.Wrapf(err, "creating %s failed", outputPath)
		return
	}
	defer f.Close()

	// Execute template
	astilog.Debugf("Executing template %s", templateName)
	if err = t.ExecuteTemplate(f, templateName, data); err != nil {
		err = errors.Wrapf(err, "executing template %s failed", templateName)
		return
	}
	return
}
