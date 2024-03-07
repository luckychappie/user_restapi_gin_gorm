package model

import (
	"time"
)

type User struct {
	Id      *string    `json:"id"`
	Name    *string    `json:"name"`
	Email   *string    `json:"email"`
	Address *string    `json:"address"`
	Dob     *time.Time `json:"dob"`
}
