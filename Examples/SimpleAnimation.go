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

	fmt.Println("starting")

	// create the foundation: a new window, basic scene, and game component (node)
	GT.EngineStart()
	simpleScene, _ := Graphics.NewBasicScene()
	node := Components.NewNode("new_image")

	// create our renderer and our animation objects
	rend := Components.NewSpriteRenderer()

	manager := rend.AnimationManager // will be replaced by GetComponent(AnimationManager) when scripting etc.

	anim := Components.SpliceSpriteSheetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 1)

	// you can append other animations or images
	anim.AppendAnimation(Components.SpliceSpriteSheetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 6, 2))
	anim.AppendAnimation(Components.SpliceSpriteSheetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 5, 3))
	anim.AppendAnimation(Components.SpliceSpriteSheetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 4))
	anim.AppendAnimation(Components.SpliceSpriteSheetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 3, 5))
	anim.AppendAnimation(Components.SpliceSpriteSheetAnimation(GT.AssetsImages+"Dog.png", 90, 58, 4, 6))

	anim.Append(GT.AssetsImages + "test.png")

	// if you can make use of a full sprite sheet for one animation, then use the SpliceFullSpriteSheetAnimation method
	//anim2.SpliceFullSpriteSheetAnimation(GT.AssetsImages+"tomatohead1.png", 12, 12)

	anim.SetFrequency(0.25, false) // set frequency of toggle to every 1/4 second
	// anim.SetFrequency(10, true) // set frequency of toggle to every 10 frames

	// you can swap images in the animation, or remove them by number in the list
	anim.Reorder(1, 27) // one-based input
	anim.Remove(1)      // remove test.png...

	manager.AddAnimation(anim, "dog_left_facing")
	manager.SetCurrentAnimation("dog_left_facing") // script components can make use of this...

	// animations can be retrieved for further configuration if desired
	anim = manager.GetAnimation("dog_left_facing")
	anim.SetAsOneTimeOnly(true)
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
