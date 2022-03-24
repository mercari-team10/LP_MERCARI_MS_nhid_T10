package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dryairship/techmeet22-nhid-service/models"
)

func AddNewHospital(hospital *models.Hospital) error {
	_, err := NhidDb.HospitalsCollection.InsertOne(context.Background(), hospital)
	return err
}

func FindHospitalByHospitalId(hospitalId string) (models.Hospital, error) {
	var hospital models.Hospital
	err := NhidDb.HospitalsCollection.FindOne(context.Background(), bson.M{"_id": hospitalId}).Decode(&hospital)
	return hospital, err
}

func FindHospitalByUsername(username string) (models.Hospital, error) {
	var hospital models.Hospital
	err := NhidDb.HospitalsCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&hospital)
	return hospital, err
}
