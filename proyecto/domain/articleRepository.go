package domain

type ArticleRW interface {
	Create(Article) (*Article, error)
	Save(Article) (*Article, error)
	GetBySlug(slug string) (*Article, error)
	GetByAuthorsNameOrderedByMostRecentAsc(usernames []string) ([]Article, error)
	GetRecentFiltered(filters []ArticleFilter) ([]Article, error)
	Delete(slug string) error
}
