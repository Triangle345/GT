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

func main() {

	fmt.Println("starting")

	// // defer profile.Start(profile.CPUProfile).Stop()
	w := Window.NewWindowedWindow("test", 600, 800)
	s, _ := Graphics.NewBasicScene(&w)
	g := TestGame{BaseScene: &s}

	//aggrImg := Graphics.NewAggregateImage("./")
	//aggrImg.Print("./aggregate.png")
	for i := 0; i < 1; i++ {
		nodebak := Components.NewNode("Person")
		nodebak.Transform().Translate(100, 100)
		
		node := Components.NewNode("Person2")

		textRend := Graphics.NewTextRenderer()
		textRend.SetText("Hello World From GT!")
		textRend.SetSize(50)
		node.Transform().Translate(100, 100)
		
		rend := Graphics.NewSpriteRenderer()
		rend.SetImage("test.png")
		//nodebak.Transform().Rotate(20)
		nodebak.Transform().Scale(.3,.3)
		nodebak.Transform().Rotate(1.5)
		nodebak.AddComponent(rend)
		node.AddComponent(textRend)
		//node.Transform().Rotate(1.5)
		//node.Transform().Scale(.5,.5)
		g.BaseScene.RootNode.AddNode(node)
		g.BaseScene.RootNode.AddNode(nodebak)
	}
	g.Start()

	w.Close()

}
