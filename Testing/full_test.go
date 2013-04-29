// full_test
package Testing

import (
	"Graphics"
	"Window"
	"fmt"
	"math/rand"
	"testing"
	//"time"
)

func random(min, max int) int {
	//srand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func Test_Graphics(t *testing.T) {
	w := Window.NewWindowedWindow("test", 800, 600)

	w.Open()

	img, e := Graphics.NewImage("test.png")
	s1 := Graphics.NewBasicSprite(img)
	img2, e := Graphics.NewImage("smiley.png")
	sprites := []*Graphics.Sprite{}

	for i := 0; i < 5000; i += 1 {
		x := random(1, 800)
		y := random(1, 600)

		fmt.Println("x:", x)
		fmt.Println("y:", y)

		s := Graphics.NewBasicSprite(img2)
		s.SetLocation(float64(x), float64(y))

		sprites = append(sprites, &s)
	}

	//glfw.SwapBuffers()

	s2 := Graphics.NewBasicSprite(img2)
	s2.SetLocation(400, 400)

	for w.IsOpen() == true {
		w.Clear()
		s1.Move(.1, .2)
		s2.Move(-.1, -.2)
		w.Draw(s1)
		w.Draw(s2)

		for _, mys := range sprites {
			dx := random(0, 20) / 10
			dy := random(0, 20) / 10

			mys.Move(float64(dx), float64(dy))

			w.Draw(mys)
		}

		w.Refresh()
	}
	//time.Sleep(5 * 10e8)

	if e != nil {
		t.Error(e.Error())
	}

}
