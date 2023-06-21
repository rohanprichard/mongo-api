package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func HandleReqs() {
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user", ShowUsers).Methods("GET")
	router.HandleFunc("/user/{id}", ShowUser).Methods("GET")
	router.HandleFunc("/user/{id}/delete", DeleteUser).Methods("GET")
	http.ListenAndServe(":9000", router)
}

func main() {
	fmt.Println("Start Here")
	ctx, _:= context.WithTimeout(context.Background(), 10*time.Second)

	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://rohanrichard:hahahaha@cluster0.ebrz5ph.mongodb.net/?retryWrites=true&w=majority"))
	HandleReqs()
}
