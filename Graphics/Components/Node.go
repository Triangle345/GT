package Components

import (
	"fmt"
	"reflect"
)

func NewNode(name string) *Node {
	return &Node{transform: NewTransform(), Name: name}
}

type Child interface {
	SetParent(node *Node)
	GetParent() *Node
}

// type GameNode interface {
// 	Component
// 	Transform() *Transform
// 	AddNode(node GameNode)
// 	AddComponent(component Component)
// 	GetComponent(componentType string) Component
// }

type Node struct {
	transform Transform
	ChildComponent
	children   []*Node
	components []Component
	Name       string
}

func (this *Node) Transform() *Transform {
	return &this.transform
}

func (this *Node) Initialize() {
	for _, child := range this.children {
		child.Initialize()
	}
}

func (this *Node) Update(delta float32) {
	// if n, ok := this.Parent.(*Node); ok {

	// need to check if this is the top most node
	if this.Parent != nil {
		this.transform.model = this.Parent.transform.GetUpdatedModel()
	}

	for _, component := range this.components {
		component.Update(delta)
	}

	for _, child := range this.children {
		child.Update(delta)
	}
}

func (this *Node) AddNode(node *Node) {

	// if n, ok := node.(Child); ok {

	node.SetParent(this)
	// } else {
	// 	fmt.Printf("No parent to set for child node: %s.\n", this.Name)
	// }

	this.children = append(this.children, node)

	node.Initialize()
}

func (this *Node) AddComponent(component Component) {
	if n, ok := component.(Child); ok {

		n.SetParent(this)
	} else {
		fmt.Printf("No parent to set for child component.\n")
	}

	this.components = append(this.components, component)

	component.Initialize()
}

func (this *Node) GetComponent(componentType string) Component {
	for _, c := range this.components {
		fmt.Println("type:")
		fmt.Println(reflect.TypeOf(c).Elem().Name())

		cType := reflect.TypeOf(c).Elem().Name()

		// need to handle defer here as reflect may throw panic if type not a pointer
		defer func() {
			cType = reflect.TypeOf(c).Name()
		}()

		if cType == componentType {
			return c
		}
	}

	fmt.Println("Cannot find component: " + componentType)
	return nil
}
