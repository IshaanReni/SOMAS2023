package modules

import (
	"SOMAS2023/internal/common/utils"
	"math"
)

type ForceVector struct {
	X float64
	Y float64
}

func (fv *ForceVector) Dot(vec ForceVector) float64 {
	return fv.X*vec.X + fv.Y*vec.Y
}

func (fv *ForceVector) Magnitude() float64 {
	return math.Sqrt(fv.X*fv.X + fv.Y*fv.Y)
}

func (fv *ForceVector) CosineSimilarity(vec ForceVector) float64 {
	return fv.Dot(vec) / (fv.Magnitude() * vec.Magnitude())
}

func GetForceVector(force utils.Forces) *ForceVector {
	return &ForceVector{
		X: force.Pedal * float64(math.Cos(float64(math.Pi*force.Turning.SteeringForce))),
		Y: force.Pedal * float64(math.Sin(float64(math.Pi*force.Turning.SteeringForce))),
	}
}
