package monster

import (
	"time"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/labstack/gommon/log"
)

func (s *_Service) Update(monsterID int, req *entity.MonsterPayload) (*entity.Monster, error) {
	var (
		timeNow = time.Now().Local()
		payl    = entity.Monster{}
	)

	dataMonster, err := s.repo.MonsterRepo.GetIDAll(1, monsterID)
	if err != nil {
		log.Printf("error getting id: %v", err)
		return nil, err
	}

	payl.ID = dataMonster.ID

	if req.Name == "" {
		payl.Name = dataMonster.Name
	} else {
		payl.Name = req.Name
	}

	if req.MonsterCategoryID == 0 {
		payl.MonsterCategoryID = dataMonster.MonsterCategory.ID
	} else {
		payl.MonsterCategoryID = req.MonsterCategoryID
	}

	if req.Description == "" {
		payl.Description = dataMonster.Description
	} else {
		payl.Description = req.Description
	}

	if len(req.TypesID) == 0 {
		var typesIDs []int
		for _, v := range dataMonster.Types {
			typesIDs = append(typesIDs, v.ID)
		}
		payl.TypesID = typesIDs
	} else {
		payl.TypesID = req.TypesID
	}

	if req.Height == 0.0 {
		payl.Height = dataMonster.Height
	} else {
		payl.Height = req.Height
	}

	if req.Weight == 0.0 {
		payl.Weight = dataMonster.Weight
	} else {
		payl.Weight = req.Weight
	}

	if req.StatsHP == 0 {
		payl.StatsHP = dataMonster.StatsHP
	} else {
		payl.StatsHP = req.StatsHP
	}

	if req.StatsAttack == 0 {
		payl.StatsAttack = dataMonster.StatsAttack
	} else {
		payl.StatsAttack = req.StatsAttack
	}

	if req.StatsDefense == 0 {
		payl.StatsDefense = dataMonster.StatsDefense
	} else {
		payl.StatsDefense = req.StatsDefense
	}

	if req.StatsSpeed == 0 {
		payl.StatsSpeed = dataMonster.StatsSpeed
	} else {
		payl.StatsSpeed = req.StatsSpeed
	}

	payl.CreatedAt = dataMonster.CreatedAt
	payl.UpdatedAt = timeNow

	if err := s.repo.MonsterRepo.Update(monsterID, &payl); err != nil {
		log.Printf("error update monster: %v", err)
		return nil, err
	}

	return &payl, nil
}
