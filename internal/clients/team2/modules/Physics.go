package modules

import "math"

type ForceVector struct {
	X float64
	Y float64
}

func (fv *ForceVector) dot(vec ForceVector) float64 {
	return fv.X*vec.X + fv.Y*vec.Y
}

func (fv *ForceVector) magnitude() float64 {
	return math.Sqrt(fv.X*fv.X + fv.Y*fv.Y)
}

func (fv *ForceVector) cosineSimilarity(vec ForceVector) float64 {
	return fv.dot(vec) / (fv.magnitude() * vec.magnitude())
}
