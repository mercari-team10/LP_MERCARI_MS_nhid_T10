package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/dryairship/techmeet22-nhid-service/config"
)

type NHIDDatabase struct {
	Session             *mongo.Client
	PatientsCollection  *mongo.Collection
	RecordsCollection   *mongo.Collection
	DoctorsCollection   *mongo.Collection
	LabsCollection      *mongo.Collection
	HospitalsCollection *mongo.Collection
}

var NhidDb NHIDDatabase

func init() {
	var connectURL string

	if config.MongoUsingAuth {
		connectURL = fmt.Sprintf(
			"mongodb://%s:%s@%s/%s",
			url.QueryEscape(config.MongoUsername),
			url.QueryEscape(config.MongoPassword),
			config.MongoDialURL,
			config.MongoDbName,
		)
	} else {
		connectURL = fmt.Sprintf(
			"mongodb://%s/%s",
			config.MongoDialURL,
			config.MongoDbName,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURL))
	if err != nil {
		log.Fatalf("[ERROR] Cannot connect to Mongo. Error: %v\n", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("[ERROR] Cannot Ping Mongo. Error: %v\n", err)
	} else {
		log.Println("[INFO] Successfully pinged Mongo")
		NhidDb.Session = mongoClient
		NhidDb.PatientsCollection = mongoClient.Database(config.MongoDbName).Collection("patients")
		NhidDb.RecordsCollection = mongoClient.Database(config.MongoDbName).Collection("records")
		NhidDb.DoctorsCollection = mongoClient.Database(config.MongoDbName).Collection("doctors")
		NhidDb.LabsCollection = mongoClient.Database(config.MongoDbName).Collection("labs")
		NhidDb.HospitalsCollection = mongoClient.Database(config.MongoDbName).Collection("hospitals")
	}
}
