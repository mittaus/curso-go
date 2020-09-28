package DIP

import (
	"fmt"
)

type User struct {
	ID   uint
	Name string
}

type IUserRepository interface {
	Search(id int) User
}

type UserRepository struct {
	listaUsuarios []User
}

func NewUserRepository() *UserRepository {
	repo := new(UserRepository)
	repo.listaUsuarios = []User{{1, "Julio"}, {2, "Juan"}}
	return repo
}

func (u UserRepository) Search(id uint) User {
	var usuario User
	for i := 0; i < len(u.listaUsuarios); i++ {
		if u.listaUsuarios[i].ID == id {
			usuario = u.listaUsuarios[i]
			break
		}
	}

	return usuario
}

type IUserService interface {
	Search(id int) User
}

type UserService struct {
	repository *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
	return &UserService{repository: r}
}
func (s UserService) Search(id uint) User {
	return s.repository.Search(id)
}

func Run() {
	fmt.Println("Dependency Inversion principle (DIP)")
	repositorio := NewUserRepository()
	servicio := NewUserService(repositorio)
	usuario := servicio.Search(1)
	fmt.Println(usuario)
	fmt.Println("........................................................")
}
