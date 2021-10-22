package models

import "time"

type Transaction struct {
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time
}

type Transaction_Response struct {
	ID      int
	Product []Transaction_Response
}

type Transaction_Input struct {
	Courier   string `json:"courier" form:"courier"`
	ProductID []int  `json:"product_id" form:"product_id"`
}

