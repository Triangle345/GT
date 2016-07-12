// SimpleWindow : This code demonstrates how to create a simple (windowed) game instance while rendering a png image

package main

import (
	"GT"
	"GT/Graphics"
	"GT/Graphics/Components"
	"fmt"
	"math/rand"
)

func random(min, max int) float32 {
	//srand.Seed(time.Now().Unix())
	return float32(rand.Intn(max-min) + min)
}

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
	node.Transform().Translate(-4, -3, 0)
	// create a model based on obj file
	rend := Components.NewModelRenderer()
	rend.SetModel(GT.AssetsModels+"textured_box.obj", GT.AssetsModels+"textured_box.mtl")
	node.AddComponent(rend)

	node2 := Components.NewNode("new_image2")
	// node2.AddComponent(&Spin{})
	// create a model based on obj file
	rend2 := Components.NewModelRenderer()
	rend2.SetModel(GT.AssetsModels+"walk.obj", GT.AssetsModels+"walk.mtl")
	node2.AddComponent(rend2)
	node2.Transform().Translate(4, -3, 0)
	node2.Transform().Rotate(.3, 0, 1, 0)

	node3 := Components.NewNode("new_image3")
	node3.AddComponent(&Spin{})
	// node2.AddComponent(&Spin{})
	// create a model based on obj file
	rend3 := Components.NewModelRenderer()
	rend3.SetModel(GT.AssetsModels+"test.obj", GT.AssetsModels+"test.mtl")
	node3.AddComponent(rend3)
	node3.Transform().Translate(4, 3, 0)
	node3.Transform().Rotate(.3, 0, 1, 0)

	node4 := Components.NewNode("new_image4")
	node4.AddComponent(&Spin{})
	rend4 := Components.NewModelRenderer()
	node4.AddComponent(rend4)
	// node2.AddComponent(&Spin{})
	// create a model based on obj file
	//node4.Transform().Translate(0, 0, 0)
	node4.Transform().Rotate(.3, 0, 1, 0)
	anim := Components.OBJAnimation(GT.AssetsModels + "SimpleAnimation")

	rend4.AnimationManager.AddAnimation(anim, "anim1")
	rend4.AnimationManager.SetCurrentAnimation("anim1")

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)
	simpleScene.RootNode.AddNode(node2)
	simpleScene.RootNode.AddNode(node3)
	simpleScene.RootNode.AddNode(node4)

	// start the scene to render our setup
	simpleScene.Start()
}
