package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	UserID       int    `gorm:"type:SMALLINT(30)  NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	EmailAddress string `gorm:"type:varchar(75)"  json:"EmailAddress"`
	Password     string `gorm:"type:varchar(60)"  json:"Password"`
	RoomNum      string `gorm:"type:varchar(4)"   json:"RoomNum"`
	IsAdmin      bool   `gorm:"type:varchar(5)"   json:"IsAdmin"`
	IsActive     bool   `gorm:"type:varchar(5)"   json:"IsActive"`
	Wallet       uint64 `gorm:"type:SMALLINT(30)" json:"Wallet"`
	FirstName    string `gorm:"type:varchar(45)"  json:"FirstName"`
	LastName     string `gorm:"type:varchar(45)"  json:"LastName"`
	PhoneNum     string `gorm:"type:varchar(11)"  json:"PhoneNum"`
}

type Order struct {
	OrderID        int       `gorm:"type:SMALLINT(30) NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	CustomerID     int       `gorm:"type:int(11)" json:"CustomerID"`
	CourierID      int       `gorm:"type:int(11)" json:"CourierID"`
	DeliveryStatus string    `gorm:"type:varchar(15);default:'pending'" json:"DeliveryStatus"`
	TimeDelivered  time.Time `gorm:"type:timestamp;default:null" json:"TimeDelivered"`
	TimeCreated    time.Time `gorm:"type:timestamp;DEFAULT:CURRENT_TIMESTAMP" json:"TimeCreated"`
	OrderTotal     uint64    `gorm:"type:SMALLINT(30)" json:"OrderTotal"`
}

func (p *User) Activate() {
	p.IsActive = true
}

func (p *User) Deactivate() {
	p.IsActive = false
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Order{})
	return db
}
