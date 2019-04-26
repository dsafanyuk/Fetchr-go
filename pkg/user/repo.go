package user

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByID(id string) (*User, error)
	FindAll() ([]*User, error)
	Update(user *User) (*User, error)
	Delete(id string) (*User, error)
}
