package user

import (
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID           int64      `db:"user_id"       json:"user_id,omitempty"`
	TimeCreated  *time.Time `db:"time_created"  json:"time_created,omitempty"`
	EmailAddress string     `db:"email_address" json:"email_address,omitempty"`
	Password     string     `db:"password"      json:"password,omitempty"`
	RoomNum      string     `db:"room_num"      json:"room_num,omitempty"`
	IsAdmin      bool       `db:"is_admin"      json:"is_admin,omitempty"`
	IsActive     bool       `db:"is_active"     json:"is_active,omitempty"`
	Wallet       uint64     `db:"wallet"        json:"wallet,omitempty"`
	FirstName    string     `db:"first_name"    json:"first_name,omitempty"`
	LastName     string     `db:"last_name"     json:"last_name,omitempty"`
	PhoneNum     string     `db:"phone_number"     json:"phone_number,omitempty"`
}
