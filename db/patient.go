package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dryairship/techmeet22-nhid-service/models"
)

func AddNewPatient(patient *models.Patient) error {
	_, err := NhidDb.PatientsCollection.InsertOne(context.Background(), patient)
	return err
}

func FindPatientByNHID(nhid string) (models.Patient, error) {
	var patient models.Patient
	err := NhidDb.PatientsCollection.FindOne(context.Background(), bson.M{"_id": nhid}).Decode(&patient)
	return patient, err
}

func FindPatientByUsername(username string) (models.Patient, error) {
	var patient models.Patient
	err := NhidDb.PatientsCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&patient)
	return patient, err
}

func FindPatientByAadhar(aadhar string) (models.Patient, error) {
	var patient models.Patient
	err := NhidDb.PatientsCollection.FindOne(context.Background(), bson.M{"aadhar": aadhar}).Decode(&patient)
	return patient, err
}
