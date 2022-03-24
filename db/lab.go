package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dryairship/techmeet22-nhid-service/models"
)

func AddNewLab(lab *models.Lab) error {
	_, err := NhidDb.LabsCollection.InsertOne(context.Background(), lab)
	return err
}

func FindLabByLabId(labId string) (models.Lab, error) {
	var lab models.Lab
	err := NhidDb.LabsCollection.FindOne(context.Background(), bson.M{"_id": labId}).Decode(&lab)
	return lab, err
}

func FindLabByUsername(username string) (models.Lab, error) {
	var lab models.Lab
	err := NhidDb.LabsCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&lab)
	return lab, err
}
