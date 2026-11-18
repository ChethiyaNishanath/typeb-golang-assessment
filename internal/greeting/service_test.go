package greeting

import (
	"errors"
	"testing"
)

func TestGreet_ValidLowercaseName(t *testing.T) {
	tt := struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		name:    "valid lowercase name",
		input:   "akila",
		want:    "Hello akila",
		wantErr: nil,
	}

	got, err := NewService().Greet(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Greet(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
	if got != tt.want {
		t.Errorf("%s: Greet(%q) = %q, want %q", tt.name, tt.input, got, tt.want)
	}
}

func TestGreet_ValidUppercaseName(t *testing.T) {
	tt := struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		name:    "valid uppercase name keeps its casing",
		input:   "Malindu",
		want:    "Hello Malindu",
		wantErr: nil,
	}

	got, err := NewService().Greet(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Greet(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
	if got != tt.want {
		t.Errorf("%s: Greet(%q) = %q, want %q", tt.name, tt.input, got, tt.want)
	}
}

func TestGreet_SingleValidLetter(t *testing.T) {
	tt := struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		name:    "single valid letter",
		input:   "K",
		want:    "Hello K",
		wantErr: nil,
	}

	got, err := NewService().Greet(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Greet(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
	if got != tt.want {
		t.Errorf("%s: Greet(%q) = %q, want %q", tt.name, tt.input, got, tt.want)
	}
}

func TestGreet_EmptyName(t *testing.T) {
	tt := struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		name:    "empty name returns empty name error",
		input:   "",
		want:    "",
		wantErr: ErrEmptyName,
	}

	got, err := NewService().Greet(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Greet(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
	if got != tt.want {
		t.Errorf("%s: Greet(%q) = %q, want %q", tt.name, tt.input, got, tt.want)
	}
}

func TestGreet_NameNotInFirstHalf(t *testing.T) {
	tt := struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		name:    "name in second half returns error",
		input:   "Zeebra",
		want:    "",
		wantErr: ErrNameNotInFirstHalf,
	}

	got, err := NewService().Greet(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Greet(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
	if got != tt.want {
		t.Errorf("%s: Greet(%q) = %q, want %q", tt.name, tt.input, got, tt.want)
	}
}

func TestGreet_FirstInvalidLetterN(t *testing.T) {
	tt := struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		name:    "boundary letter n returns error",
		input:   "nalinda",
		want:    "",
		wantErr: ErrNameNotInFirstHalf,
	}

	got, err := NewService().Greet(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Greet(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
	if got != tt.want {
		t.Errorf("%s: Greet(%q) = %q, want %q", tt.name, tt.input, got, tt.want)
	}
}

func TestGreet_StartsWithDigit(t *testing.T) {
	tt := struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		name:    "name starting with digit returns error",
		input:   "1chethiya",
		want:    "",
		wantErr: ErrNameNotInFirstHalf,
	}

	got, err := NewService().Greet(tt.input)
	if !errors.Is(err, tt.wantErr) {
		t.Errorf("%s: Greet(%q) error = %v, want %v", tt.name, tt.input, err, tt.wantErr)
	}
	if got != tt.want {
		t.Errorf("%s: Greet(%q) = %q, want %q", tt.name, tt.input, got, tt.want)
	}
}
