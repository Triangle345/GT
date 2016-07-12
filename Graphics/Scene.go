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
	drawStart := int(time.Now().Unix())
	drawEnd := int(time.Now().Unix())

	// continue to render the scene unless signalled for termination
	for running {

		if GlobalInput.CheckForUpdates() {
			running = !(GlobalInput.GetInputStatus("Esc") || GlobalInput.GetInputStatus("Quit"))
		}

		s.window.Clear()
		drawEnd = int(time.Now().Unix())

		// TODO: implement or store this data. Implement this with the fps limiting mentioned below
		delta := float32(drawEnd - drawStart)
		s.RootNode.Update(delta)
		drawStart = int(time.Now().Unix())
		s.Draw()

		// fps limitation is optional, sleep for at least 1 ns to smooth rendering
		if s.targetFps > 0 {
			if delta < float32(time.Second/time.Duration(s.targetFps)) {
				// (1 second / n frames - time elapsed) = leftover sleep needed to reach target fps
				time.Sleep((time.Second/time.Duration(s.targetFps) - time.Duration(delta)))
			}
		} else {
			time.Sleep(time.Nanosecond)
		}
		s.window.Refresh()
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
