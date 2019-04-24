package model

import (
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ID          int        `db:"product_id" json:"product_id"`
	ProductName string     `db:"product_name" json:"product_name"`
	Price       uint64     `db:"price" json:"price"`
	Category    string     `db:"category" json:"category"`
	ProductURL  string     `db:"product_url" json:"product_url"`
	TimeCreated *time.Time `db:"time_created" json:"time_created"`
	IsActive    bool       `db:"is_active" json:"is_active"`
}
