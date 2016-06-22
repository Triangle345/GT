// full_test
package main

import (
	"GT/Graphics"
	"GT/Graphics/Components"
	"GT/Window"
	"fmt"
	"math/rand"
	"path/filepath"
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

	myTestImage, _ := filepath.Abs("../Assets/Images/test.png")

	for i := 0; i < 1; i++ {
		nodebak := Components.NewNode("Person")
		nodebak.Transform().Translate(100, 100)

		node := Components.NewNode("Person2")

		textRend := Graphics.NewTextRenderer()
		textRend.SetText("Hello World From GT!")
		textRend.SetSize(15)
		node.Transform().Translate(150, 100)

		rend := Graphics.NewSpriteRenderer()
		rend.SetImage(myTestImage)
		nodebak.Transform().Scale(.5, .5)
		nodebak.Transform().Rotate(1.5)
		nodebak.AddComponent(rend)
		node.AddComponent(textRend)
		g.BaseScene.RootNode.AddNode(node)
		g.BaseScene.RootNode.AddNode(nodebak)
	}
	g.Start()

	w.Close()

}
