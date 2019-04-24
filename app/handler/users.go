package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dsafanyuk/fetchr-go/app/model"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func GetAllUsers(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	user := []model.User{}
	db.Select(&user, "SELECT * FROM users")
	respondJSON(w, http.StatusOK, user)
}

func CreateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	createdUser := model.User{}
	query := `
	INSERT INTO users (
		email_address, password, room_num, first_name, last_name, phone_number
	) VALUES (
		:email_address, :password, :room_num, :first_name, :last_name, :phone_number
	) RETURNING user_id, email_address, password, room_num, first_name, last_name, phone_number, is_active, is_admin, time_created, wallet`

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	rows, err := db.NamedQuery(query, &user)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer rows.Close()
	// Retreive newly created user_id
	if rows.Next() {
		err := rows.StructScan(&createdUser)
		if err != nil {
			log.Fatal(err)
		}
	}
	respondJSON(w, http.StatusCreated, createdUser)
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

func UpdateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	updatedUser := model.User{}
	query := `
	UPDATE users SET
		email_address = :email_address,
		password = :password,
		wallet = :wallet,
		is_active = :is_active,
		is_admin = :is_admin,
		room_num = :room_num,
		first_name = :first_name,
		last_name = :last_name,
		phone_number = :phone_number
	WHERE user_id = :user_id
	RETURNING user_id, email_address, password, room_num, first_name, last_name, phone_number, is_active, is_admin, time_created, wallet`

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

	rows, err := db.NamedQuery(query, &user)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Retreive newly created user_id
	if rows.Next() {
		err := rows.StructScan(&updatedUser)
		if err != nil {
			log.Fatal(err)
			respondError(w, http.StatusInternalServerError, err.Error())
		}
	}
	respondJSON(w, http.StatusOK, updatedUser)
}

func DeleteUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	query := `
	UPDATE users
		SET is_active = false
	WHERE user_id = :user_id
		RETURNING user_id, email_address, password, room_num, first_name, last_name, phone_number, is_active, is_admin, time_created, wallet`

	user := getUserOr404(db, userID, w, r)
	if user == nil {
		return
	}
	rows, err := db.NamedQuery(query, user)
	if rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
			respondError(w, http.StatusInternalServerError, err.Error())
		}
	}
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getUserOr404(db *sqlx.DB, userID string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	err := db.Get(&user, "SELECT * FROM users WHERE user_id = $1", userID)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
