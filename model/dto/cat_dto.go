package dto

import (
	"github.com/google/uuid"
)

type CatRequestDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required,min=1,max=30"`
	Race        string    `json:"race" validate:"required,oneof=Persian MaineCoon Siamese Ragdoll Bengal Sphynx BritishShorthair Abyssinian ScottishFold Birman"`
	Sex         string    `json:"sex" validate:"required,oneof=male female"`
	AgeInMonth  int       `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string    `json:"description" validate:"required,min=1,max=200"`
	ImageUrls   []string  `json:"imageUrls" validate:"required,min=1,dive,url"`
	HasMatched  bool      `json:"hasMatched"`
}

type CatFilter struct {
	ID               string // `id` limit the output based on the cat’s id
	Limit            int    // `limit` limit the output of the data
	Offset           int    // `offset` offset the output of the data
	Race             string // `race` filter based on race
	Sex              string // `sex` filter based on sex
	IsAlreadyMatched bool   // `isAlreadyMatched` filter based on match
	AgeInMonth       string // `ageInMonth` filter based on age
	Owned            bool   // `owned` filter based on the cat that the user own
	Search           string // `search` display information that contains the name of search
}
