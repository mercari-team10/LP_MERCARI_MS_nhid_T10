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

func LabLogin(c *gin.Context) {
	var labDetails models.LabLoginInput
	err := c.BindJSON(&labDetails)
	if err != nil {
		log.Println("[WARN] Lab Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid login details")
		return
	}

	foundLab, err := db.FindLabByUsername(labDetails.Username)
	if err != nil {
		log.Println("[WARN] Lab Username not found: ", labDetails.Username, err.Error())
		c.String(http.StatusNotFound, "Invalid Username")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundLab.Password), []byte(labDetails.Password))
	if err == nil {
		jwt, err := CreateLabJWT(&foundLab)
		if err != nil {
			log.Println("[ERROR] JWT generation Error: ", err.Error())
			c.String(http.StatusInternalServerError, "A server error occurred")
			return
		} else {
			c.SetCookie("LabAuth", jwt, 14*24*60*60, "/", "", false, false) // cookie expires after 14 days
			c.String(http.StatusOK, "Login Successful")
		}
	} else {
		log.Println("[WARN] Invalid Password by lab: ", labDetails.Username)
		c.String(http.StatusUnauthorized, "Invalid Password")
		return
	}
}

func RegisterNewLab(c *gin.Context) {
	var labDetails models.LabRegistrationInput
	err := c.BindJSON(&labDetails)
	if err != nil {
		log.Println("[WARN] Lab Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid lab details")
		return
	}

	if foundLab, err := db.FindLabByUsername(labDetails.Username); err == nil {
		log.Println("[WARN] Username already in use: ", foundLab)
		c.String(http.StatusBadRequest, "A Lab with this username has already registered")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(labDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] Error while encrypting password: ", err.Error())
		c.String(http.StatusInternalServerError, "A server error occurred")
		return
	}

	newLab := models.Lab{
		LabId:      uuid.New().String(),
		Name:       labDetails.Name,
		Phone:      labDetails.Phone,
		Username:   labDetails.Username,
		Address:    labDetails.Address,
		HospitalId: labDetails.HospitalId,
		Password:   string(hashedPassword),
	}

	err = db.AddNewLab(&newLab)
	if err != nil {
		log.Println("[ERROR] Error while adding new lab to db: ", err.Error())
		c.String(http.StatusInternalServerError, "A database error occurred")
		return
	}

	c.String(http.StatusOK, "Lab Successfully Registered")
}

func GetLabDetails(c *gin.Context) {
	labId := c.Param("id")
	foundLab, err := db.FindLabByLabId(labId)
	if err != nil {
		log.Println("[WARN] Error while fetching lab details: ", err.Error())
		c.String(http.StatusBadRequest, "Lab Not Found")
	} else {
		labOutput := models.LabDetailsOutput{
			LabId:      foundLab.LabId,
			Name:       foundLab.Name,
			Address:    foundLab.Address,
			Phone:      foundLab.Phone,
			HospitalId: foundLab.HospitalId,
		}
		c.JSON(http.StatusOK, &labOutput)
	}
}
