package user

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(user *User) (*User, error)
	FindUserByID(id string) (*User, error)
	FindAllUsers() ([]*User, error)
	UpdateUser(id string, user *User) (*User, error)
	DeleteUser(id string) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo,
	}
}

func (s *userService) CreateUser(user *User) (*User, error) {
	user, err := s.repo.Create(user)
	if err != nil {
		logrus.WithField("error", err).Error("Error creating user")
		return nil, err
	}
	logrus.WithField("id", user.ID).Info("Created new user")

	return user, nil
}

func (s *userService) FindUserByID(id string) (*User, error) {
	user, err := s.repo.FindByID(id)

	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Error finding user")
		return nil, err
	}
	logrus.WithField("id", id).Info("Found user")

	return user, nil
}

func (s *userService) FindAllUsers() ([]*User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Error finding all users")
		return nil, err
	}
	return users, nil
}

func (s *userService) UpdateUser(id string, payload *User) (*User, error) {
	// Ensure requested user exists
	existingUser, err := s.repo.FindByID(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Error finding user")
		return nil, err
	}

	// Merge the payload and existing user struct to provide updated struct to  repo
	updatedUser := merge(existingUser, payload)
	user, err := s.repo.Update(updatedUser.(*User))
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Error updating user")
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id string) (*User, error) {
	user, err := s.repo.Delete(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Error deleting user")

		return nil, err
	}
	return user, nil
}

// merges two structs, where a's values take precendence over b's values (a's values will be kept over b's if each field has a value)
func merge(a, b interface{}) interface{} {
	jb, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Marshal error b:", err)
	}
	err = json.Unmarshal(jb, &a)
	if err != nil {
		fmt.Println("Unmarshal error b-a:", err)
	}

	return a
}
