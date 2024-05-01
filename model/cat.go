package model

import (
	"time"
)

type Cat struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name" validate:"required,min=1,max=30"`
	Race        string    `json:"race" validate:"required,oneof=Persian MaineCoon Siamese Ragdoll Bengal Sphynx BritishShorthair Abyssinian ScottishFold Birman"`
	Sex         string    `json:"sex" validate:"required,oneof=male female"`
	AgeInMonth  int       `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string    `json:"description" validate:"required,min=1,max=200"`
	ImageUrls   []string  `json:"imageUrls" validate:"required,min=1,dive,url"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt" validate:"required,time"`
}
