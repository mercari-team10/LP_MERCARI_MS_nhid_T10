package models

type Patient struct {
	NHID         string `json:"_id" bson:"_id"`
	Username     string
	Name         string
	Phone        string
	Password     string
	AadharNumber string
	DateOfBirth  int64
	Gender       string
}

type PatientRegistrationInput struct {
	Name         string `json:"name" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Password     string `json:"password" binding:"required"`
	AadharNumber string `json:"aadhar" binding:"required"`
	DateOfBirth  int64  `json:"dob" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
}

type PatientLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PatientLoginOutput struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	NHID         string `json:"nhid"`
	Phone        string `json:"phone"`
	AadharNumber string `json:"aadhar"`
	DateOfBirth  int64  `json:"dob"`
	Gender       string `json:"gender"`
}
