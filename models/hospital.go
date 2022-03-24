package models

type Hospital struct {
	HospitalId string `json:"_id" bson:"_id"`
	Username   string
	Name       string
	Address    string
	Phone      string
	Password   string
}

type HospitalRegistrationInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type HospitalLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type HospitalDetailsOutput struct {
	HospitalId string `json:"hospitalId"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}
