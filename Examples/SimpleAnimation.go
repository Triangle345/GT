// SimpleAnimation : This code demonstrates how to create a 2d animation from a sprite sheet

package main

import (
	"GT"
	"GT/Graphics"
	"GT/Graphics/Components"
	"fmt"
	"path/filepath"
)

func main() {

	fmt.Println("starting")

	// create the foundation: a new window, basic scene, and game component (node)
	GT.EngineStart()
	simpleScene, _ := Graphics.NewBasicScene()
	node := Components.NewNode("new_image")

	// splice a sprite sheet into multiple nodes... method tbd

	// anim := Graphics.NewAnimation()
	// myImagePath, _ := filepath.Abs("../Assets/Images/tomatohead1.png")
	// anim.SpliceAndSetSheet(myImagePath)
	// anim.Frequency(5) // number of frames it takes to toggle animation
	// node.AddComponent(anim)
	// OR
	rend := Graphics.NewSpriteRenderer()
	myImagePath, _ := filepath.Abs("../Assets/Images/tomatohead1.png")

	// right now it simply iterates the spliced images
	// we may need to have an option to identify the rows and animate them differently...
	// i.e. if a sheet has idle animations as well as movement animations,
	//      we need to change which to iterate based on input to determine row etc.
	//      this involves determining an expected / standardized sprite sheet format
	// also we may want to abstract this outside of sprite renderer?
	rend.SpliceAndSetSheet(myImagePath, 12, 12)
	simpleScene.SetFPS(4)

	// attach the sprite to our node, and transform if desired
	node.AddComponent(rend)
	node.Transform().Scale(3, 3)
	node.Transform().Translate(300, 400)

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)

	// start the scene to render our setup
	simpleScene.Start()

}
