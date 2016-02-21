package Components

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
	children []GameNode
	Name     string
}

func (this Node) InitializeAll() {
	for _, child := range this.children {
		child.InitializeAll()
	}
}

func (this Node) UpdateAll(delta float32) {
	for _, child := range this.children {
		child.UpdateAll(delta)
	}
}

func (this *Node) addNode(node GameNode) {
	this.children = append(this.children, node)
}
