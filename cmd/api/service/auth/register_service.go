package auth

import (
	"github.com/DitoAdriel99/go-monsterdex/pkg/hashing"

	"log"
	"time"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
)

func (s *_Service) Register(payload *entity.RegisterPayload) (*entity.Register, error) {
	var (
		time = time.Now().Local()
	)

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.AuthRepo.CheckEmail(payload.Email); err != nil {
		log.Println("Check Email error : ", err)
		return nil, err
	}

	hashedPass, err := hashing.HashPassword(payload.Password)
	if err != nil {
		log.Println("Hash Password error : ", err)
		return nil, err
	}

	dataRegister := entity.Register{
		FullName:  payload.FullName,
		Email:     payload.Email,
		Password:  hashedPass,
		Role:      "user",
		CreatedAt: time,
		UpdatedAt: &time,
	}

	currId, err := s.repo.AuthRepo.Register(&dataRegister)
	if err != nil {
		log.Println("Register error : ", err)
		return nil, err
	}

	dataRegister.ID = currId

	return &dataRegister, nil
}
