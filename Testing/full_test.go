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

	for i := 0; i < 500; i += 1 {
		x := random(1, 800)
		y := random(1, 600)

		fmt.Println("x:", x)
		fmt.Println("y:", y)

		s := Graphics.NewBasicSprite(img2)
		s.SetLocation(float32(x), float32(y))

		sprites = append(sprites, &s)
	}

	s2 := Graphics.NewBasicSprite(img2)
	s2.SetLocation(400, 400)

	for w.IsOpen() == true {
		w.Clear()
		s1.Move(.1, .2)
		s2.Move(-.1, -.2)
		s2.Rotate(60.0)

		w.Draw(s1)
		w.Draw(s2)

		for _, mys := range sprites {
			dx := random(0, 20) / 10
			dy := random(0, 20) / 10

			mys.Move(float32(dx), float32(dy))
			mys.Rotate(float32(dx) * 10)

			w.Draw(mys)
		}

		w.Refresh()
	}

	if e != nil {
		t.Error(e.Error())
	}

}
