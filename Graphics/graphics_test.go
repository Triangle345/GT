// window_test.go
package Graphics

import (
	"Window"
	"fmt"
	"testing"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
)

func Test_Graphics(t *testing.T) {
	w := Window.NewWindowedWindow("test", 800, 600)

	w.Open()

	// TODO put this in the scene class or soemthing
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	img, e := NewImage("test.png")
	s1 := NewBasicSprite(img)
	s1.draw()

	glfw.SwapBuffers()
	img2, e := NewImage("smiley.png")
	s2 := NewBasicSprite(img2)
	s2.SetLocation(400, 400)
	s2.draw()

	glfw.SwapBuffers()
	time.Sleep(5 * 10e8)

	if e != nil {
		t.Error(e.Error())
	}

	fmt.Println("ID")
	fmt.Println(img.textureId)
	fmt.Println(img2.textureId)

}
