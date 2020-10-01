package domain

//UserRW repositorio
type UserRW interface {
	Create(username, email, password string) (*User, error)
	GetByName(userName string) (*User, error)
	GetByEmailAndPassword(email, password string) (*User, error)
	Save(user User) error
}
