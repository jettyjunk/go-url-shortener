package service

import (
	"context"
	"testing"
)

type stubRepo struct {
	createErrs []error
	calls      int
}

func (s *stubRepo) Create(ctx context.Context, code string, originalURL string) error {
	err := s.createErrs[s.calls]
	s.calls++
	return err
}

func (s *stubRepo) GetCode(ctx context.Context, code string) (string, error) {
	return "", nil
}

func TestShortenSuccess(t *testing.T) {
	repo := &stubRepo{
		createErrs: []error{nil},
	}

	svc := NewLinkService(repo)

	code, err := svc.Shorten(context.Background(), "htps://example.com/")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(code) != codeLength {
		t.Errorf("expected length %d, got %d", codeLength, len(code))
	}

	if repo.calls != 1 {
		t.Fatalf("expected 1 create call, got %d", repo.calls)
	}

}
