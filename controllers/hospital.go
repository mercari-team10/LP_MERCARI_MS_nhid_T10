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

func HospitalLogin(c *gin.Context) {
	var hospitalDetails models.HospitalLoginInput
	err := c.BindJSON(&hospitalDetails)
	if err != nil {
		log.Println("[WARN] Hospital Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid login details")
		return
	}

	foundHospital, err := db.FindHospitalByUsername(hospitalDetails.Username)
	if err != nil {
		log.Println("[WARN] Hospital Username not found: ", hospitalDetails.Username, err.Error())
		c.String(http.StatusNotFound, "Invalid Username")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundHospital.Password), []byte(hospitalDetails.Password))
	if err == nil {
		jwt, err := CreateHospitalJWT(&foundHospital)
		if err != nil {
			log.Println("[ERROR] JWT generation Error: ", err.Error())
			c.String(http.StatusInternalServerError, "A server error occurred")
			return
		} else {
			c.SetCookie("HospitalAuth", jwt, 14*24*60*60, "/", "", false, false) // cookie expires after 14 days
			c.String(http.StatusOK, "Login Successful")
		}
	} else {
		log.Println("[WARN] Invalid Password by hospital: ", hospitalDetails.Username)
		c.String(http.StatusUnauthorized, "Invalid Password")
		return
	}
}

func RegisterNewHospital(c *gin.Context) {
	var hospitalDetails models.HospitalRegistrationInput
	err := c.BindJSON(&hospitalDetails)
	if err != nil {
		log.Println("[WARN] Hospital Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid hospital details")
		return
	}

	if foundHospital, err := db.FindHospitalByUsername(hospitalDetails.Username); err == nil {
		log.Println("[WARN] Username already in use: ", foundHospital)
		c.String(http.StatusBadRequest, "A Hospital with this username has already registered")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(hospitalDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] Error while encrypting password: ", err.Error())
		c.String(http.StatusInternalServerError, "A server error occurred")
		return
	}

	newHospital := models.Hospital{
		HospitalId: uuid.New().String(),
		Name:       hospitalDetails.Name,
		Phone:      hospitalDetails.Phone,
		Username:   hospitalDetails.Username,
		Address:    hospitalDetails.Address,
		Password:   string(hashedPassword),
	}

	err = db.AddNewHospital(&newHospital)
	if err != nil {
		log.Println("[ERROR] Error while adding new hospital to db: ", err.Error())
		c.String(http.StatusInternalServerError, "A database error occurred")
		return
	}

	c.String(http.StatusOK, "Hospital Successfully Registered")
}

func GetHospitalDetails(c *gin.Context) {
	hospitalId := c.Param("id")
	foundHospital, err := db.FindHospitalByHospitalId(hospitalId)
	if err != nil {
		log.Println("[WARN] Error while fetching hospital details: ", err.Error())
		c.String(http.StatusBadRequest, "Hospital Not Found")
	} else {
		hospitalOutput := models.HospitalDetailsOutput{
			HospitalId: foundHospital.HospitalId,
			Name:       foundHospital.Name,
			Address:    foundHospital.Address,
			Phone:      foundHospital.Phone,
		}
		c.JSON(http.StatusOK, &hospitalOutput)
	}
}
