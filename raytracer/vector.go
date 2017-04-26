package raytracer

import "math"

type Vector struct {
	X, Y, Z float32
}

func times(v Vector, k float32) Vector {
	return Vector{k * v.X, k * v.Y, k * v.Z}
}

func minus(v1, v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func dot(v1, v2 Vector) float32 {
	return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func mag(v Vector) float32 {
	formula := v.X * v.X + v.Y * v.Y + v.Z * v.Z
	sqrt := math.Sqrt(float64(formula))
	return float32(sqrt)
}

func norm(v Vector) Vector {
	var vmag = mag(v)
	var div float32

	if (vmag == 0.0) {
		div = math.MaxFloat32
	} else {
		div = 1 * vmag
	}
	return times(v, div)
}

func cross(v1, v2 Vector) Vector {
	return Vector{v1.Y * v2.Z - v1.Z * v2.Y, v1.Z * v2.X - v1.X * v2.Z, v1.X * v2.Y - v1.Y * v2.X}
}
