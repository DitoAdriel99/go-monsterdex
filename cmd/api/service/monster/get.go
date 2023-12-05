package monster

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
	"github.com/DitoAdriel99/go-monsterdex/pkg/meta"
	"github.com/go-redis/redis/v8"
)

func (s *_Service) Get(bearer string, m *meta.Metadata) (*presentation.Monsters, error) {
	claims, err := s.token.GetClaimsFromToken(bearer)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.MonsterRepo.Get(claims.ID, m)
	if err != nil {
		return nil, err
	}

	urlCh := make(chan string)
	errCh := make(chan error)

	var wg sync.WaitGroup
	for i := range *data {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ctx := context.Background()
			urlRedis, err := checkRedisData(ctx, s.rdb, fmt.Sprintf("%s%d", entity.MonsterRedisKey, (*data)[i].ID))
			if err == redis.Nil {
				url, err := s.gcs.ResignUrl(ctx, s.cfg.GCS.Storage.Bucket, (*data)[i].Image)
				if err != nil {
					errCh <- err
					return
				}
				urlCh <- url
				chacheData(ctx, s.rdb, fmt.Sprintf("%s%d", entity.MonsterRedisKey, (*data)[i].ID), url)
			} else {
				urlCh <- urlRedis
			}

		}(i)
	}

	go func() {
		wg.Wait()
		close(urlCh)
	}()

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

func checkRedisData(ctx context.Context, rdb *redis.Client, key string) (string, error) {
	data, err := rdb.Get(ctx, key).Result()
	return data, err
}

func chacheData(ctx context.Context, rdb *redis.Client, key string, content string) {
	rdb.Set(ctx, key, content, time.Minute*5)
}
