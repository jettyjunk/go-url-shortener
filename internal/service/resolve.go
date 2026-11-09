package service

import "context"

func (s *LinkService) Resolve(ctx context.Context, code string) (string, error) {
	resCode, err := s.linkRepository.GetCode(ctx, code)
	if err != nil {
		return "", err
	}

	return resCode, nil
}
