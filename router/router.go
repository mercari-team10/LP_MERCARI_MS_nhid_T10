package router

import (
	"github.com/gin-gonic/gin"

	"github.com/dryairship/techmeet22-nhid-service/controllers"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.RespondToPing)
	r.GET("/publicKey", controllers.GetPublicKey)

	r.POST("/patient/register", controllers.RegisterNewPatient)
	r.POST("/patient/login", controllers.PatientLogin)

	r.POST("/doctor/register", controllers.RegisterNewDoctor)
	r.POST("/doctor/login", controllers.DoctorLogin)
	r.GET("/doctor/details/:id", controllers.GetDoctorDetails)

	r.POST("/lab/register", controllers.RegisterNewLab)
	r.POST("/lab/login", controllers.LabLogin)
	r.GET("/lab/details/:id", controllers.GetLabDetails)

	r.POST("/hospital/register", controllers.RegisterNewHospital)
	r.POST("/hospital/login", controllers.HospitalLogin)
	r.GET("/hospital/details/:id", controllers.GetHospitalDetails)

	r.POST("/records", controllers.EnsureDoctorOrLab(), controllers.AddNewRecord)
	r.GET("/records", controllers.EnsureAnyAuth(), controllers.FindRecords)
}
