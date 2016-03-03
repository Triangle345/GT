// full_test
package main

import (
	"GT/Graphics"
	"GT/Window"
	// "github.com/veandco/go-sdl2/sdl"
	"math/rand"
	// "time"
	"GT/Graphics/Components"
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
	for i := 0; i < 5; i++ {
		g.AddSprite("smiley"+string(i), Graphics.NewRectangularArea(0, 0, 128, 128))
		g.GetSprite("smiley"+string(i)).SetLocation(float32(random(0, 600)), 100.0)
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

	for i := 0; i < 2; i++ {
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

	// TODO maybe get root node from scene because of messup with putting wrong node in start

	nodebak := Components.NewNode("Person")
	nodebak.Translate(100, 100)
	node := Components.NewNode("Person2")
	node.Translate(100, 100)
	node.AddNode(nodebak)
	rend := Graphics.NewSpriteRenderer()
	rend.SetImageSpriteSheet("smiley.png")
	nodebak.AddComponent(rend)

	g.BaseScene.RootNode.AddNode(node)

	g.Start()

	w.Close()

}
