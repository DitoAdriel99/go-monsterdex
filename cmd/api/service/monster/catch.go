package monster

import (
	"log"
	"time"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/pkg/jwt_parse"
)

func (s *_Service) Catch(bearer string, monsterID int) (*bool, error) {
	var (
		timeNow = time.Now().Local()
	)

	claims, err := jwt_parse.GetClaimsFromToken(bearer)
	if err != nil {
		return nil, err
	}

	data := entity.Catch{
		UserID:    claims.ID,
		MonsterID: monsterID,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	isCatch, err := s.repo.MonsterRepo.Catch(&data)
	if err != nil {
		log.Printf("error catch: %v", err)
		return nil, err
	}

	return isCatch, nil
}
