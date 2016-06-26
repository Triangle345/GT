package Physics

import (
	"math"
)

type vector3 struct {
	X, Y, Z float64
}

// NewVector3 returns a new vector3 instance
func NewVector3(x, y, z float64) vector3 {
	return vector3{x, y, z}
}

// Add adds two vectors
func (this vector3) Add(other vector3) vector3 {
	return vector3{
		this.X + other.X,
		this.Y + other.Y,
		this.Z + other.Z}
}

// Add adds two vectors
func (this vector3) Sub(other vector3) vector3 {
	return vector3{
		this.X - other.X,
		this.Y - other.Y,
		this.Z - other.Z}
}

// Magreturns the length of the vector
func (this vector3) Mag() float64 {
	return math.Sqrt(math.Pow(this.X, 2) + math.Pow(this.Y, 2) + math.Pow(this.Z, 2))
}

// SqMag returns the length of the vector(no sqrt)
func (this vector3) SqMag() float64 {
	return math.Pow(this.X, 2) + math.Pow(this.Y, 2) + math.Pow(this.Z, 2)
}

// Scale returns the matrix multiplied by the scaler
func (this vector3) Scale(scalar float64) vector3 {
	return vector3{this.X * scalar, this.Y * scalar, this.Z * scalar}
}
