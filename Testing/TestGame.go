// TestGame runs a simple sprite render to test simple graphics and performance indications. For now needs to be run with -Assets "../Assets"
package main

import (
	"GT"
	"GT/Graphics"
	"GT/Graphics/Components"
	"fmt"
	"image"
	"math/rand"
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
	// w := Window.NewWindowedWindow("test", 600, 800)
	GT.EngineStart()
	s, _ := Graphics.NewBasicScene()
	g := TestGame{BaseScene: &s}

	// 500 sprites does about 38 fps on my laptop
	for i := 0; i < 500; i++ {
		nodebak := Components.NewNode("Person")
		nodebak.Transform().Translate(400, 400,0)
		node := Components.NewNode("Person2")

		node.Transform().Translate(100, 100,0)
		node.AddNode(nodebak)
		rend := Components.NewSpriteRenderer()
		rend.SetImage(GT.AssetsImages + "test.png")
		nodebak.Transform().Rotate(20,0,0,1)
		nodebak.AddComponent(rend)

		rend2 := Components.NewSpriteRenderer()
		rend2.SetSubImage(GT.AssetsImages+"smiley.png", image.Rectangle{image.Point{30, 30}, image.Point{50, 50}})
		nodebak2 := Components.NewNode("Person")
		nodebak2.Transform().Translate(50, 50,0)
		nodebak2.AddComponent(rend2)
		// sr := nodebak2.GetComponent("SpriteRenderer")
		// fmt.Println(sr)

		node.AddNode(nodebak2)

		g.BaseScene.RootNode.AddNode(node)
	}
	g.Start()

}
