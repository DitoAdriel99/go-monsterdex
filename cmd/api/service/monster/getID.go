package monster

import (
	"context"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
)

func (s *_Service) GetID(bearer string, monsterID int) (*presentation.Monster, error) {
	claims, err := s.token.GetClaimsFromToken(bearer)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.MonsterRepo.GetID(claims.ID, monsterID)
	if err != nil {
		return nil, err
	}

	url, err := s.gcs.Get(context.Background(), s.cfg.GCS.Storage.Bucket, data.Image)
	if err != nil {
		return nil, err
	}

	data.Image = string(url)

	return data, nil
}
