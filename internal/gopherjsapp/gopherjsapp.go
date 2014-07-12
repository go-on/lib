/*
Package gopherjsapp provides a way to have "mountable" javascript apps written in gopherjs.

An app is a package based on gopherjs that has an exported function Mount that
takes a mountpoint (string) and returns nothing.

This mount function is the main entry point for the whole app. No init nor main functions may be used.
*/

package gopherjsapp

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/go-on/gopherjslib"
)

// Mount compiles the given package and mounts it with the given mountpoint.
// It returns an handler that serves the resulting JavaScript

// An error is returned if compilation fails.
func Mount(pkg string, mountPoint string) (http.Handler, error) {
	code := strings.NewReader(`
    package main
    import app "` + pkg + `"
    func main() {
    	app.Mount("` + mountPoint + `")
    }
  `)

	var js bytes.Buffer

	err := gopherjslib.Build(code, &js, nil)

	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		rw.Write(js.Bytes())
	}), nil
}

// MountedJS compiles the pkg mounting it at mountPoint and returns the js bytes
// It panics if a compilation error occurs
func MountedJS(pkg string, mountPoint string) []byte {
	code := strings.NewReader(`
    package main
    import app "` + pkg + `"
    func main() {
    	app.Mount("` + mountPoint + `")
    }
  `)

	var js bytes.Buffer

	err := gopherjslib.Build(code, &js, nil)

	if err != nil {
		panic(err.Error())
	}
	return js.Bytes()
}
