package api

import (
	. "DIaaS/diaas/auth"
	. "DIaaS/diaas/db"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (srv *Server) setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (srv *Server) error(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(struct {
		Error   bool
		Message string
	}{
		Error:   true,
		Message: msg,
	})
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		return
	}
	log.Printf("Error while serving request: %s\n", msg)
	w.Header().Add("Content-Type", "application/json")
}

func (srv *Server) deploy(w http.ResponseWriter, r *http.Request) {

	srv.setupCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		encoder := json.NewEncoder(w)

		decoder := json.NewDecoder(r.Body)
		var creds DeployBody
		err := decoder.Decode(&creds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		deployment, err := srv.Registry.Deploy(r.URL.Query().Get("lang"), creds)
		if err != nil {
			srv.error(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = encoder.Encode(deployment)
		if err != nil {
			log.Fatal("deployment encoding failed")
			return
		}
	}
	return
}

func (srv *Server) logs(w http.ResponseWriter, r *http.Request) {
	srv.setupCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}
	id := filepath.Base(r.URL.Path)
	deployment := srv.Registry.Resolve(id)
	if deployment == nil {
		deployment = srv.Registry.Find(id)
	}
	if deployment == nil {
		srv.error(w, http.StatusNotFound, "Deployment not found")
		return
	}
	logfile, err := os.Open(deployment.BuildLogs)
	if err != nil {
		srv.error(w, http.StatusNotFound, "BuildLogs not found")
		return
	}
	defer func(logfile *os.File) {
		err := logfile.Close()
		if err != nil {
			log.Fatal("logfile closure defer failed")
		}
	}(logfile)
	_, err = io.Copy(w, logfile)
	if err != nil {
		log.Fatal("io copy failed")
		return
	}
}

func (srv *Server) myfunctions(w http.ResponseWriter, r *http.Request) {

	srv.setupCORS(&w)
	var db *mongo.Database
	db = DbConnection()
	if r.Method == "OPTIONS" {
		return
	}
	if r.Method == "GET" {
		coll := db.Collection("registry")
		encoder := json.NewEncoder(w)
		results := []*Results{}
		user := r.URL.Query().Get("users")
		cursor, err := coll.Find(context.TODO(), bson.D{{"user", user}})
		if err != nil {
			log.Fatal(err)
		}

		if err = cursor.All(context.TODO(), &results); err != nil {
			log.Fatal(err)
		}
		err = encoder.Encode(results)
		if err != nil {
			log.Fatal("failed to encode results")
			return
		}
	}
	if r.Method == "DELETE" {
		coll := db.Collection("registry")
		encoder := json.NewEncoder(w)
		results := []*Results{}
		user := r.URL.Query().Get("users")
		id := filepath.Base(r.URL.Path)
		filter := bson.M{}
		if id == "" || user == "" {
			return
		}
		filter["id"] = id
		filter["user"] = user
		_, err := coll.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		err = encoder.Encode(results)
		if err != nil {
			log.Fatal("failed to encode results")
			return
		}
	}

}

func (srv *Server) describe(w http.ResponseWriter, r *http.Request) {
	srv.setupCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}
	id := filepath.Base(r.URL.Path)
	deployment := srv.Registry.Resolve(id)
	if deployment == nil {
		deployment = srv.Registry.Find(id)
	}
	if deployment == nil {
		srv.error(w, http.StatusNotFound, "Deployment not found")
		return
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(deployment)
	if err != nil {
		log.Fatal("failed to encode deployment")
		return
	}
}

func (srv *Server) bind(w http.ResponseWriter, r *http.Request) {
	err := Authenticate(w, r)
	if err != nil {
		log.Println("failed to authenticate. aborting bind")
		return
	}
	srv.setupCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}
	id := filepath.Base(r.URL.Path)
	alias := r.URL.Query().Get("to")
	deployment := srv.Registry.Resolve(id)
	if deployment == nil {
		srv.error(w, http.StatusNotFound, "Deployment does not exist")
		return
	}
	item := srv.Registry.Bind(alias, deployment)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(item)
	if err != nil {
		log.Fatal("failed to encode item")
		return
	}
}

func (srv *Server) trigger(w http.ResponseWriter, r *http.Request) {
	// err := Authenticate(w, r)
	// if err != nil {
	// 	srv.error(w, http.StatusUnauthorized, err.Error())
	// 	return
	// }
	srv.setupCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}
	id := filepath.Base(r.URL.Path)
	params := r.FormValue("params")
	decoder := json.NewDecoder(r.Body)
	var res TriggerRequestBody
	err := decoder.Decode(&res)
	duration, err := time.ParseDuration(res.Duration)
	if err != nil {
		log.Print("error in duration parsing", err)
		return
	}
	err = srv.Registry.Run(id, r.Body, w, w, params, duration)
	if err != nil {
		srv.error(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (srv *Server) listAlias(w http.ResponseWriter, r *http.Request) {
	srv.setupCORS(&w)
	var db *mongo.Database
	db = DbConnection()

	if r.Method == "OPTIONS" {
		return
	}
	encoder := json.NewEncoder(w)
	aliases := []*Alias{}
	coll := db.Collection("alias")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cur.All(context.TODO(), &aliases); err != nil {
		log.Fatal(err)
	}
	err = encoder.Encode(aliases)
	if err != nil {
		log.Fatal("failed to encode aliases")
		return
	}

}

func (srv *Server) list(w http.ResponseWriter, r *http.Request) {
	srv.setupCORS(&w)
	var db *mongo.Database
	db = DbConnection()
	if r.Method == "OPTIONS" {
		return
	}
	coll := db.Collection("registry")
	encoder := json.NewEncoder(w)
	results := []*Results{}
	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	err = encoder.Encode(results)
	if err != nil {
		log.Fatal("failed to encode results")
		return
	}
}
