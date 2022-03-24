package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/dryairship/techmeet22-nhid-service/models"
)

func parseJWTandGetClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Println("[WARN] Unexpected signing method: ", token.Header["alg"])
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return RSAPublicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Invalid token claims")
	}
}

func CreatePatientJWT(patient *models.Patient) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"NHID": patient.NHID,
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().AddDate(0, 0, 14).Unix(), // Exppire after 14 days of issue
	})
	return token.SignedString(RSAPrivateKey)
}

func CreateDoctorJWT(doctor *models.Doctor) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"DoctorId": doctor.DoctorId,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().AddDate(0, 0, 14).Unix(), // Exppire after 14 days of issue
	})
	return token.SignedString(RSAPrivateKey)
}

func CreateLabJWT(lab *models.Lab) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"LabId": lab.LabId,
		"nbf":   time.Now().Unix(),
		"exp":   time.Now().AddDate(0, 0, 14).Unix(), // Exppire after 14 days of issue
	})
	return token.SignedString(RSAPrivateKey)
}

func CreateHospitalJWT(hospital *models.Hospital) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"HospitalId": hospital.HospitalId,
		"nbf":        time.Now().Unix(),
		"exp":        time.Now().AddDate(0, 0, 14).Unix(), // Exppire after 14 days of issue
	})
	return token.SignedString(RSAPrivateKey)
}

func GetPatientFromJWT(tokenString string) (string, error) {
	claims, err := parseJWTandGetClaims(tokenString)
	if err != nil {
		return "", err
	}
	return claims["NHID"].(string), nil
}

func GetDoctorFromJWT(tokenString string) (string, error) {
	claims, err := parseJWTandGetClaims(tokenString)
	if err != nil {
		return "", err
	}
	return claims["DoctorId"].(string), nil
}

func GetLabFromJWT(tokenString string) (string, error) {
	claims, err := parseJWTandGetClaims(tokenString)
	if err != nil {
		return "", err
	}
	return claims["LabId"].(string), nil
}

func GetHospitalFromJWT(tokenString string) (string, error) {
	claims, err := parseJWTandGetClaims(tokenString)
	if err != nil {
		return "", err
	}
	return claims["HospitalId"].(string), nil
}
