package pwcheck_test

import (
	"testing"

	"github.com/christopher-kleine/pwcheck"
)

func TestGetBase(t *testing.T) {
	actual := pwcheck.GetBase("abcd")
	expected := len(pwcheck.LowerChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetBase("abcdA")
	expected = len(pwcheck.LowerChars) + len(pwcheck.UpperChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetBase("A")
	expected = len(pwcheck.UpperChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetBase("^_")
	expected = len(pwcheck.OtherSpecialChars) + len(pwcheck.SepChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetBase("^")
	expected = len(pwcheck.OtherSpecialChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetBase("!")
	expected = len(pwcheck.ReplaceChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetBase("123")
	expected = len(pwcheck.DigitsChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetBase("123ü")
	expected = len(pwcheck.DigitsChars) + 1
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}
