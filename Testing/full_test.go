// full_test
package Testing

import (
	"Graphics"
	"Window"
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"math/rand"
	"testing"
	"time"
)

func random(min, max int) int {
	//srand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func Test_VBO(t *testing.T) {
	w := Window.NewWindowedWindow("test", 800, 600)

	w.Open()

	w.Clear()
	w.Refresh()

	ID := make([]gl.Uint, 1)
	data := make([]float32, 6)
	data[0] = 50
	data[1] = 50
	data[2] = 100
	data[3] = 50
	data[4] = 74
	data[5] = 100

	//gl.GenVertexArrays(1, &ID[0]) // Create our Vertex Array Object
	//gl.BindVertexArray(ID[0])     // Bind our Vertex Array Object so we can use it

	gl.GenBuffers(1, &ID[0])
	gl.BindBuffer(gl.ARRAY_BUFFER, ID[0])
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(4*len(data)), gl.Pointer(&data[0]), gl.STATIC_DRAW)

	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Color3f(0.0, 0.0, 0.0)
	gl.BindBuffer(gl.ARRAY_BUFFER, ID[0])
	// 4 is size of our float, but 2 objects each
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, gl.FALSE, 0, gl.Pointer(&data[0]))

	gl.VertexPointer(2, gl.FLOAT, 2*4, gl.Offset(nil, 0))
	//glEnableVertexAttribArray(0); // Disable our Vertex Array Object
	//glBindVertexArray(0); // Disable our Vertex Buffer Object
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.Flush()
	w.Refresh()
	time.Sleep(3 * 10e8)
}

func Test_Graphics(t *testing.T) {
	w := Window.NewWindowedWindow("test", 800, 600)

	w.Open()

	img, e := Graphics.NewImage("test.png")
	s1 := Graphics.NewBasicSprite(img)
	img2, e := Graphics.NewImage("smiley.png")
	sprites := []*Graphics.Sprite{}

	for i := 0; i < 500; i += 1 {
		x := random(1, 800)
		y := random(1, 600)

		fmt.Println("x:", x)
		fmt.Println("y:", y)

		s := Graphics.NewBasicSprite(img2)
		s.SetLocation(float32(x), float32(y))

		scale := random(2, 6)
		fmt.Println("scale", float32(scale)/10.0)
		s.Scale(float32(scale)/10.0, float32(scale)/10.0)
		sprites = append(sprites, &s)
	}

	s2 := Graphics.NewBasicSprite(img2)
	s2.SetLocation(400, 400)

	for w.IsOpen() == true {
		w.Clear()
		s1.Move(.1, .2)
		s2.Move(-.1, -.2)
		s2.Rotate(20.0)

		w.Draw(s1)
		w.Draw(s2)

		for _, mys := range sprites {
			dx := random(0, 20) / 10
			dy := random(0, 20) / 10

			mys.Move(float32(dx), float32(dy))
			mys.Rotate(float32(dx))

			w.Draw(mys)
		}

		w.Refresh()
	}

	if e != nil {
		t.Error(e.Error())
	}

}
