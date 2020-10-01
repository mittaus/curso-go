package commentRW

import (
	"sync"

	"errors"

	"time"

	"example.com/mittaus/ddd-example/domain"
	"example.com/mittaus/ddd-example/application"
)

type rw struct {
	store *sync.Map
}

func New() application.CommentRW {
	return rw{
		store: &sync.Map{},
	}
}

func (rw rw) Create(comment domain.Comment) (*domain.Comment, error) {
	if _, err := rw.GetByID(comment.ID); err == nil {
		return nil, application.ErrAlreadyInUse
	}
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	rw.store.Store(comment.ID, comment)

	return rw.GetByID(comment.ID)
}

func (rw rw) GetByID(id int) (*domain.Comment, error) {
	value, ok := rw.store.Load(id)
	if !ok {
		return nil, application.ErrNotFound
	}

	comment, ok := value.(domain.Comment)
	if !ok {
		return nil, errors.New("not an article stored at key")
	}

	return &comment, nil
}

func (rw rw) Delete(id int) error {
	rw.store.Delete(id)

	return nil
}
