package Components

type Component interface {
	Child
	Initialize()
	Update(delta float32)
}

type ChildComponent struct {
	Parent *Node
}

func (this *ChildComponent) SetParent(node *Node) {
	this.Parent = node
}

func (this *ChildComponent) GetParent() *Node {
	return this.Parent
}

func (this *ChildComponent) Intialize() {
	//nothing
}

func (this *ChildComponent) Update(delta float32) {
	// nothing
}
