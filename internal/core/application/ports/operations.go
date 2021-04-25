package ports

import (
	"context"

	"github.com/armin-agahian/hexagonal-sample/internal/core/domain/entities"
)

type ArticleService interface {
	Get(ctx context.Context, id string) (entities.Article, error)
	Create(ctx context.Context, title string, body string) (entities.Article, error)
}
