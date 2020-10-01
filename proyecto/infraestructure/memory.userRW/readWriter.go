package userRW

import (
	"errors"
	"sync"

	"log"

	"time"

	"example.com/mittaus/ddd-example/application"
	"example.com/mittaus/ddd-example/domain"
	"example.com/mittaus/ddd-example/testData"
	"github.com/spf13/viper"
)

type rw struct {
	store *sync.Map // map username:user
	//TODO : ADD a password hasher here
}

func New() application.UserRW {
	rw := rw{
		store: &sync.Map{},
	}

	if viper.GetBool("populate") {
		log.Print("append user")
		rick := testData.User("rick")
		rw.Create(rick.Name, rick.Email, rick.Password)
	}

	return rw
}

func (rw rw) Create(username, email, password string) (*domain.User, error) {
	if _, err := rw.GetByName(username); err == nil {
		return nil, application.ErrAlreadyInUse
	}

	rw.store.Store(username, domain.User{
		Name:      username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return rw.GetByName(username)
}

func (rw rw) GetByName(userName string) (*domain.User, error) {
	value, ok := rw.store.Load(userName)
	if !ok {
		return nil, application.ErrNotFound
	}

	user, ok := value.(domain.User)
	if !ok {
		return nil, errors.New("not a user stored at key")
	}

	return &user, nil
}

func (rw rw) GetByEmailAndPassword(email, password string) (*domain.User, error) {
	var err error
	var foundUser domain.User

	rw.store.Range(func(key, value interface{}) bool {
		user, ok := value.(domain.User)
		if !ok {
			err = errors.New("failed to assert to domain.User")
			return false
		}

		if user.Email == email && user.Password == password {
			foundUser = user
			return false // stop range
		}

		return true // keep iterating
	})

	return &foundUser, err
}

func (rw rw) Save(user domain.User) error {
	if user, _ := rw.GetByName(user.Name); user == nil {
		return application.ErrNotFound
	}

	user.UpdatedAt = time.Now()
	rw.store.Store(user.Name, user)

	return nil
}
