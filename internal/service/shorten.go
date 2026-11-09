package service

import (
	"context"
	"errors"
	"net/url"

	"github.com/jettyjunk/go-url-shortener/internal/codegen"
	"github.com/jettyjunk/go-url-shortener/internal/repository"
)

const (
	codeLength  = 7
	maxAttempts = 5
)

func (s *LinkService) Shorten(ctx context.Context, rawURL string) (string, error) {
	origURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", ErrInvalidURL
	}
	if origURL.Scheme != "http" && origURL.Scheme != "https" {
		return "", ErrInvalidURL
	}
	if origURL.Host == "" {
		return "", ErrInvalidURL
	}

	for i := 0; i < maxAttempts; i++ {
		code, err := codegen.Generate(codeLength)
		if err != nil {
			return "", err
		}

		err = s.linkRepository.Create(ctx, code, rawURL)
		if err == nil {
			return code, nil
		}

		if errors.Is(err, repository.ErrCodeExists) {
			continue
		}

		return "", err
	}

	return "", ErrCodeGeneration
}
