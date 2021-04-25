package ports

import (
	"context"

	"github.com/armin-agahian/hexagonal-sample/internal/core/domain/entities"
)

type ArticleRepository interface {
	Get(ctx context.Context, id string) (entities.Article, error)
	Save(ctx context.Context, article entities.Article) error
}
