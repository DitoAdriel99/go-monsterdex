package monster

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/labstack/gommon/log"
)

func (s *_Service) SetStatus(monsterID int, req *entity.StatusPayload) error {
	if err := req.Validate(); err != nil {
		return err
	}
	dataMonster, err := s.repo.MonsterRepo.GetIDByMonsterID(monsterID)
	if err != nil {
		log.Printf("error getting id: %v", err)
		return err
	}

	if dataMonster.IsActive == req.Status {
		if req.Status {
			return entity.ErrAlreadyActive
		} else {
			return entity.ErrAlreadyDeactive
		}
	}

	if err := s.repo.MonsterRepo.SetStatus(monsterID, req.Status); err != nil {
		log.Printf("error set status: %v", err)
		return err
	}

	return nil
}
