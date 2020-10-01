package persistence

import (
	"errors"
	"strings"

	"time"

	"example.com/mittaus/ddd-example/application"
	"example.com/mittaus/ddd-example/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserRW {
	return &UserRepository{db}
}

func (rw UserRepository) Create(username, email, password string) (*domain.User, error) {

	if _, err := rw.GetByName(username); err == nil {
		return nil, application.ErrAlreadyInUse
	}

	rw.db.Create(domain.User{
		Name:      username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return rw.GetByName(username)
}

func (rw UserRepository) GetByName(userName string) (*domain.User, error) {
	var user domain.User
	err := rw.db.Where("name = ?", userName).Order("email ASC").Take(&user).Error
	if err != nil {
		return nil, err
	}

	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (rw UserRepository) GetByEmailAndPassword(email, password string) (*domain.User, error) {

	var user domain.User
	err := rw.db.Where("email=? AND password=?", email, password).Order("email ASC").Take(&user).Error
	if err != nil {
		return nil, err
	}

	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("failed to assert to domain.User")
	}

	return &user, nil

}

func (rw UserRepository) Save(user domain.User) error {
	if user, _ := rw.GetByName(user.Name); user == nil {
		return application.ErrNotFound
	}

	user.UpdatedAt = time.Now()
	err := rw.db.Save(user).Error

	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return errors.New("title already taken")
		}
		//any other db error
		return errors.New("database error")
	}

	return nil
}
