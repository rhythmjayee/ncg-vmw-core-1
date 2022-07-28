package api

import (
	. "DIaaS/diaas/constants"
	. "DIaaS/diaas/db"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (registry *Registry) Deploy(lang string, source DeployBody) (*Deployment, error) {
	var db *mongo.Database
	db = DbConnection()
	coll := db.Collection("registry")
	deployment := NewDeployment(source, lang)
	if err := deployment.Build(); err != nil {
		return nil, fmt.Errorf("failed to build: %v", err)
	}
	filter := bson.D{{"id", deployment.ID}}
	var res *Deployment
	_ = coll.FindOne(context.TODO(), filter).Decode(&res)
	if !(res == nil) {
		return nil, fmt.Errorf("aldredy exists: %v", deployment.ID)
	}
	result, err := coll.InsertOne(context.TODO(), deployment)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return deployment, nil
}

func (registry *Registry) Run(id string, stdin io.Reader, stdout, stderr io.Writer, params string, duration time.Duration) error {
	deployment := registry.Resolve(id)
	if deployment == nil {
		return fmt.Errorf("deployment not found")
	}
	err := deployment.Run(stdin, stdout, stderr, params, duration)
	if err != nil {
		return err
	}
	return nil
}

func (registry *Registry) Resolve(identifier string) *Deployment {
	var db *mongo.Database
	db = DbConnection()
	if strings.HasPrefix(identifier, AliasPrefix) {
		coll := db.Collection("alias")
		filter := bson.D{{"name", strings.TrimPrefix(identifier, AliasPrefix)}}
		var result *Alias
		err := coll.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		identifier = result.For
	}
	coll := db.Collection("registry")
	var result *Deployment
	filter := bson.D{{"id", identifier}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (registry *Registry) Find(identifier string) *Deployment {
	var (
		item         *Deployment
		longestMatch int
	)
	for _, d := range registry.Items {
		if strings.HasPrefix(d.ID, identifier) && len(identifier) > longestMatch {
			item = d
			longestMatch = len(identifier)
		}
	}
	return item
}

func (registry *Registry) Bind(alias string, deployment *Deployment) *Alias {
	var db *mongo.Database
	db = DbConnection()
	item := &Alias{
		Name: alias,
		For:  deployment.ID,
		URL:  TriggerPrefix + AliasPrefix + alias,
		Date: time.Now(),
	}
	coll := db.Collection("alias")
	result, err := coll.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return item
}
