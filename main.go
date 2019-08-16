package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/tzapil/english/collections"
	"github.com/tzapil/english/common"
	"github.com/tzapil/english/ping"
	"github.com/tzapil/english/words"
)

type Collection struct {
	_id  uint
	Name string
	Date time.Time
}

type Trainer struct {
	name string
	age  uint
	city string
}

func db() {
	client := common.GetDB()

	collection := client.Database("test").Collection("trainers")

	// create a value into which the result can be decoded
	var result Trainer

	filter := bson.D{{"name", "Ash"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	// collection := client.Database("english").Collection("collections")

	// ash := Collection{0, "Robot123", time.Now()}

	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// trainers := []interface{}{misty, brock}

	// insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// filter := bson.D{{"name", "Robot"}}

	// // create a value into which the result can be decoded
	// var result Collection

	// err = collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Found a single document: %+v\n", result)
}

func serve() {
	// creating of new router
	r := gin.Default()

	// make all handlers v1 api version
	v1 := r.Group("/api/v1")

	collections.CollectionRegister(v1)
	collections.CollectionsRegister(v1)

	words.WordsRegister(v1)
	words.WordRegister(v1)

	ping.PingRegister(v1)

	// run default listen and serve on 0.0.0.0:8080
	r.Run()
}

func main() {
	common.Init()
	defer common.Close()

	serve()

	// db()
}
