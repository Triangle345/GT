// window project window.go
package Window

import (
	"GT/Graphics/Opengl"
	"errors"
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
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

	if err := gl.Init(); err != nil {
		return errors.New("Cannot initialize OGL: ")
	}

	gl.Viewport(0, 0, int32(w.Width), int32(w.Height))
	Opengl.CreateBuffers(w.Width, w.Height)

	return nil
}

func (w Window) Clear() {
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w Window) Refresh() {
	sdl.Delay(1)
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
	sdl.GL_DeleteContext(w.contextSDL)
	w.windowSDL.Destroy()
	sdl.Quit()
}
