package monster

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/pkg/storage"
	"github.com/google/uuid"
)

func (s *_Service) Create(req *entity.MonsterPayload) (*entity.Monster, error) {
	var (
		timeNow = time.Now().Local()
	)

	if err := req.Validate(); err != nil {
		return nil, err
	}

	if _, err := s.repo.MonsterCategoryRepo.GetId(req.MonsterCategoryID); err != nil {
		return nil, err
	}

	if _, err := s.repo.MonsterTypeRepo.GetIds(req.TypesID); err != nil {
		return nil, err
	}

	data := entity.Monster{
		Name:              req.Name,
		MonsterCategoryID: req.MonsterCategoryID,
		Description:       req.Description,
		TypesID:           req.TypesID,
		Image:             req.Image,
		Height:            req.Height,
		Weight:            req.Weight,
		StatsHP:           req.StatsHP,
		StatsAttack:       req.StatsAttack,
		StatsDefense:      req.StatsDefense,
		StatsSpeed:        req.StatsSpeed,
		IsActive:          true,
		CreatedAt:         timeNow,
		UpdatedAt:         timeNow,
	}

	mType := getFileExtension(data.Image)

	decoded, err := base64.StdEncoding.DecodeString(data.Image)
	if err != nil {
		return nil, err
	}

	objectName := fmt.Sprintf("%s%s.%s", os.Getenv("GCS_PREFIX"), uuid.New(), mType)

	// Use a channel to communicate completion of the upload
	uploadComplete := make(chan *string)
	defer close(uploadComplete)

	go func() {
		imageStore, err := storage.StoreDataToGCS(context.Background(), os.Getenv("GCS_BUCKET"), objectName, decoded, false, mType)
		if err != nil {
			log.Printf("error storedata to gcs", err)
			uploadComplete <- nil
			return
		}

		uploadComplete <- imageStore
	}()

	select {
	case imageStore := <-uploadComplete:
		if imageStore != nil {
			data.Image = *imageStore
		} else {
			// Handle the error if the upload fails
			return nil, errors.New("error storing data to GCS")
		}
	case <-time.After(5 * time.Second): // Add a timeout to avoid waiting indefinitely
		log.Println("Upload operation timed out")
		return nil, errors.New("upload operation timed out")
	}

	currId, err := s.repo.MonsterRepo.Create(1, &data)
	if err != nil {
		log.Printf("error create monster", err)
		return nil, err
	}

	signedUrl, err := storage.SignedURL(context.Background(), os.Getenv("GCS_BUCKET"), objectName)
	if err != nil {
		log.Printf("error signed to gcs", err)
		return nil, err
	}

	data.ID = currId
	data.Image = *signedUrl

	return &data, nil
}

func getFileExtension(base64String string) string {
	// Decode the base64 string
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return ""
	}

	// Magic numbers or headers of different file types
	magicNumbers := map[string]string{
		"\x89PNG":          "png",
		"BM":               "bmp",
		"GIF":              "gif",
		"\xFF\xD8\xFF":     "jpeg",
		"\x25\x50\x44\x46": "pdf",
		// Add more magic numbers for other file types if needed
	}

	// Iterate over magic numbers to find a match
	for magic, ext := range magicNumbers {
		if strings.HasPrefix(string(data), magic) {
			return ext
		}
	}

	// If no match found, return empty string
	return ""
}
