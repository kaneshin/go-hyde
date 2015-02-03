package hyde

type HydeFunc func(float64) float64

const hydeUnit = 156.0

func FromHyde(v float64) float64 {
	return v * hydeUnit
}

func ToHyde(v float64) float64 {
	return v / hydeUnit
}
