package user

type User struct {
	Id    string `json:"id"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
}

type UserRepository interface {
	CreateUser(g *User) error
	FetchUsers() ([]*User, error)
	DeleteUser(Id string) error
	FetchUserById(Id string) (*User, error)
}
