package pwcheck

import "strings"

var (
	ReplaceChars      = `!@$&*`
	SepChars          = `_-., `
	OtherSpecialChars = `"#%'()+/:;<=>?[\]^{|}~`
	LowerChars        = `abcdefghijklmnopqrstuvwxyz`
	UpperChars        = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	DigitsChars       = `0123456789`
)

func GetBase(password string) int {
	chars := map[rune]struct{}{}
	for _, c := range password {
		chars[c] = struct{}{}
	}

	hasReplace := false
	hasSep := false
	hasOtherSpecial := false
	hasLower := false
	hasUpper := false
	hasDigits := false
	base := 0

	for c := range chars {
		switch {
		case strings.ContainsRune(ReplaceChars, c):
			hasReplace = true
		case strings.ContainsRune(SepChars, c):
			hasSep = true
		case strings.ContainsRune(OtherSpecialChars, c):
			hasOtherSpecial = true
		case strings.ContainsRune(LowerChars, c):
			hasLower = true
		case strings.ContainsRune(UpperChars, c):
			hasUpper = true
		case strings.ContainsRune(DigitsChars, c):
			hasDigits = true
		default:
			base++
		}
	}

	if hasReplace {
		base += len(ReplaceChars)
	}
	if hasSep {
		base += len(SepChars)
	}
	if hasOtherSpecial {
		base += len(OtherSpecialChars)
	}
	if hasLower {
		base += len(LowerChars)
	}
	if hasUpper {
		base += len(UpperChars)
	}
	if hasDigits {
		base += len(DigitsChars)
	}
	return base
}
