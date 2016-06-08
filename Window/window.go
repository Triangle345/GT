// window project window.go
package Window

import (
	// "GT/Graphics/Opengl"

	"GT/Graphics/Opengl"
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	title         string
	Width, Height int
	full          bool
	windowSDL     *sdl.Window
	contextSDL    sdl.GLContext
	running       bool
}

func NewWindowedWindow(title string, width, height int) Window {
	// FOR GLFW event handling
	runtime.LockOSThread()

	windowSDL, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_OPENGL)

	if err != nil {
		panic(err)
	}
	if windowSDL == nil {
		panic(sdl.GetError())
	}

	context, err := sdl.GL_CreateContext(windowSDL)

	// call init right after creating context
	if err := gl.Init(); err != nil {
		fmt.Println("Cannot initialize OGL: " + err.Error())
	}

	// set viewport for window
	gl.Viewport(0, 0, int32(width), int32(height))

	sdl.GL_MakeCurrent(windowSDL, context)

	if err != nil {
		panic("Error loading window: " + err.Error())
	}

	w := Window{title: title, Width: width, Height: height, full: false, windowSDL: windowSDL, contextSDL: context, running: false}

	w.init()
	var current sdl.DisplayMode
	if err := sdl.GetCurrentDisplayMode(0, &current); err != nil {
		fmt.Println("COuld not get display mode: " + err.Error())
	}
	fmt.Printf("Display #%d: current display mode is %dx%dpx @ %dhz. \n", 0, current.W, current.H, current.RefreshRate)

	return w
}

func (w Window) init() error {

	sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	sdl.GL_SetAttribute(sdl.GL_RED_SIZE, 8)
	sdl.GL_SetAttribute(sdl.GL_GREEN_SIZE, 8)
	sdl.GL_SetAttribute(sdl.GL_BLUE_SIZE, 8)
	sdl.GL_SetAttribute(sdl.GL_ALPHA_SIZE, 8)

	// this only needs to be called once
	Opengl.CreateBuffers()

	//TODO: find a better way to load all images in
	//however; this has logic to handle more than one call
	//// TODO: put this in opengl along with aggregate image
	//Image.AggrImg.Bind2GL()

	return nil
}

func (w Window) Clear() {
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w Window) Refresh() {
	sdl.Delay(0)
	sdl.GL_SwapWindow(w.windowSDL)

}

// func (w window) Draw(entity ...Graphics.Drawable) {
// 	for _, v := range entity {
// 		v.Draw(w.projectionM, w.viewM)
// 	}

// }

// func (w window) IsOpen() bool {
// 	if glfw.WindowParam(glfw.Opened) == 1 {
// 		return true
// 	}

// 	return false
// }

func (w Window) Close() {
	// sdl.GL_DeleteContext(w.contextSDL)
	// w.windowSDL.Destroy()
	// sdl.Quit()
}
