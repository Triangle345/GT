package Components

// ScriptComponent becomes a user defined script
// Contains helper methods
type ScriptComponent struct {
	ChildComponent
}

// Transform - returns the parent's Transform
func (this ScriptComponent) Transform() *Transform {
	return &this.GetParent().transform
}
