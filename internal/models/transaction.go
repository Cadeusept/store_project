package models

import "time"

type Transaction struct {
	Id        uint64    `json:"id" db:"id"`
	UserId    uint64    `json:"userId" db:"user_id" binding:"required"`
	UserEmail string    `json:"userEmail" db:"user_email" binding:"required"`
	Amount    uint64    `json:"amount" db:"amount" binding:"required"`
	Currency  string    `json:"currency" db:"currency" binding:"required"`
	Created   time.Time `json:"created" db:"created"`
	Changed   time.Time `json:"changed" db:"changed"`
	Status    string    `json:"status" db:"stat"`
}
