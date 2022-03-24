package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dryairship/techmeet22-nhid-service/db"
	"github.com/dryairship/techmeet22-nhid-service/models"
)

func AddNewRecord(c *gin.Context) {
	var newRecordDetails models.AddRecordInput
	err := c.BindJSON(&newRecordDetails)
	if err != nil {
		log.Println("[WARN] Record Details are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid record details")
		return
	}

	newRecord := models.Record{
		RecordId:    uuid.New().String(),
		NHID:        newRecordDetails.NHID,
		RecordType:  newRecordDetails.RecordType,
		Comments:    newRecordDetails.Comments,
		Attachments: newRecordDetails.Attachments,
		Timestamp:   time.Now().Unix(),
	}

	if doctorId, isDoctor := c.Get("DoctorID"); isDoctor {
		newRecord.IssuerType = "doctor"
		newRecord.IssuerId = doctorId.(string)
		err = db.AddNewRecord(&newRecord)
		if err == nil {
			c.String(http.StatusOK, "Record added")
		}
	} else if labId, isLab := c.Get("LabID"); isLab {
		newRecord.IssuerType = "lab"
		newRecord.IssuerId = labId.(string)
		err = db.AddNewRecord(&newRecord)
		if err == nil {
			c.String(http.StatusOK, "Record added")
		}
	}
	log.Println("[WARN] Error while addding record: ", err.Error())
	c.String(http.StatusInternalServerError, "A server error occurred")

}

func FindRecords(c *gin.Context) {
	var findRecordsFilters models.FindRecordsInput
	err := c.BindJSON(&findRecordsFilters)

	isAuthorizedRequest := false
	if patientId, isPatient := c.Get("PatientId"); isPatient {
		isAuthorizedRequest = (findRecordsFilters.NHID == patientId)
	} else if doctorId, isDoctor := c.Get("DoctorId"); isDoctor {
		isAuthorizedRequest = verifyDoctorIsAuthorized(doctorId.(string), findRecordsFilters.NHID)
	} else if labId, isLab := c.Get("LabId"); isLab {
		isAuthorizedRequest = verifyLabIsAuthorized(labId.(string), findRecordsFilters.NHID)
	}

	if !isAuthorizedRequest {
		c.String(http.StatusUnauthorized, "This user is not authorized to view these records")
		return
	}

	if err != nil {
		log.Println("[WARN] Record Filters are invalid: ", err.Error())
		c.String(http.StatusBadRequest, "Invalid record filers")
		return
	}

	if findRecordsFilters.EndDate == 0 {
		findRecordsFilters.EndDate = time.Now().Unix()
	}

	records, err := db.FindRecords(findRecordsFilters.NHID, findRecordsFilters.StartDate, findRecordsFilters.EndDate)
	if err != nil {
		c.String(http.StatusInternalServerError, "A server error occurred")
	} else {
		c.JSON(http.StatusOK, &records)
	}
}

func verifyDoctorIsAuthorized(doctorId, patientId string) bool {
	return true
}

func verifyLabIsAuthorized(labId, patientId string) bool {
	return true
}
