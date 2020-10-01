package application

import (
	"errors"

	"example.com/mittaus/ddd-example/domain"
)

func (i interactor) getArticleAndCheckUser(username, slug string) (*domain.Article, error) {
	completeArticle, err := i.articleRW.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	if completeArticle == nil {
		return nil, errArticleNotFound
	}

	// check only if a username is specified
	if username != "" && completeArticle.Author.Name != username {
		return nil, errors.New("article not owned by user")
	}

	return completeArticle, nil
}
