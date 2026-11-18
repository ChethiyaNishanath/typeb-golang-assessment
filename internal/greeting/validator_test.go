package greeting

import (
	"errors"
	"testing"
)

func TestValidate_LowercaseFirstLetterA(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "lowercase first letter a",
		input:   "akila",
		wantErr: nil,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_LowercaseLastValidLetterM(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "lowercase last valid letter m",
		input:   "malindu",
		wantErr: nil,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_UppercaseFirstLetterA(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "uppercase first letter A",
		input:   "Akila",
		wantErr: nil,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_UppercaseFirstLetterF(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "uppercase first letter A",
		input:   "Fahim",
		wantErr: nil,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_UppercaseLastValidLetterM(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "uppercase last valid letter M",
		input:   "Malindu",
		wantErr: nil,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_SingleValidUppercaseLetter(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "single valid letter",
		input:   "K",
		wantErr: nil,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_SingleValidLowercaseLetter(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "single valid letter",
		input:   "l",
		wantErr: nil,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_EmptyName(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "empty name",
		input:   "",
		wantErr: ErrEmptyName,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_LowercaseFirstInvalidLetterN(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "lowercase first invalid letter n",
		input:   "nalinda",
		wantErr: ErrNameNotInFirstHalf,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_LowercaseZ(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "lowercase z",
		input:   "zeebra",
		wantErr: ErrNameNotInFirstHalf,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_UppercaseN(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "uppercase N",
		input:   "Nalinda",
		wantErr: ErrNameNotInFirstHalf,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_UppercaseZ(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "uppercase Z",
		input:   "Zeebra",
		wantErr: ErrNameNotInFirstHalf,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_StartsWithDigit(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "starts with digit",
		input:   "1chethiya",
		wantErr: ErrNameNotInFirstHalf,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_StartsWithSpace(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "starts with space",
		input:   " chethiya",
		wantErr: ErrNameNotInFirstHalf,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}

func TestValidate_StartsWithSymbol(t *testing.T) {
	tt := struct {
		name    string
		input   string
		wantErr error
	}{
		name:    "starts with symbol",
		input:   "@chethiya",
		wantErr: ErrNameNotInFirstHalf,
	}

	err := ValidateName(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Validate(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
}
