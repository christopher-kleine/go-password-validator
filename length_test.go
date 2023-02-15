package pwcheck_test

import (
	"testing"

	"github.com/christopher-kleine/pwcheck"
)

func TestRemoveMoreThanTwoFromSequence(t *testing.T) {
	actual := pwcheck.RemoveMoreThanTwoFromSequence("12345678", "0123456789")
	expected := "12"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.RemoveMoreThanTwoFromSequence("abcqwertyabc", "qwertyuiop")
	expected = "abcqwabc"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.RemoveMoreThanTwoFromSequence("", "")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.RemoveMoreThanTwoFromSequence("", "12345")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

func TestGetReversedString(t *testing.T) {
	actual := pwcheck.GetReversedString("abcd")
	expected := "dcba"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetReversedString("1234")
	expected = "4321"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

func TestRemoveRepeatingChars(t *testing.T) {
	actual := pwcheck.RemoveMoreThanTwoRepeatingChars("aaaa")
	expected := "aa"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.RemoveMoreThanTwoRepeatingChars("bbbbbbbaaaaaaaaa")
	expected = "bbaa"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.RemoveMoreThanTwoRepeatingChars("ab")
	expected = "ab"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.RemoveMoreThanTwoRepeatingChars("")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

func TestGetLength(t *testing.T) {
	actual := pwcheck.GetLength("aaaa")
	expected := 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetLength("11112222")
	expected = 4
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetLength("aa123456")
	expected = 4
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetLength("876543")
	expected = 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pwcheck.GetLength("qwerty123456z")
	expected = 5
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}
