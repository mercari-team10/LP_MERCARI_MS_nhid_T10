package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/dryairship/techmeet22-nhid-service/db"
	"github.com/dryairship/techmeet22-nhid-service/models"
)

func PatientLogin(c *gin.Context) {
	var patientDetails models.PatientLoginInput
	err := c.BindJSON(&patientDetails)
	if err != nil {
		log.Println("[WARN] Patient Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid login details")
		return
	}

	foundPatient, err := db.FindPatientByUsername(patientDetails.Username)
	if err != nil {
		log.Println("[WARN] Patient Username not found: ", patientDetails.Username, err.Error())
		c.String(http.StatusNotFound, "Invalid Username")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundPatient.Password), []byte(patientDetails.Password))
	if err == nil {
		jwt, err := CreatePatientJWT(&foundPatient)
		if err != nil {
			log.Println("[ERROR] JWT generation Error: ", err.Error())
			c.String(http.StatusInternalServerError, "A server error occurred")
			return
		} else {
			c.SetCookie("PatientAuth", jwt, 14*24*60*60, "/", "", false, false) // cookie expires after 14 days
			patientOutput := models.PatientLoginOutput{
				NHID:         foundPatient.NHID,
				Username:     foundPatient.Username,
				Name:         foundPatient.Name,
				Phone:        foundPatient.Phone,
				Gender:       foundPatient.Gender,
				AadharNumber: foundPatient.AadharNumber,
				DateOfBirth:  foundPatient.DateOfBirth,
			}
			c.JSON(http.StatusOK, &patientOutput)
		}
	} else {
		log.Println("[WARN] Invalid Password by patient: ", patientDetails.Username)
		c.String(http.StatusUnauthorized, "Invalid Password")
		return
	}
}

func RegisterNewPatient(c *gin.Context) {
	var patientDetails models.PatientRegistrationInput
	err := c.BindJSON(&patientDetails)
	if err != nil {
		log.Println("[WARN] Patient Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid patient details")
		return
	}

	if foundPatient, err := db.FindPatientByUsername(patientDetails.Username); err == nil {
		log.Println("[WARN] Username already in use: ", foundPatient)
		c.String(http.StatusBadRequest, "A Patient with this username has already registered")
		return
	}

	if foundPatient, err := db.FindPatientByAadhar(patientDetails.AadharNumber); err == nil {
		log.Println("[WARN] Aaadhar Number already in use: ", foundPatient)
		c.String(http.StatusBadRequest, "A Patient with this Aadhar Number has already registered")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(patientDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] Error while encrypting password: ", err.Error())
		c.String(http.StatusInternalServerError, "A server error occurred")
		return
	}

	newPatient := models.Patient{
		NHID:         uuid.New().String(),
		Name:         patientDetails.Name,
		Phone:        patientDetails.Phone,
		Username:     patientDetails.Username,
		AadharNumber: patientDetails.AadharNumber,
		Password:     string(hashedPassword),
		DateOfBirth:  patientDetails.DateOfBirth,
		Gender:       patientDetails.Gender,
	}

	err = db.AddNewPatient(&newPatient)
	if err != nil {
		log.Println("[ERROR] Error while adding new patient to db: ", err.Error())
		c.String(http.StatusInternalServerError, "A database error occurred")
		return
	}

	c.String(http.StatusOK, "Patient Successfully Registered")
}
