// SimpleWindow : This code demonstrates how to create a simple (windowed) game instance while rendering a png image

package main

import (
	"GT"
	"GT/Graphics"
	"GT/Graphics/Components"
	"GT/Graphics/Components/Model"
	"fmt"
)

type Spin struct {
	Components.ScriptComponent
	rot float32
}

func (this *Spin) Initialize() {
	//no initialize for RunLeft
	this.Transform().Rotate(0, .1, .5, .2)

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
	// create a model based on obj file
	rend := Model.NewModelRenderer()
	rend.SetModel(GT.AssetsModels+"textured_box.obj", GT.AssetsModels+"textured_box.mtl")
	node.AddComponent(rend)

	node2 := Components.NewNode("new_image")
	node2.AddComponent(&Spin{})
	// create a model based on obj file
	rend2 := Model.NewModelRenderer()
	rend2.SetModel(GT.AssetsModels+"test.obj", GT.AssetsModels+"test.mtl")
	node2.AddComponent(rend2)
	node2.Transform().Translate(3, -3, 0)
	// node2.Transform().Rotate(.3, .2, 1, .2)

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)
	simpleScene.RootNode.AddNode(node2)

	// start the scene to render our setup
	simpleScene.Start()
}
