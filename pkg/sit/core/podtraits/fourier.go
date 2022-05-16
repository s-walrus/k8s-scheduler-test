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
	for i, b := range f.cosKs {
		y += b * math.Cos((float64)(i+1)*x)
	}
	return y
}

func (f *FiniteFourierSeries) Integrate(a float64, b float64) float64 {
	var sinKs, cosKs []float64
	for i, a := range f.sinKs {
		cosKs = append(cosKs, -a/float64(i+1))
	}
	for i, b := range f.cosKs {
		sinKs = append(sinKs, b/float64(i+1))
	}
	intFourier := NewFiniteFourierSeries(0, sinKs, cosKs)
	return intFourier.GetValue(b) + f.k*b - intFourier.GetValue(a) - f.k*a
}
