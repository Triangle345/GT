// SimpleWindow : This code demonstrates how to create a simple (windowed) game instance while rendering a png image

package main

import (
	"GT"
	"GT/Graphics"
	"GT/Graphics/Components"
	"fmt"
)

type Spin struct {
	Components.ScriptComponent
}

func (this *Spin) Initialize() {
	//no initialize for RunLeft
	this.Transform().Rotate(0, .2, 1, .2)

}

func (this *Spin) Update(delta float32) {
	this.Transform().Rot += .01

}

func main() {

	fmt.Println("starting")
	GT.EngineStart()

	simpleScene, _ := Graphics.New3DScene()
	node := Components.NewNode("new_image")
	node.AddComponent(&Spin{})
	// create a png based sprite
	rend := Components.NewModelRenderer()
	// myImagePath, _ := filepath.Abs("../Assets/Images/test.png")
	//rend.SetImage(GT.AssetsImages + "test.png")

	// attach the sprite to our node, and transform if desired
	node.AddComponent(rend)
	//node.Transform().Rotate(.2, 0, 1, 0)
	// node.Transform().Translate(1, 2, 0)

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)

	// start the scene to render our setup
	simpleScene.Start()
}
