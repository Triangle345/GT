package Components

type Component interface {
	Child
	Initialize()
	Update(delta float32)
}

type GameComponent struct {
	Parent GameNode
}

func (this *GameComponent) SetParent(node GameNode) {
	this.Parent = node
}

func (this *GameComponent) GetParent() GameNode {
	return this.Parent
}

func (this *GameComponent) Intialize() {
	//nothing
}

func (this *GameComponent) Update(delta float32) {
	// nothing
}
