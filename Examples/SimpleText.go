// full_test
package main

import (
	"GT/Graphics"
	"GT/Graphics/Components"
	//"GT/Window"
	"GT"
	"fmt"
	"math/rand"
)

func random(min, max int) int {
	//srand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {

	fmt.Println("starting")
	GT.EngineStart()

	s, _ := Graphics.NewBasicScene()

	// create text rendering components
	textRend := Graphics.NewTextRenderer()
	textRend.SetFont("Raleway")
	textRend.SetText("Hello World From GT!?")
	textRend.SetSize(14)
	textRend.SetColor(1, .2, .1, 1)

	textRend2 := Graphics.NewTextRenderer()
	textRend2.SetFont("Fantasque Sans Mono Regular")
	textRend2.SetSize(60)
	textRend2.SetText("Hello World Pt.2")

	// create nodes for them so we can apply transformation
	node1 := Components.NewNode("Text1")
	node1.Transform().Translate(100, 100)
	node1.AddComponent(textRend)

	node2 := Components.NewNode("Text2")
	node2.Transform().Rotate(.3)
	node2.Transform().Translate(150, 100)
	node2.AddComponent(textRend2)

	s.RootNode.AddNode(node1)
	s.RootNode.AddNode(node2)

	s.Start()

}
