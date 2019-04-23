package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dsafanyuk/fetchr-go/app/structs"
	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
)

func GetAllUsers(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	user := []structs.User{}
	db.Select(&user, "SELECT * FROM users ORDER BY first_name ASC")
	respondJSON(w, http.StatusOK, user)
}

func CreateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	user := structs.User{}
	query := `
	INSERT INTO users (
		email_address, password, room_num, first_name, last_name, phone_number
	) VALUES (
		:email_address, :password, :room_num, :first_name, :last_name, :phone_number
	)`
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	result, err := db.NamedExec(query, &user)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	user.ID, err = result.LastInsertId()
	respondJSON(w, http.StatusCreated, user)
}

func GetUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID := vars["userID"]

	user := getUserOr404(db, userID, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

// func UpdateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	userID, _ := strconv.ParseInt(vars["userID"], 0, 64)
// 	user := getUserOr404(db, userID, w, r)
// 	if user == nil {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&user); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Create(&user).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, user)
// }

func DeleteUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	query := `
	UPDATE users
		SET is_active = 0
	WHERE user_id = ?
	`
	checkUser := getUserOr404(db, userID, w, r)
	if checkUser == nil {
		return
	}

	_, err := db.Exec(query, userID)

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, userID)
}

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getUserOr404(db *sqlx.DB, userID string, w http.ResponseWriter, r *http.Request) *structs.User {
	user := structs.User{}
	err := db.Get(&user, "SELECT * FROM users WHERE user_id = ?", userID)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
