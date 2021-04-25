package articlesrv

import (
	"context"

	"github.com/armin-agahian/hexagonal-sample/internal/core/application/ports"
	"github.com/armin-agahian/hexagonal-sample/internal/core/domain/entities"
	"github.com/google/uuid"
)

type service struct {
	articleRepository ports.ArticleRepository
}

func New(articleRepo ports.ArticleRepository) ports.ArticleService {
	return &service{articleRepository: articleRepo}
}

func (srv *service) Get(ctx context.Context, id string) (entities.Article, error) {
	article, err := srv.articleRepository.Get(ctx, id)
	if err != nil {
		article = entities.Article{}
	}
	return article, err
}

func (srv *service) Create(ctx context.Context, title string, body string) (entities.Article, error) {
	article := entities.NewArticle(uuid.New().String(), title, body)
	err := srv.articleRepository.Save(ctx, article)
	if err != nil {
		article = entities.Article{}
	}
	return article, err
}
