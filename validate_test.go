package pwcheck_test

import (
	"testing"

	"github.com/christopher-kleine/pwcheck"
)

func TestValidate(t *testing.T) {
	err := pwcheck.Validate("mypass", 50)
	expectedError := "insecure password, try including more special characters, using uppercase letters, using numbers or using a longer password"
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}

	err = pwcheck.Validate("MYPASS", 50)
	expectedError = "insecure password, try including more special characters, using lowercase letters, using numbers or using a longer password"
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}

	err = pwcheck.Validate("mypassword", 4)
	if err != nil {
		t.Errorf("Err should be nil")
	}

	err = pwcheck.Validate("aGoo0dMi#oFChaR2", 80)
	if err != nil {
		t.Errorf("Err should be nil")
	}

	expectedError = "insecure password, try including more special characters, using lowercase letters, using uppercase letters or using a longer password"
	err = pwcheck.Validate("123", 60)
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}
}
