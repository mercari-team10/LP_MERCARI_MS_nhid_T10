package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	ApplicationPort   string
	MongoDialURL      string
	MongoDbName       string
	MongoUsername     string
	MongoUsingAuth    bool
	MongoPassword     string
	RSAPublicKeyFile  string
	RSAPrivateKeyFile string
)

func init() {
	viper.SetConfigName("nhid-config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("[WARN] Unable to locate configuration file: ", err.Error())
	}

	viper.AutomaticEnv()

	ApplicationPort = viper.GetString("ApplicationPort")

	MongoDialURL = viper.GetString("Mongo.URL")
	MongoDbName = viper.GetString("Mongo.Database")
	MongoUsername = viper.GetString("Mongo.Username")
	MongoPassword = viper.GetString("Mongo.Password")
	MongoUsingAuth = viper.GetBool("Mongo.UsingAuth")

	RSAPublicKeyFile = viper.GetString("RSA.PublicKeyFile")
	RSAPrivateKeyFile = viper.GetString("RSA.PrivateKeyFile")
}
