// full_test
package main

import (
	"GT/Graphics"
	"GT/Graphics/Components"
	"fmt"
	"GT"
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

type Bunny struct {
	Components.ChildComponent
	posX, posY, speedX, speedY float32
}

//TODO: create base class for component and/or node to impelemnt parent stuff because i forgot

func (this *Bunny) Initialize() {

}

func (this *Bunny) Update(delta float32) {

	this.posX += this.speedX
	this.posY += this.speedY
	this.speedY += 9.8

	if this.posX > 100 {
		this.speedX *= -1
		this.posX = 100
	}

	g := float32(random(0, 10) + 3)
	g *= .4

}

func main() {

	fmt.Println("starting")
	// // defer profile.Start(profile.CPUProfile).Stop()
	// w := Window.NewWindowedWindow("test", 600, 800)
GT.EngineStart()
	s, _ := Graphics.NewBasicScene()
	g := TestGame{BaseScene: &s}


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
		nodebak.AddComponent(&Bunny{})

		rend2 := Graphics.NewSpriteRenderer()
		rend2.SetSubImage("smiley.png", image.Rectangle{image.Point{30, 30}, image.Point{50, 50}})
		nodebak2 := Components.NewNode("Person")
		nodebak2.Transform().Translate(50, 50)
		nodebak2.AddComponent(rend2)
		// sr := nodebak2.GetComponent("SpriteRenderer")
		// fmt.Println(sr)

		node.AddNode(nodebak2)

		g.BaseScene.RootNode.AddNode(node)
	}
	g.Start()

}
