// full_test
package Testing

import (
	"GT/Graphics"
	// "GT/Scene"
	"GT/Window"
	// "fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"testing"
	// "time"
	"fmt"
)

func random(min, max int) int {
	//srand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func Test_Graphics(t *testing.T) {

	w := Window.NewWindowedWindow("test", 800, 600)
	
	s, _ := Graphics.NewBasicScene("smiley.png", 800, 600)

	s.AddSprite("smiley", Graphics.NewImageSection(0, 0, 128, 128))
	s.GetSprite("smiley").SetPosition(100, 100)

	// var window *sdl.Window
	// var context sdl.GLContext
	var event sdl.Event
	var running bool
	// var err error

	running = true

	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:

				fmt.Println(string(t.Timestamp))
			}
		}
		w.Clear()
		s.Draw()
		w.Refresh()
	}

	// // e := w.Open()

	// // if e != nil {
	// // 	t.Error("Window open failure: " + e.Error())
	// // }

	// // if w.isOpen() == false {
	// // 	t.Error("Window should be open but it's not")
	// // }

	// // if w.Width != 800 {
	// // 	t.Error("Window width should be 800")
	// // }

	// // if w.Height != 600 {
	// // 	t.Error("Window height should be 600")
	// // }

	// running := true

	// for running == true {

	// 	w.Refresh()

	// }

	w.Close()

}
