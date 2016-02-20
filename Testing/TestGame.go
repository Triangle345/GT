// full_test
package main

import (
	"GT/Graphics"
	"GT/Window"
	// "github.com/veandco/go-sdl2/sdl"
	"math/rand"
	// "time"
	"fmt"
	// "github.com/davecheney/profile"
)

func random(min, max int) int {
	//srand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

type TestGame struct {
	*Graphics.BaseScene
}

var first bool
var skip bool

func (g *TestGame) Load() {
	for i := 0; i < 15000; i++ {
		g.AddSprite("smiley"+string(i), Graphics.NewRectangularArea(0, 0, 128, 128))
		g.GetSprite("smiley"+string(i)).SetLocation(float32(random(0, 500)), 100.0)
	}

	// first = true
	// skip = true
}

func (g *TestGame) Update() {
	// if skip == true {
	//
	// 	skip = false
	// 	return
	// }
	// if first == true {
	// 	for i := 10; i < 2000; i++ {
	// 		g.AddSprite("smiley"+string(i), Graphics.NewRectangularArea(0, 0, 128, 128))
	// 		g.GetSprite("smiley"+string(i)).SetLocation(float32(random(0, 500)), 100.0)
	// 	}
	// 	first = false
	// }
	for i := 0; i < 20; i++ {
		x, y := g.GetSprite("smiley" + string(i)).GetLocation()
		x += 1
		g.GetSprite("smiley"+string(i)).SetLocation(float32(x), y)
		// fmt.Printf("smileyend %d has y %f\n", i, y)
	}

}

func main() {
	fmt.Println("starting")
	// // defer profile.Start(profile.CPUProfile).Stop()
	w := Window.NewWindowedWindow("test", 600, 800)
	s, _ := Graphics.NewBasicScene("smiley.png", &w)
	g := TestGame{BaseScene: &s}

	g.LoadHandler = g.Load
	g.UpdateHandler = g.Update

	// fmt.Println(g)
	// for i := 0; i < 25000; i++ {
	// s.AddSprite("smiley"+string(i), Graphics.NewImageSection(0, 0, 128, 128))
	// s.GetSprite("smiley"+string(i)).SetLocation(float32(0), float32(random(0, 500)))
	// _, y := s.GetSprite("smiley" + string(i)).GetLocation()
	// fmt.Printf("smileystart %d has y %f\n", i, y)
	// }

	// var window *sdl.Window
	// var context sdl.GLContext
	// var event sdl.Event
	// var running bool
	// var err error

	// running = true
	// // x := 0
	// for running {
	// 	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch t := event.(type) {
	// 		case *sdl.QuitEvent:
	// 			running = false
	// 		case *sdl.MouseMotionEvent:

	// 			fmt.Println(string(t.Timestamp))
	// 		}
	// 	}
	// w.Clear()
	// for i := 0; i < 25000; i++ {
	// 	_, y := s.GetSprite("smiley" + string(i)).GetLocation()
	// 	s.GetSprite("smiley"+string(i)).SetLocation(float32(x), y)
	// 	// fmt.Printf("smileyend %d has y %f\n", i, y)
	// }
	// s.Draw()
	// w.Refresh()

	g.Start()

	// x += 1
	// }

	// // e := w.Open()

	// // if e != nil {
	// //   t.Error("Window open failure: " + e.Error())
	// // }

	// // if w.isOpen() == false {
	// //   t.Error("Window should be open but it's not")
	// // }

	// // if w.Width != 800 {
	// //   t.Error("Window width should be 800")
	// // }

	// // if w.Height != 600 {
	// //   t.Error("Window height should be 600")
	// // }

	// running := true

	// for running == true {

	//  w.Refresh()

	// }

	w.Close()

}
