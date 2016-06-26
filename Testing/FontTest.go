// full_test
package main

import (
	"GT/Graphics"
	"GT/Graphics/Components"
	//"GT/Window"
	"fmt"
	"math/rand"
	"GT"
)

func random(min, max int) int {
	//srand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}



func main() {

	fmt.Println("starting")
	GT.EngineStart()

	s, _ := Graphics.NewBasicScene()


	nodebak := Components.NewNode("Person")
	nodebak.Transform().Translate(100, 100)

	node := Components.NewNode("Person2")

	textRend := Graphics.NewTextRenderer()
	textRend.SetFont("Raleway")
	textRend.SetText("Hello World From GT!?")
	
	textRend.SetSize(14)
	textRend.SetColor(1,.2,.1,1)
	node.Transform().Translate(150, 100)

	textRend2 := Graphics.NewTextRenderer()
	textRend2.SetFont("Fantasque Sans Mono Regular")
	textRend2.SetSize(60)
	textRend2.SetText("This is Second Font Fantasque Bold")


	rend := Graphics.NewSpriteRenderer()
	rend.SetImage(GT.AssetsImages + "/test.png")
	nodebak.Transform().Scale(.5, .5)
	nodebak.Transform().Rotate(1.2)
	nodebak.AddComponent(rend)
	nodebak.AddComponent(textRend2)
	node.AddComponent(textRend)
	s.RootNode.AddNode(node)
	s.RootNode.AddNode(nodebak)
	
	s.Start()


}
