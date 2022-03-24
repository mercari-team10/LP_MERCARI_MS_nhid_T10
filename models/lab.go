package models

type Lab struct {
	LabId      string `json:"_id" bson:"_id"`
	Username   string
	Name       string
	Address    string
	HospitalId string
	Phone      string
	Password   string
}

type LabRegistrationInput struct {
	Name       string `json:"name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Password   string `json:"password" binding:"required"`
	HospitalId string `json:"hospitalId" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

type LabLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LabDetailsOutput struct {
	LabId      string `json:"labId"`
	HospitalId string `json:"hospitalId"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}
