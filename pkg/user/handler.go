package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return &userHandler{
		userService,
	}
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.FindAllUsers()
	if err != nil {
		logrus.WithField("error", err).Error("Unable to find all users")
		http.Error(w, "Unable to find all users", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}

func (h *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if `id` is a valid user id
	if _, err := strconv.Atoi(id); err != nil {
		logrus.WithFields(logrus.Fields{"id": id}).Error("Invalid User ID")
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	user, err := h.userService.FindUserByID(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Unable to find user")
		http.Error(w, "Unable to find user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {

	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		logrus.Error("Unable to decode user")
		http.Error(w, "Bad format for user", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		logrus.WithField("error", err).Error("Unable to create user")
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(createdUser)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}

}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if `id` is a valid user id
	if _, err := strconv.Atoi(id); err != nil {
		logrus.WithFields(logrus.Fields{"id": id}).Error("Invalid User ID")
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	var payload User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		logrus.Error("Unable to decode user")
		http.Error(w, "Bad format for user", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.userService.UpdateUser(id, &payload)
	response, err := json.Marshal(updatedUser)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}

}
func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := h.userService.DeleteUser(id)

	response, err := json.Marshal(user)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}
