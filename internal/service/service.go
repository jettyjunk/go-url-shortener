package service

import "context"

type LinkService struct {
	linkRepository LinkRepository
}

type LinkRepository interface {
	Create(
		ctx context.Context,
		code string,
		originalURL string,
	) error

	GetCode(
		ctx context.Context,
		code string,
	) (string, error)
}

func NewLinkService(linkRepository LinkRepository) *LinkService {
	return &LinkService{
		linkRepository: linkRepository,
	}
}
