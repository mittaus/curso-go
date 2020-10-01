package articleValidator

import (
	"example.com/mittaus/ddd-example/application"
	"example.com/mittaus/ddd-example/domain"
)

type validator struct {
}

func New() application.ArticleValidator {
	return validator{}
}

func (validator) BeforeCreationCheck(article *domain.Article) error { return nil }
func (validator) BeforeUpdateCheck(article *domain.Article) error   { return nil }
