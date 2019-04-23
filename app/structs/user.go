package structs

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID           int64      `db:"user_id"       json:"user_id"`
	TimeCreated  *time.Time `db:"time_created"  json:"omitempty"`
	EmailAddress string     `db:"email_address" json:"email_address"`
	Password     string     `db:"password"      json:"password"`
	RoomNum      string     `db:"room_num"      json:"room_num"`
	IsAdmin      bool       `db:"is_admin"      json:"is_admin"`
	IsActive     bool       `db:"is_active"     json:"is_active"`
	Wallet       float32    `db:"wallet"        json:"wallet"`
	FirstName    string     `db:"first_name"    json:"first_name"`
	LastName     string     `db:"last_name"     json:"last_name"`
	PhoneNum     string     `db:"phone_number"     json:"phone_number"`
}
