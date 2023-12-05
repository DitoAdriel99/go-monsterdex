package dependency

import (
	"context"
	"fmt"
	"log"

	"github.com/DitoAdriel99/go-monsterdex/bootstrap"
	"github.com/DitoAdriel99/go-monsterdex/config"
	midd "github.com/DitoAdriel99/go-monsterdex/middleware"
	"github.com/DitoAdriel99/go-monsterdex/pkg/storage"
	"github.com/DitoAdriel99/go-monsterdex/pkg/tokenizer"
	"github.com/go-redis/redis/v8"
)

type Dependency struct {
	Redis     *redis.Client
	GcsClient storage.Storage
	Token     tokenizer.JWT
	RBAC      midd.RBAC
}

func NewDependency(cfg config.Cfg) *Dependency {
	ctx := context.Background()
	gcsClient, err := storage.NewGCS(ctx, cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("connect to GCS client %v", err))
	}

	redisCLient := bootstrap.NewRedisClient(cfg)

	token := tokenizer.JWT{
		Cfg:    cfg,
		Issuer: "monsterdex-app",
	}

	rbac := midd.RBAC{
		Token: token,
	}

	return &Dependency{
		Redis:     redisCLient,
		GcsClient: gcsClient,
		Token:     token,
		RBAC:      rbac,
	}

}
