package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/asticode/go-astilog"
	"github.com/asticode/go-astitools/flag"
	"github.com/asticode/go-astitools/template"
	"github.com/pkg/errors"
)

// Vars
var t *template.Template

// Flags
var (
	oses = astiflag.Strings{}
)

func main() {
	// Parse flags
	flag.Var(&oses, "os", "one of the OS the making will be done for")
	flag.Parse()
	astilog.FlagInit()

	// Default oses
	if len(oses) == 0 {
		oses = astiflag.Strings{
			"darwin",
			"linux",
			"windows",
		}
	}

	// Load templates
	astilog.Debug("Loading templates")
	var err error
	if t, err = astitemplate.ParseDirectory("./splashmake", ".tmpl"); err != nil {
		astilog.Fatal(errors.Wrap(err, "loading templates failed"))
	}

	// Make
	astilog.Debug("Making")
	if err = makeAll(oses...); err != nil {
		astilog.Fatal(errors.Wrap(err, "building failed"))
	}
}

// makeAll makes all proper steps for all OSes
func makeAll(oses ...string) (err error) {
	for _, os := range oses {
		astilog.Debugf("Making for %s", os)
		switch os {
		case "darwin":
			err = makeDarwin()
		case "linux":
			err = makeLinux()
		case "windows":
			err = makeWindows()
		default:
			err = fmt.Errorf("making for %s not yet implemented", os)
		}
		if err != nil {
			err = errors.Wrapf(err, "making for %s failed", os)
			return
		}
	}
	return
}

// TemplateData represents template data
type TemplateData struct {
	Binary string
}

// executeTemplate executes a template
func executeTemplate(data interface{}, outputPath string) (err error) {
	// Create output
	astilog.Debugf("Creating %s", outputPath)
	var f *os.File
	if f, err = os.Create(outputPath); err != nil {
		err = errors.Wrapf(err, "creating %s failed", outputPath)
		return
	}
	defer f.Close()

	// Execute template
	astilog.Debug("Executing template")
	if err = t.ExecuteTemplate(f, "/shared.tmpl", data); err != nil {
		err = errors.Wrap(err, "executing template failed")
		return
	}
	return
}
