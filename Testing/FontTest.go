// full_test
package main

import (
	"GT/Graphics"
	"GT/Graphics/Components"
	//"GT/Window"
	"fmt"
	"math/rand"
	"path/filepath"
	"GT"
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
	GT.EngineStart()

	s, _ := Graphics.NewBasicScene()
	g := TestGame{BaseScene: &s}

	myTestImage, _ := filepath.Abs("../Assets/Images/test.png")

	for i := 0; i < 1; i++ {
		nodebak := Components.NewNode("Person")
		nodebak.Transform().Translate(100, 100)

		node := Components.NewNode("Person2")

		textRend := Graphics.NewTextRenderer()
		textRend.SetText("Hello World From GT!")
		textRend.SetSize(50)
		textRend.SetColor(1,.2,.1,1)
		node.Transform().Translate(150, 100)

		textRend2 := Graphics.NewTextRenderer()
		textRend2.SetFont("Fantasque Sans Mono Regular")
		textRend2.SetSize(70)
		textRend2.SetText("This is the second font")


		rend := Graphics.NewSpriteRenderer()
		rend.SetImage(myTestImage)
		nodebak.Transform().Scale(.5, .5)
		nodebak.Transform().Rotate(1.2)
		nodebak.AddComponent(rend)
		nodebak.AddComponent(textRend2)
		node.AddComponent(textRend)
		g.BaseScene.RootNode.AddNode(node)
		g.BaseScene.RootNode.AddNode(nodebak)
	}
	g.Start()


}
