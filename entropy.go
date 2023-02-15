package pwcheck

import (
	"math"
)

// GetEntropy returns the entropy in bits for the given password
// See the ReadMe for more information
func GetEntropy(password string) float64 {
	return getEntropy(password)
}

func getEntropy(password string) float64 {
	base := GetBase(password)
	length := GetLength(password)

	// calculate log2(base^length)
	return LogPow(float64(base), length, 2)
}

func logX(base, n float64) float64 {
	if base == 0 {
		return 0
	}
	// change of base formulae
	return math.Log2(n) / math.Log2(base)
}

// logPow calculates log_base(x^y)
// without leaving logspace for each multiplication step
// this makes it take less space in memory
func LogPow(expBase float64, pow int, logBase float64) float64 {
	// logb (MN) = logb M + logb N
	total := 0.0
	for i := 0; i < pow; i++ {
		total += logX(logBase, expBase)
	}
	return total
}

// GetEntropySlice returns the entropy in bits for each password in a slice
// See the ReadMe for more information
func GetEntropySlice(passwords []string) []float64 {
	entropy := make([]float64, len(passwords))
	for i, password := range passwords {
		entropy[i] = getEntropy(password)
	}
	return entropy
}
