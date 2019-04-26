package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func ensureUserTableExists() {
	if _, err := a.DB.Exec(userTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearUserTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_user_id_seq RESTART WITH 1")
}

const userTableCreationQuery = `create table if not exists test_db.public.users
(
  user_id       serial              not null
    constraint users_pk
      primary key,
  time_created  timestamp default CURRENT_TIMESTAMP,
  email_address varchar(75)         not null,
  password      varchar(256)        not null,
  is_admin      boolean   default false,
  wallet        integer   default 0 not null,
  first_name    varchar(75),
  last_name     varchar(75),
  is_active     boolean   default true,
  phone_number  varchar(10),
  room_num      numeric(4)
);
`

func TestEmptyUserTable(t *testing.T) {
	clearUserTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	clearUserTable()

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "sql: no rows in result set" {
		t.Errorf("Expected the 'error' key of the response to be set to 'sql: no rows in result set'. Got '%s'", m["error"])
	}
}
func TestCreateUser(t *testing.T) {
	clearUserTable()

	payload := []byte(`{
	"email_address":"dave@email.com",
	"password": "david1",
	"room_num": "1234",
	"first_name": "davy",
	"last_name": "safanyuk",
	"phone_number": "123456789"
	}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["email_address"] != "dave@email.com" {
		t.Errorf("Expected user name to be 'dave@email.com'. Got '%v'", m["email_address"])
	}

	if m["password"] != "david1" {
		t.Errorf("Expected user password to be 'david1'. Got '%v'", m["password"])
	}

	if m["room_num"] != "1234" {
		t.Errorf("Expected user room_num to be '1234'. Got '%v'", m["room_num"])
	}
}

func TestGetUser(t *testing.T) {
	clearUserTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
func TestUpdateUser(t *testing.T) {
	clearUserTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	payload := []byte(`{"first_name":"Not Davy"}`)

	req, _ = http.NewRequest("PUT", "/users/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["user_id"] != originalUser["user_id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalUser["user_id"], m["user_id"])
	}

	if m["first_name"] == originalUser["first_name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalUser["first_name"], m["first_name"], m["first_name"])
	}
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		statement := fmt.Sprintf(`
		INSERT INTO users (
			email_address, password, room_num, first_name, last_name, phone_number
		) VALUES (
			%s, %s, %s, %s, %s, %s
		)`, "'dave@email.com'", "'david1'", "'1234'", "'davy'", "'safanyuk'", "'1234567890'")
		a.DB.Exec(statement)
	}
}
