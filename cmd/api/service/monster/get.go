package monster

import (
	"context"
	"os"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
	"github.com/DitoAdriel99/go-monsterdex/pkg/jwt_parse"
	"github.com/DitoAdriel99/go-monsterdex/pkg/meta"
	"github.com/DitoAdriel99/go-monsterdex/pkg/storage"
)

func (s *_Service) Get(bearer string, m *meta.Metadata) (*presentation.Monsters, error) {
	claims, err := jwt_parse.GetClaimsFromToken(bearer)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.MonsterRepo.Get(claims.ID, m)
	if err != nil {
		return nil, err
	}

	urlCh := make(chan string)
	errCh := make(chan error)
	defer close(urlCh)
	defer close(errCh)

	for i := range *data {
		go func(i int) {
			url, err := storage.SignedURL(context.Background(), os.Getenv("GCS_BUCKET"), (*data)[i].Image)
			if err != nil {
				errCh <- err
				return
			}
			urlCh <- *url
		}(i)
	}

	for i := range *data {
		select {
		case url := <-urlCh:
			(*data)[i].Image = url
		case err := <-errCh:
			return nil, err
		}
	}

	return data, nil
}
