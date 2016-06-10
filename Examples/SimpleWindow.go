// This code demonstrates how to create a simple (windowed) game instance while rendering a png image
package main

import (
	"GT/Graphics"
	"GT/Graphics/Components"
	"GT/Window"
	"fmt"
	"time"
)

func main() {

	fmt.Println("starting")

	// create the foundation: a new window, basic scene, and game component (node)
	simpleWindow := Window.NewWindowedWindow("test", 600, 800)
	simpleScene, _ := Graphics.NewBasicScene(&simpleWindow)
	node := Components.NewNode("new_image")

	// create a png based sprite
	rend := Graphics.NewSpriteRenderer()
	rend.SetImage("test.png")

	// attach the sprite to our node, and transform if desired
	node.AddComponent(rend)
	node.Transform().Rotate(0)
	node.Transform().Translate(300, 400)

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)

	// manually draw the scene
	simpleScene.RootNode.Update(.34)
	simpleScene.Draw()
	simpleWindow.Refresh()

	time.Sleep(5 * 10e8)

	// close the window
	simpleWindow.Close()
}
