package Components

import (
	"fmt"
)

func NewNode(name string) Node {
	return Node{Transform: NewTransform(), Name: name}
}

type GameNode interface {
	InitializeAll()
	UpdateAll(delta float32)
}

type Component interface {
	Initialize()
	Update(delta float32)
}

type Node struct {
	Transform
	children   []GameNode
	components []Component
	Name       string
}

func (this Node) InitializeAll() {
	for _, child := range this.children {
		child.InitializeAll()
	}
}

func (this Node) UpdateAll(delta float32) {

	for _, component := range this.components {
		component.Update(delta)
	}

	for _, child := range this.children {
		child.UpdateAll(delta)
	}
}

// TODO make sure each node only has one of each type
func (this *Node) AddNode(node GameNode) {
	if n, ok := node.(Node); ok {
		n.Transform.model = this.GetUpdatedModel()
	} else {
		fmt.Printf("Cannot add node.\n")
	}

	this.children = append(this.children, node)
}

func (this *Node) AddComponent(component Component) {
	this.components = append(this.components, component)
}
