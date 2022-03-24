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

func DoctorLogin(c *gin.Context) {
	var doctorDetails models.DoctorLoginInput
	err := c.BindJSON(&doctorDetails)
	if err != nil {
		log.Println("[WARN] Doctor Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid login details")
		return
	}

	foundDoctor, err := db.FindDoctorByUsername(doctorDetails.Username)
	if err != nil {
		log.Println("[WARN] Doctor Username not found: ", doctorDetails.Username, err.Error())
		c.String(http.StatusNotFound, "Invalid Username")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundDoctor.Password), []byte(doctorDetails.Password))
	if err == nil {
		jwt, err := CreateDoctorJWT(&foundDoctor)
		if err != nil {
			log.Println("[ERROR] JWT generation Error: ", err.Error())
			c.String(http.StatusInternalServerError, "A server error occurred")
			return
		} else {
			c.SetCookie("DoctorAuth", jwt, 14*24*60*60, "/", "", false, false) // cookie expires after 14 days
			doctorOutput := models.DoctorLoginOutput{
				DoctorId: foundDoctor.DoctorId,
				Name:     foundDoctor.Name,
				Phone:    foundDoctor.Phone,
				UPRN:     foundDoctor.UPRN,
				Username: foundDoctor.Username,
			}
			c.JSON(http.StatusOK, &doctorOutput)
		}
	} else {
		log.Println("[WARN] Invalid Password by doctor: ", doctorDetails.Username)
		c.String(http.StatusUnauthorized, "Invalid Password")
		return
	}
}

func RegisterNewDoctor(c *gin.Context) {
	var doctorDetails models.DoctorRegistrationInput
	err := c.BindJSON(&doctorDetails)
	if err != nil {
		log.Println("[WARN] Doctor Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid doctor details")
		return
	}

	if foundDoctor, err := db.FindDoctorByUsername(doctorDetails.Username); err == nil {
		log.Println("[WARN] Username already in use: ", foundDoctor)
		c.String(http.StatusBadRequest, "A Doctor with this username has already registered")
		return
	}

	if foundDoctor, err := db.FindDoctorByUPRN(doctorDetails.UPRN); err == nil {
		log.Println("[WARN] UPRN already in use: ", foundDoctor)
		c.String(http.StatusBadRequest, "A Doctor with this UPRN has already registered")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(doctorDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] Error while encrypting password: ", err.Error())
		c.String(http.StatusInternalServerError, "A server error occurred")
		return
	}

	newDoctor := models.Doctor{
		DoctorId: uuid.New().String(),
		Name:     doctorDetails.Name,
		Phone:    doctorDetails.Phone,
		Username: doctorDetails.Username,
		UPRN:     doctorDetails.UPRN,
		Password: string(hashedPassword),
	}

	err = db.AddNewDoctor(&newDoctor)
	if err != nil {
		log.Println("[ERROR] Error while adding new doctor to db: ", err.Error())
		c.String(http.StatusInternalServerError, "A database error occurred")
		return
	}

	c.String(http.StatusOK, "Doctor Successfully Registered")
}

func GetDoctorDetails(c *gin.Context) {
	doctorId := c.Param("id")
	foundDoctor, err := db.FindDoctorByDoctorId(doctorId)
	if err != nil {
		log.Println("[WARN] Error while fetching doctor details: ", err.Error())
		c.String(http.StatusBadRequest, "Doctor Not Found")
	} else {
		doctorOutput := models.DoctorDetailsOutput{
			DoctorId: foundDoctor.DoctorId,
			Name:     foundDoctor.Name,
			Phone:    foundDoctor.Phone,
		}
		c.JSON(http.StatusOK, &doctorOutput)
	}
}
