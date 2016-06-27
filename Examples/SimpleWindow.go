// SimpleWindow : This code demonstrates how to create a simple (windowed) game instance while rendering a png image

package main

import (
	"GT"
	"GT/Graphics"
	"GT/Graphics/Components"
	"fmt"
)

func main() {

	fmt.Println("starting")
	GT.EngineStart()

	simpleScene, _ := Graphics.NewBasicScene()
	node := Components.NewNode("new_image")

	// create a png based sprite
	rend := Components.NewSpriteRenderer()
	// myImagePath, _ := filepath.Abs("../Assets/Images/test.png")
	rend.SetImage(GT.AssetsImages + "test.png")

	// attach the sprite to our node, and transform if desired
	node.AddComponent(rend)
	node.Transform().Rotate(0)
	node.Transform().Translate(400, 400)

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)

	// start the scene to render our setup
	simpleScene.Start()
}
