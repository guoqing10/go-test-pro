package main

import (
	"errors"
	"fmt"
	internal "go-test-pro/intenal"
)

type User struct {
	ID   int
	Name string
}

func NewUser(id int, name string) (User, error) {
	if name == "" {
		return User{}, errors.New("name is empty")
	}
	return User{ID: id, Name: name}, nil
}

type UserRepo interface {
	Save(User)
	Get(int) (User, bool)
}

type UserService struct{ repo UserRepo }

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(id int, name string) error {
	u, err := NewUser(id, name)
	if err != nil {
		return err
	}
	s.repo.Save(u)
	return nil
}

func (s *UserService) Find(id int) (User, bool) {
	return s.repo.Get(id)
}

type MemUserRepo struct{ data map[int]User }

func NewMemUserRepo() *MemUserRepo {
	return &MemUserRepo{data: map[int]User{}}
}

func (r *MemUserRepo) Save(u User) {
	r.data[u.ID] = u
}

func (r *MemUserRepo) Get(id int) (User, bool) {
	u, ok := r.data[id]
	return u, ok
}

func main() {
	repo := NewMemUserRepo()
	svc := NewUserService(repo)
	_ = svc.Register(1, "tom")
	u, _ := svc.Find(1)
	internal.Service()
	fmt.Println(u)
}
