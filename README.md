GT
==

Go Graphics/Game engine

This is a simple Game engine written in Go.  It is really a thin wrapper around the go-gl wrapper.  Written a test of what it an do in the "Testing" section.  So far it seems pretty responsive with 5000 sprites drawn at once.

This project is still in it's infancy and has a minimal set of features.  Primarily; it can draw sprites on the screen and move them.

# Getting Started:
Note: we will likely have Travis CI set up *at some point* for builds and will include MinGW dlls and whatnot. 
However, this is the current installation procedure:
* [Install Go and configure your workspace.](https://golang.org/doc/install)
* [Install MinGW (64 bit) using Win-builds](http://win-builds.org/doku.php)
    * needed for 64 bit gcc
    * ensure that your path is updated with the bin contents after installing
* Clone this repo in your $GOPATH/src directory
* Install some Go dependancies:
    * go get github.com/go-gl/gl/v3.3-core/gl
    * go get github.com/go-gl/
    * go get github.com/veandco/go-sdl2
* Validate the install
    * Navigate to $GOPATH/src/GT/Testing
    * go run TestGame.go
    * Ctrl-C
