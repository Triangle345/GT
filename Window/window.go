// window project window.go
package Window

import (
	"Graphics"
	"errors"
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"os"
)

type window struct {
	title         string
	width, height int
	full          bool
}

func NewWindowedWindow(title string, width, height int) window {

	w := window{title, width, height, false}

	return w
}

func (w window) init() error {
	if err := gl.Init(); err != nil {
		return errors.New("Cannot initialize OGL: " + err.Error())
	}

	gl.MatrixMode(gl.PROJECTION)
	gl.Disable(gl.DEPTH_TEST)

	gl.Ortho(0, gl.Double(w.width), gl.Double(w.height), 0, -1, 1)

	return nil
}

func (w window) Clear() {
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w window) Refresh() {
	glfw.SwapBuffers()
}

func (w window) Draw(entity ...Graphics.Drawable) {
	for _, v := range entity {
		v.Draw()
	}

}

func (w window) IsOpen() bool {
	if glfw.WindowParam(glfw.Opened) == 1 {
		return true
	}

	return false
}

func (w window) Close() {
	glfw.CloseWindow()
	glfw.Terminate()
}

func (w window) Open() error {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return err
	}

	// TODO need to make boolean for no resize, etc

	if err := glfw.OpenWindow(
		w.width, w.height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return err
	}

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(w.title)

	err := w.init()

	if err != nil {
		return err
	}

	return nil
}
