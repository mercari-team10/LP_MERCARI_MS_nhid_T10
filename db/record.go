package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dryairship/techmeet22-nhid-service/models"
)

var findRecordsOptions *options.FindOptions

func init() {
	findRecordsOptions = options.Find()
	findRecordsOptions.SetSort(bson.D{{"timestamp", -1}})
	findRecordsOptions.SetLimit(20)
}

func AddNewRecord(record *models.Record) error {
	_, err := NhidDb.RecordsCollection.InsertOne(context.Background(), record)
	return err
}

func FindRecords(nhid string, startDate, endDate int64) ([]models.Record, error) {
	var records []models.Record
	cursor, err := NhidDb.PatientsCollection.Find(context.Background(),
		bson.M{
			"NHID": nhid,
			"timestamp": bson.M{
				"$le": endDate,
				"$ge": startDate,
			},
		},
		findRecordsOptions,
	)
	if err != nil {
		return records, err
	}

	err = cursor.All(context.Background(), &records)
	return records, err
}
