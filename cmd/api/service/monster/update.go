package monster

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/rules"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func (s *_Service) Update(monsterID int, req *entity.MonsterPayload) (*entity.Monster, error) {
	var (
		timeNow = time.Now().Local()
		payl    = entity.Monster{}
	)

	dataMonster, err := s.repo.MonsterRepo.GetIDByMonsterID(monsterID)
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
		if _, err := s.repo.MonsterCategoryRepo.GetId(req.MonsterCategoryID); err != nil {
			return nil, err
		}
		payl.MonsterCategoryID = req.MonsterCategoryID
	}

	if req.Description == "" {
		payl.Description = dataMonster.Description
	} else {
		payl.Description = req.Description
	}

	uploadComplete := make(chan *string)
	defer close(uploadComplete)

	if req.Image == "" {
		payl.Image = dataMonster.Image
	} else {
		if isValid := rules.IsValidBase64(req.Image); !isValid {
			return nil, entity.CustomError("Image is not valid!")
		}

		mType := getFileExtension(req.Image)

		if isValid := rules.IsAllowedImageExtension(mType); !isValid {
			return nil, entity.CustomError("Extension is not valid!")
		}

		decoded, err := base64.StdEncoding.DecodeString(req.Image)
		if err != nil {
			return nil, err
		}

		objectName := fmt.Sprintf("%s%s.%s", s.cfg.GCS.Storage.Prefix, uuid.New(), mType)

		// Use a channel to communicate completion of the upload

		go func() {
			err := s.gcs.Put(context.Background(), s.cfg.GCS.Storage.Bucket, objectName, decoded, false, mType)
			if err != nil {
				log.Printf("error storedata to gcs", err)
				uploadComplete <- nil
				return
			}

			uploadComplete <- &objectName
		}()
		select {
		case imageStore := <-uploadComplete:
			if imageStore != nil {
				payl.Image = *imageStore
			} else {
				// Handle the error if the upload fails
				return nil, entity.CustomError("error storing data to GCS")
			}
		case <-time.After(5 * time.Second): // Add a timeout to avoid waiting indefinitely
			log.Println("Upload operation timed out")
			return nil, entity.CustomError("upload operation timed out")
		}

	}

	if len(req.TypesID) == 0 {
		var typesIDs []int
		for _, v := range dataMonster.Types {
			typesIDs = append(typesIDs, v.ID)
		}
		payl.TypesID = typesIDs
	} else {
		if _, err := s.repo.MonsterTypeRepo.GetIds(req.TypesID); err != nil {
			return nil, err
		}
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

	removeRedisData(context.Background(), s.rdb, fmt.Sprintf("%s%d", entity.MonsterRedisKey, monsterID))

	return &payl, nil
}

func removeRedisData(ctx context.Context, rdb *redis.Client, keys ...string) {
	rdb.Del(ctx, keys...)
}
