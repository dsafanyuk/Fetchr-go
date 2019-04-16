package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dsafanyuk/fetchr-go/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, _ := strconv.ParseInt(vars["userID"], 0, 64)
	fmt.Println(userID)
	user := getUserOr404(db, userID, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, _ := strconv.ParseInt(vars["userID"], 0, 64)
	user := getUserOr404(db, userID, w, r)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, _ := strconv.ParseInt(vars["userID"], 0, 64)
	user := getUserOr404(db, userID, w, r)
	if user == nil {
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, userID int64, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, userID).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
