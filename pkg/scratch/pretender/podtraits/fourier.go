package podtraits

import "math"

type FiniteFourierSeries struct {
	k     float64
	sinKs []float64
	cosKs []float64
}

func NewFiniteFourierSeries(k float64, sinCoefficients []float64, cosCoefficients []float64) *FiniteFourierSeries {
	return &FiniteFourierSeries{k: k, sinKs: sinCoefficients, cosKs: cosCoefficients}
}

func (f *FiniteFourierSeries) GetValue(x float64) float64 {
	y := f.k
	for i, a := range f.sinKs {
		y += a * math.Sin((float64)(i+1)*x)
	}
	for i, a := range f.sinKs {
		y += a * math.Cos((float64)(i+1)*x)
	}
	return y
}
