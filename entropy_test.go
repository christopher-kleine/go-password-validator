package pwcheck_test

import (
	"math"
	"testing"

	"github.com/christopher-kleine/pwcheck"
)

func TestLogPow(t *testing.T) {
	expected := math.Round(math.Log2(math.Pow(7, 8)))
	actual := math.Round(pwcheck.LogPow(7, 8, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log2(math.Pow(10, 11)))
	actual = math.Round(pwcheck.LogPow(10, 11, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log2(math.Pow(11, 17)))
	actual = math.Round(pwcheck.LogPow(11, 17, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log10(math.Pow(13, 21)))
	actual = math.Round(pwcheck.LogPow(13, 21, 10))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
