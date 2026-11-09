package codegen

import (
	"strings"
	"testing"
)

func TestGenerateLength(t *testing.T) {
	code, err := Generate(7)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(code) != 7 {
		t.Errorf("expected lenght 7, got %d", len(code))
	}
}

func TestGenerateAlphabet(t *testing.T) {
	code, err := Generate(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, r := range code {
		if !strings.ContainsRune(alphabet, r) {
			t.Fatalf("expected character: %q", r)
		}
	}
}

func TestGenerateUnique(t *testing.T) {
	seen := make(map[string]struct{})

	for i := 0; i < 1000; i++ {
		code, err := Generate(7)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if _, ok := seen[code]; ok {
			t.Fatalf("duplicate code: %q", code)
		}

		seen[code] = struct{}{}
	}
}

func TestGenerateInvalidLength(t *testing.T) {
	if _, err := Generate(0); err == nil {
		t.Errorf("expected error for length 0")
	}
}
