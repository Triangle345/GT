// SimpleAnimation : This code demonstrates how to create a 2d animation from a sprite sheet

package main

import (
	"GT"
	"GT/Graphics"
	"GT/Graphics/Components"
	"fmt"
)

//////////////////////////////////////////////////////
// TODO: clean this up to use an external script file instead of coding it all here
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
	// NOTE/TODO: we are going to have to set some ground rules for animations.
	// 1) the user will be able to append their own images (pngs) to an animation
	//    i.e. rend.Animation[rowNum?].AppendImage(imgName)
	// 2) the user will be able to splice a sprite sheet into one or more animations
	//    i) the user will provide frame dimensions to splice a single sprite, then the number of frames to splice
	//    ii) animations are assumed to be one row within the sprite sheet (for now?)
	//        i.e. rend.SpliceAndSetSheet(imgName, frameX, frameY, noOfFrames, rowNum)
	//
	// NOTE: since we may be using multiple animations we may want to abstract an animation component to attach and map to the renderer
	//

	fmt.Println("starting")

	// create the foundation: a new window, basic scene, and game component (node)
	GT.EngineStart()
	simpleScene, _ := Graphics.NewBasicScene()
	node := Components.NewNode("new_image")

	// splice a sprite sheet into multiple nodes... method tbd

	rend := Components.NewSpriteRenderer()
	anim := Components.NewSpriteAnimation()

	// SetAnimation to grab specific frames from a row in a spritesheet
	anim.SpliceAndSetAnimation(GT.AssetsImages+"Dog.png", 58, 90, 5, 3)

	// SetFullSheetAnimation to use entire sprite sheet as one animation
	//anim.SpliceAndSetFullSheetAnimation(GT.AssetsImages+"tomatohead1.png", 12, 12)
	//anim.SpliceAndSetFullSheetAnimation(GT.AssetsImages+"Dog.png", 58, 90) //58, 90)
	anim.Frequency = 6

	rend.AddAnimation(anim, "run_left")
	rend.SetCurrentAnimation("run_left") // script components can make use of this...
	simpleScene.SetFPS(30)

	// attach the sprite to our node, and transform if desired
	node.AddComponent(rend)
	node.AddComponent(&RunLeft{})
	//node.Transform().Scale(3, 3)
	node.Transform().Translate(400, 400)

	// attach the node to our scene
	simpleScene.RootNode.AddNode(node)

	// start the scene to render our setup
	simpleScene.Start()

}
