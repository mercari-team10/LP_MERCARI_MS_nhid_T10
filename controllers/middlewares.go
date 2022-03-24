package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func EnsureDoctorOrLab() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		if doctorJWT, err := c.Cookie("DoctorAuth"); err == nil {
			doctorId, err := GetDoctorFromJWT(doctorJWT)
			if err == nil {
				c.Set("DoctorID", doctorId)
				c.Next()
			}
		} else if labJWT, err := c.Cookie("LabAuth"); err == nil {
			labId, err := GetLabFromJWT(labJWT)
			if err == nil {
				c.Set("LabID", labId)
				c.Next()
			}
		}
		log.Println("[WARN] Unauthorized access attempt: ", err.Error())
		c.Abort()
	}
}

func EnsureAnyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		if patientJWT, err := c.Cookie("PatientAuth"); err == nil {
			nhid, err := GetPatientFromJWT(patientJWT)
			if err == nil {
				c.Set("PatientID", nhid)
				c.Next()
			}
		} else if doctorJWT, err := c.Cookie("DoctorAuth"); err == nil {
			doctorId, err := GetDoctorFromJWT(doctorJWT)
			if err == nil {
				c.Set("DoctorID", doctorId)
				c.Next()
			}
		} else if labJWT, err := c.Cookie("LabAuth"); err == nil {
			labId, err := GetLabFromJWT(labJWT)
			if err == nil {
				c.Set("LabID", labId)
				c.Next()
			}
		}
		log.Println("[WARN] Unauthorized access attempt: ", err.Error())
		c.Abort()
	}
}
