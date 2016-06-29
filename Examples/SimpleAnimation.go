// SimpleAnimation : This code demonstrates how to create a 2d animation from a sprite sheet

package main

import (
	"GT"
	"GT/Graphics"
	"GT/Graphics/Components"
	"fmt"
)

//////////////////////////////////////////////////////
// TODO: remove this when a scripting example is made
//    clean this up to use an external script file instead of coding it all here

// RunLeft is a sample script we can use for actions
type RunLeft struct {
	Components.ScriptComponent
}

func (this *RunLeft) Initialize() {
	//no initialize for RunLeft
}

func (this *RunLeft) Update(delta float32) {

	this.Transform().X -= .7 //* delta
}

//////////////////////////////////////////////////////

func main() {
	// NOTE: we are going to have to set some ground rules for animations.
	// 1) the user will be able to append their own images (pngs) to an animation
	// 2) the user will be able to splice a sprite sheet into one or more animations
	//    i) the user will provide frame dimensions to splice a single sprite, then the number of frames to splice
	//    ii) animations are assumed to be one row within the sprite sheet (for now?)
	//        i.e. amin.SpliceAndSetSheet(imgName, frameX, frameY, noOfFrames, rowNum)
	//
	// NOTE: since we may be using multiple animations we may want to abstract an animation component to attach and map to the renderer
	//

	fmt.Println("starting")

	// create the foundation: a new window, basic scene, and game component (node)
	GT.EngineStart()
	simpleScene, _ := Graphics.NewBasicScene()
	node := Components.NewNode("new_image")

	// create our renderer and our animation objects
	rend := Components.NewSpriteRenderer()
	anim := Components.NewSpriteAnimation()
	anim2 := Components.NewSpriteAnimation()
	anim3 := Components.NewSpriteAnimation()

	// set or append our animation(s) based on an image, and user defined framing / splice logic
	anim.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 1)
	anim.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 6, 2)
	anim2.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 5, 3)
	anim2.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 4)
	anim3.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 3, 5)
	anim3.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 6)

	// if you can make use of a full sprite sheet for one animation, then use the SpliceAndSetFullSheet method
	//anim2.SpliceAndSetFullSheetAnimation(GT.AssetsImages+"tomatohead1.png", 12, 12)

	// you can append other animations or images
	anim.AppendAnimation(anim2)
	anim.AppendAnimation(anim3)
	anim.AppendImage(GT.AssetsImages + "test.png")
	anim.Frequency(0.25, false) // set frequency of toggle to every 15 frames

	// you can swap images in the animation, or remove them by number in the list
	anim.ReorderImage(1, 27) // one-based input
	anim.RemoveImage(1)      // remove test.png...

	rend.AddAnimation(anim, "dog_left_facing")
	rend.SetCurrentAnimation("dog_left_facing") // script components can make use of this...
	//simpleScene.SetFPS(30)

	// attach the sprite to our node, and transform if desired
	node.AddComponent(rend)
	node.AddComponent(&RunLeft{}) // TODO: remove this for a scripting example
	//node.Transform().Scale(3, 3) // NOTE: scaling doesn't seem perfect for sub images/animations?
	node.Transform().Translate(400, 400)

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)

	// start the scene to render our setup
	simpleScene.Start()

}
