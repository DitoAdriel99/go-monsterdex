package monster

import (
	"context"
	"os"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
	"github.com/DitoAdriel99/go-monsterdex/pkg/jwt_parse"
	"github.com/DitoAdriel99/go-monsterdex/pkg/storage"
)

func (s *_Service) GetID(bearer string, monsterID int) (*presentation.Monster, error) {
	claims, err := jwt_parse.GetClaimsFromToken(bearer)
	if err != nil {
		return nil, err
	}

	data, err := s.repo.MonsterRepo.GetID(claims.ID, monsterID)
	if err != nil {
		return nil, err
	}

	url, err := storage.SignedURL(context.Background(), os.Getenv("GCS_BUCKET"), data.Image)
	if err != nil {
		return nil, err
	}

	data.Image = *url

	return data, nil
}
