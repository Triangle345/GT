// full_test
package main

import (
	"GT/Graphics"
	"GT/Window"
	// "github.com/veandco/go-sdl2/sdl"
	"math/rand"
	// "time"
	"GT/Graphics/Components"
	"GT/Graphics/Font"
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

func main() {

	Font.ReadFonts("./")

	fmt.Println("starting")

	// // defer profile.Start(profile.CPUProfile).Stop()
	w := Window.NewWindowedWindow("test", 600, 800)
	s, _ := Graphics.NewBasicScene(&w)
	g := TestGame{BaseScene: &s}

	//aggrImg := Graphics.NewAggregateImage("./")
	//aggrImg.Print("./aggregate.png")
	for i := 0; i < 1; i++ {
		nodebak := Components.NewNode("Person")
		nodebak.Transform().Translate(400, 400)
		node := Components.NewNode("Person2")

		node.Transform().Translate(100, 100)
		node.AddNode(nodebak)
		rend := Graphics.NewSpriteRenderer()
		rend.SetImage("test.png")
		nodebak.Transform().Rotate(20)
		nodebak.AddComponent(rend)

		g.BaseScene.RootNode.AddNode(node)
	}
	g.Start()

	w.Close()

}
