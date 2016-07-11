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

	fmt.Println("starting")

	// create the foundation: a new window, basic scene, and game component (node)
	GT.EngineStart()
	simpleScene, _ := Graphics.NewBasicScene()
	node := Components.NewNode("new_image")

	// create our renderer and our animation objects
	rend := Components.NewSpriteRenderer()

	anim := rend.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 1)

	// you can append other animations or images
	anim.AppendAnimation(rend.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 6, 2))
	anim.AppendAnimation(rend.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 5, 3))
	anim.AppendAnimation(rend.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 4))
	anim.AppendAnimation(rend.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 3, 5))
	anim.AppendAnimation(rend.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 6))

	anim.Append(GT.AssetsImages + "test.png")

	// if you can make use of a full sprite sheet for one animation, then use the SpliceAndSetFullSheet method
	//anim2.SpliceAndSetFullSheetAnimation(GT.AssetsImages+"tomatohead1.png", 12, 12)

	anim.SetFrequency(0.25, false) // set frequency of toggle to every 1/4 second

	// you can swap images in the animation, or remove them by number in the list
	anim.Reorder(1, 27) // one-based input
	anim.Remove(1)      // remove test.png...

	anim.SetAsOneTimeOnly(true)

	rend.AnimationHandler.AddAnimation(anim, "dog_left_facing")
	rend.AnimationHandler.SetCurrentAnimation("dog_left_facing") // script components can make use of this...
	//simpleScene.SetFPS(30)

	// attach the sprite to our node, and transform if desired
	node.AddComponent(rend)
	node.AddComponent(&RunLeft{}) // TODO: remove this for a scripting example
	//node.Transform().Scale(3, 3) // NOTE: scaling doesn't seem perfect for sub images/animations?
	node.Transform().Translate(400, 400, 0)

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)

	// start the scene to render our setup
	simpleScene.Start()

}
