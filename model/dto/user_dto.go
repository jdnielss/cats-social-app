package dto

type RegistrationRequestDTO struct {
	Email    string `json:"email" validate:"required,email,min=5,max=50"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type RegistrationResponseDTO struct {
	Email       string `json:"email" validate:"required,email,min=5,max=50"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
	AccessToken string `json:"accessToken"`
}
