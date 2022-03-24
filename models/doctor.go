package models

type Doctor struct {
	DoctorId string `json:"_id" bson:"_id"`
	Username string
	Name     string
	Phone    string
	Password string
	UPRN     string
}

type DoctorRegistrationInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	UPRN     string `json:"uprn" binding:"required"`
}

type DoctorLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DoctorLoginOutput struct {
	DoctorId string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	UPRN     string `json:"uprn"`
}

type DoctorDetailsOutput struct {
	DoctorId string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}
