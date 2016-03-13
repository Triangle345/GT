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

type Bunny struct {
	Components.GameComponent
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
	w := Window.NewWindowedWindow("test", 600, 800)
	s, _ := Graphics.NewBasicScene(&w)
	g := TestGame{BaseScene: &s}

	for i := 0; i < 14000; i++ {
		nodebak := Components.NewNode("Person")
		nodebak.Translate(400, 400)
		node := Components.NewNode("Person2")
		node.Translate(100, 100)
		node.AddNode(nodebak)
		rend := Graphics.NewSpriteRenderer()
		rend.SetImageSpriteSheet("test.png")
		nodebak.AddComponent(rend)
		nodebak.AddComponent(&Bunny{})

		// rend2 := Graphics.NewSpriteRenderer()
		// rend2.SetImageSpriteSheet("smiley.png")
		// nodebak2 := Components.NewNode("Person")
		// nodebak2.Translate(50, 50)
		// nodebak2.AddComponent(rend2)

		// node.AddNode(nodebak2)

		g.BaseScene.RootNode.AddNode(node)
	}
	g.Start()

	w.Close()

}
