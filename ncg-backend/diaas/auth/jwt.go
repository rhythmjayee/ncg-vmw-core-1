package auth

import (
	. "DIaaS/diaas/db"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Todo - Deprecate this
var jwtKey = []byte("my_secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	var db *mongo.Database
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db = DbConnection()
	coll := db.Collection("user_data")
	filter := bson.D{{"username", creds.Username}}
	var expected UserRegistration
	err = coll.FindOne(context.TODO(), filter).Decode(&expected)
	if err != nil || expected.Password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Authenticate(w http.ResponseWriter, r *http.Request) error {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	var db *mongo.Database
	var userInfo UserRegistration
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db = DbConnection()
	coll := db.Collection("user_data")
	_, err = coll.InsertOne(context.TODO(), userInfo)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Println("failed to insert user credentials")
		return
	}
}
