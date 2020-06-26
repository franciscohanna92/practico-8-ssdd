package inmem

import (
	"fmt"
	"sync"

	user "../../../pkg"
)

type userRepository struct {
	mtx   sync.RWMutex
	users map[string]*user.User
}

func NewUserRepository(users map[string]*user.User) user.UserRepository {
	if users == nil {
		users = make(map[string]*user.User)
	}

	return &userRepository{
		users: users,
	}
}

func (r *userRepository) CreateUser(g *user.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if err := r.checkIfExists(g.Id); err != nil {
		return err
	}
	r.users[g.Id] = g
	return nil
}

func (r *userRepository) FetchUsers() ([]*user.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*user.User, 0, len(r.users))
	for _, value := range r.users {
		values = append(values, value)
	}
	return values, nil
}

func (r *userRepository) DeleteUser(Id string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.users, Id)

	return nil
}

func (r *userRepository) UpdateUser(Id string, g *user.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.users[Id] = g
	return nil
}

func (r *userRepository) FetchUserById(Id string) (*user.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.users {
		if v.Id == Id {
			return v, nil
		}
	}

	return nil, fmt.Errorf("The Id %s doesn't exist", Id)
}

func (r *userRepository) checkIfExists(Id string) error {
	for _, v := range r.users {
		if v.Id == Id {
			return fmt.Errorf("The user %s already exists", Id)
		}
	}

	return nil
}
