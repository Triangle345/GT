package Physics

import (
	"GT/Logging"
	"math"
)

type vector3 struct {
	X, Y, Z float64
}

// NewVector3 returns a new vector3 instance
func NewVector3(x, y, z float64) vector3 {
	return vector3{x, y, z}
}

// NewVector3 returns a new vector3 instance
func ZeroVector3() vector3 {
	return vector3{0, 0, 0}
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

// Comp is the Component Product a o b
func (this vector3) Comp(other vector3) vector3 {
	return vector3{this.X * other.X, this.Y * other.Y, this.Z * other.Z}
}

// Scalar returns the matrix multiplied by the scaler
func (this vector3) Scalar(scalar float64) vector3 {
	return vector3{this.X * scalar, this.Y * scalar, this.Z * scalar}
}

// Dot is the dot prodcut or scalar product of two vectors a . b
func (this vector3) Dot(other vector3) float64 {
	return this.X*other.X + this.Y*other.Y + this.Z*other.Z
}

// Cross returns cross product matrix of two matrices, eg. The orthogonal matrix
func (this vector3) Cross(other vector3) vector3 {
	return vector3{
		this.Y*other.Z - this.Z*other.Y,
		this.Z*other.X - this.X*other.Z,
		this.X*other.Y - this.Y*other.X}
}

// Normalizes normalizes the vector values to be in 0.0 - 1.0 range
func (this vector3) Normalize() vector3 {
	mag := this.Mag()

	if mag > 0 {
		return this.Scalar(1 / mag)
	}

	Logging.Debug("Cannot normalize vector, magnitude 0")
	return ZeroVector3()
}

// Lerp Linear Interpolation p0+(d/∥v∥)v
func (this vector3) Lerp(weight float64) vector3 {
	mag := this.Mag()
	wLen := mag * weight

	return this.Scalar(wLen / mag)
}
