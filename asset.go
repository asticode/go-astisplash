package astisplash

import (
	"os"

	"github.com/pkg/errors"
)

// disembedAsset disembeds an asset
func disembedAsset(i []byte, path string) (err error) {
	// Create file
	var f *os.File
	if f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755); err != nil {
		err = errors.Wrapf(err, "creating %s failed", path)
		return
	}
	defer f.Close()

	// Write
	if _, err = f.Write(i); err != nil {
		err = errors.Wrapf(err, "writing to %s failed", path)
		return
	}
	return
}
