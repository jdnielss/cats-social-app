package model

import "time"

type User struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email" validate:"required,email,min=5,max=50"`
	Name      string    `json:"name,omitempty" validate:"required,min=5,max=50"`
	Password  string    `json:"password" validate:"required,min=5,max=15"`
	CreatedAt time.Time `json:"createdAt,omitempty" validate:"required,time"`
}
