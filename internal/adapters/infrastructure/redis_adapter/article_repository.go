package redis_adapter

import (
	"context"
	"encoding/json"
	"time"

	"github.com/armin-agahian/hexagonal-sample/internal/core/application/ports"
	"github.com/armin-agahian/hexagonal-sample/internal/core/domain/entities"
	"github.com/go-redis/redis/v8"
)

type repository struct {
	rdb *redis.Client
}

func NewArticleRepository() ports.ArticleRepository {
	newClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &repository{
		rdb: newClient,
	}
}

func (repo *repository) Get(ctx context.Context, id string) (entities.Article, error) {
	article := entities.Article{}
	value, err := repo.rdb.Get(ctx, id).Bytes()
	if err == nil {
		json_err := json.Unmarshal(value, &article)
		if json_err != nil {
			return entities.Article{}, json_err
		}
	}
	return article, err
}

func (repo *repository) Save(ctx context.Context, article entities.Article) error {
	value, err := json.Marshal(article)
	if err == nil {
		redis_err := repo.rdb.Set(ctx, article.Id, value, time.Hour).Err()
		return redis_err
	}
	return err
}
