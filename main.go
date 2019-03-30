package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoSession *mgo.Session

// Users type
type Users struct {
	ID   bson.ObjectId `json:"_id"`
	Name string        `json:"name"`
}

const database = "graphql"

func main() {
	err := Init("mongodb://root:secret@localhost:27017/graphql?authSource=admin")

	if err != nil {
		log.Fatalf("Not connect MongoDB %v", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/api/users", usersHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := mongoSession.Copy()
	defer s.Close()
	var users []*Users
	err := s.DB(database).C("users").Find(nil).All(&users)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}

// Init mongodb
func Init(mongoURL string) error {
	var err error
	mongoSession, err = mgo.Dial(mongoURL)

	return err
}
