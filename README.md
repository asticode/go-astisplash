This package provides a way to display a cross-platform splash screen without having to feel the pain of cross compiling.

Indeed, instead of using CGO that would require to install each and every toolchains in order to cross compile, this package relies on executing already cross-compiled binaries that are embedded directly in GO.

# Installation

Run the following command:

    $ go get -u github.com/asticode/go-astisplash
    
# Usage

WARNING: the code below doesn't handle errors for readibility purposes. However you SHOULD!

```go
// Build splasher
s, _ := astisplash.New()
defer s.Close()

// Splash
sp, _ := s.Splash("/path/to/your/image.png")

// Sleep
time.Sleep(2 * time.Second)

// Close the splash
sp.Close()
```

# Developments

When modifying the C code, you'll have to run the following:

    $ cd $GOPATH/src/github.com/asticode/go-astisplash
    $ make