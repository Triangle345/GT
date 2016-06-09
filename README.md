GT
==

## Summary

Go Graphics/Game engine

This is a simple Game engine written in Go.  It is really a thin wrapper around the go-gl wrapper.  Written a test of what it an do in the "Testing" section.  So far it seems pretty responsive with 5000 sprites drawn at once.

This project is still in it's infancy and has a minimal set of features.  Primarily; it can draw sprites on the screen and move them.

## Getting Started
The current installation procedure:
* [Install Go and configure your workspace](https://golang.org/doc/install)
* Windows Only: [Install MinGW (64 bit) using Win-builds](http://win-builds.org/doku.php)
    * needed for 64 bit gcc
    * ensure that your path is updated with the bin contents after installing
* Clone this repo in your $GOPATH/src directory
* Install some Go dependancies (may need less depending on your go set-up):
	* go get github.com/go-gl/gl/v3.2-core/gl
	* go get github.com/go-gl/mathgl/mgl32
	* go get github.com/golang/freetype/truetype
	* go get github.com/veandco/go-sdl2/sdl
	* go get golang.org/x/image/font
	* go get golang.org/x/image/math/fixed
* Validate the install
    * Navigate to $GOPATH/src/GT/Testing
    * go run TestGame.go
    * Ctrl-C

## Installation Troubleshooting
If you encounter any issues after / during installation check out these pointers:
* It is important to have the 64 bit version of MinGW installed for Windows developers
* Ensure that all go dependencies have been resolved
* This engine uses OpenGL and SDL, so updated graphics drivers may be necessary
	* Because of this, running this engine within a virtual machine could be difficult

## Other Notes
We will likely have Travis CI set up *at some point* for builds and will include MinGW dlls and whatnot. So hopefully the installation process will require less work on the user's end.

We also plan to make dev installation as easy as "go get github.com/Triangle345/GT" but at the moment the project needs to be cloned (or download the zip)
