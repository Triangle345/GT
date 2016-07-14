// Scene.go : contain game components within nodes, and stucture them inside of scenes for rendering etc

package Graphics

import (
	"GT/Graphics/Components"
	"GT/Graphics/Opengl"
	"GT/Input"
	"GT/Window"
	"fmt"
	"time"
)

// BaseScene : object to contain nodes for rendering. TODO: implement more scene types from this base type
type BaseScene struct {
	RootNode *Components.Node

	window *Window.Window
	// TODO: can put fps into struct fps counter
	targetFps int
	fps       int
	timestart int32
	update    bool
}

// NewBasicScene : start a new simple 2D scene of orthographic orientation
func NewBasicScene() (BaseScene, error) {

	s := BaseScene{RootNode: Components.NewNode("RootNode"), window: &Window.MainWindow}
	s.fps = 0
	s.timestart = int32(time.Now().Unix())
	s.update = true

	Opengl.SetViewPort(int32(s.window.Width), int32(s.window.Height))
	Opengl.SetOrthographic(s.window.Width, s.window.Height)

	return s, nil
}

// New3DScene : start a new simple 3D scene of orthographic orientation
func New3DScene() (BaseScene, error) {

	s := BaseScene{RootNode: Components.NewNode("RootNode"), window: &Window.MainWindow}
	s.fps = 0
	s.timestart = int32(time.Now().Unix())
	s.update = true

	Opengl.SetViewPort(int32(s.window.Width), int32(s.window.Height))
	Opengl.SetPerspective(s.window.Width, s.window.Height)

	return s, nil
}

// Start : initiate a scene's continous animation. While the scene is running, draw it and update our nodes
func (s *BaseScene) Start() {

	running := true
	s.update = true
	drawStart := int32(time.Now().UnixNano())
	drawEnd := int32(time.Now().UnixNano())
	delta := 0
	onesec := int32(time.Second.Nanoseconds())
	desiredFps := int32(s.targetFps)

	// continue to render the scene unless signalled for termination
	for running {

		drawStart = int32(time.Now().UnixNano())

		if GlobalInput.CheckForUpdates() {
			running = !(GlobalInput.GetInputStatus("Esc") || GlobalInput.GetInputStatus("Quit"))
		}

		s.window.Clear()

		s.RootNode.Update(float32(delta))
		s.Draw()
		s.window.Refresh()

		drawEnd = int32(time.Now().UnixNano())

		delta := drawEnd - drawStart

		// fps limitation is optional, sleep for at least 1 ns to smooth rendering
		if desiredFps > 0 {
			// (1 second in ns / n frames - time elapsed in ns) = leftover sleep in ns needed to reach target fps
			if delta < onesec/desiredFps {
				time.Sleep(time.Duration(onesec/desiredFps - delta))
			}
		} else {
			time.Sleep(time.Nanosecond)
		}
	}
}

// Draw : wrap the openGl draw method and report fps info
func (s *BaseScene) Draw() {

	if s.update {
		// TODO: remove hard coded resolution
		// gl.BindTexture(gl.TEXTURE_2D, s.spriteSheet.textureId)
		Opengl.BindBuffers()
		s.update = false
	}

	Opengl.Draw()

	// calc fps
	if (int32(time.Now().Unix()) - s.timestart) >= 1 {
		fmt.Printf("FPS: %d \n", s.fps)
		s.fps = 0
		s.timestart = int32(time.Now().Unix())
	}

	s.fps = s.fps + 1
}

// SetFPS forces the loop to delay such that the fps is the number inputted
func (s *BaseScene) SetFPS(fpsInputted int) {
	s.targetFps = fpsInputted
}
