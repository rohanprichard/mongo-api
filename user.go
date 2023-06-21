package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var user User

	json.NewDecoder(r.Body).Decode(&user)
	
	coll := client.Database("db").Collection("cl")
	ctx, _:= context.WithTimeout(context.Background(), 10*time.Second)

	result,_ := coll.InsertOne(ctx, user)

	fmt.Print("\n\nUser created \n\n")
	json.NewEncoder(w).Encode(result)
}

func ShowUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var user []User

	coll := client.Database("db").Collection("cl")
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		panic(err)
	}

	result, _ := coll.Find(ctx, bson.M{}) 

	for result.Next(ctx){
		var u User
		result.Decode(&u)
		user = append(user,u)
	}

	json.NewEncoder(w).Encode(user)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	coll := client.Database("db").Collection("cl")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	result:= coll.FindOne(ctx, User{ID:id}) 

	var u User
	result.Decode(&u)
	json.NewEncoder(w).Encode(u)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	coll := client.Database("db").Collection("cl")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	result:= coll.FindOneAndDelete(ctx, User{ID:id}) 

	var u User
	result.Decode(&u)


	fmt.Print("\n\nUser deleted\n\n")
	json.NewEncoder(w).Encode(u)
}