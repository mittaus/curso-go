package application

import (
	"example.com/mittaus/ddd-example/domain"
)

func (i interactor) ArticlesFeed(username string, limit, offset int) (*domain.User, domain.ArticleCollection, int, error) {
	if limit < 0 {
		return nil, domain.ArticleCollection{}, 0, nil
	}

	var user *domain.User
	if username != "" {
		var errGet error
		user, errGet = i.userRW.GetByName(username)
		if errGet != nil {
			return nil, nil, 0, errGet
		}
	}
	articles, err := i.articleRW.GetByAuthorsNameOrderedByMostRecentAsc(user.FollowIDs)
	if err != nil {
		return nil, nil, 0, err
	}

	return user, domain.ArticleCollection(articles).ApplyLimitAndOffset(limit, offset), len(articles), nil // needs the original length
}
