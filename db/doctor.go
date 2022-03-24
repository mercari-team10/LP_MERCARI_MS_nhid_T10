package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dryairship/techmeet22-nhid-service/models"
)

func AddNewDoctor(doctor *models.Doctor) error {
	_, err := NhidDb.DoctorsCollection.InsertOne(context.Background(), doctor)
	return err
}

func FindDoctorByDoctorId(doctorId string) (models.Doctor, error) {
	var doctor models.Doctor
	err := NhidDb.DoctorsCollection.FindOne(context.Background(), bson.M{"_id": doctorId}).Decode(&doctor)
	return doctor, err
}

func FindDoctorByUsername(username string) (models.Doctor, error) {
	doctor := models.Doctor{}
	fmt.Println(NhidDb)
	err := NhidDb.DoctorsCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&doctor)
	return doctor, err
}

func FindDoctorByUPRN(uprn string) (models.Doctor, error) {
	var doctor models.Doctor
	err := NhidDb.DoctorsCollection.FindOne(context.Background(), bson.M{"uprn": uprn}).Decode(&doctor)
	return doctor, err
}
