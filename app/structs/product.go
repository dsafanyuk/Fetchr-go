package structs

import (
	"time"
)

type Product struct {
	ID          int
	ProductName string    `db:"product_name" json:"product_name"`
	Price       uint64    `db:"price" json:"price"`
	Category    string    `db:"category" json:"Category"`
	ProductURL  string    `db:"product_url" json:"ProductURL"`
	TimeCreated time.Time `db:"time_created" json:"-"`
	IsActive    bool      `db:"is_active" json:"IsActive"`
}
