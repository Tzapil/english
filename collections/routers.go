package collections

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/tzapil/english/common"
)

// CollectionsRegister registrates common routers
func CollectionsRegister(router *gin.RouterGroup) {
	router.GET("/collections", CollectionsRetrieve)
	router.GET("/collections/random", RandomCollectionRetrieve)
	router.POST("/collection", CollectionCreate)
	router.GET("/collection/:id", CollectionRetrieve)
	router.PUT("/collection/:id", CollectionUpdate)
	router.DELETE("/collection/:id", CollectionDelete)
}

func CollectionDelete(c *gin.Context) {
	name, errID := primitive.ObjectIDFromHex(c.Param("id"))

	if errID != nil {
		log.Println(errID)
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	client := common.GetDB()
	collection := client.Database("english").Collection("collections")

	var result Collection

	filter := bson.D{{"_id", name}}
	err := collection.FindOneAndDelete(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"collection": result})
}

func RandomCollectionRetrieve(c *gin.Context) {
	client := common.GetDB()
	collection := client.Database("english").Collection("collections")
	count, err := collection.CountDocuments(context.TODO(), bson.D{{}})

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, common.NewError("collections", errors.New("Something goes wrong")))
		return
	}

	if count == 0 {
		c.JSON(http.StatusBadRequest, common.NewError("collections", errors.New("No collections")))
		return
	}

	var result Collection

	findOptions := options.FindOne()
	findOptions.SetSkip(rand.Int63n(count))

	errR := collection.FindOne(context.TODO(), bson.D{{}}, findOptions).Decode(&result)

	if errR != nil {
		log.Println(errR)
		c.JSON(http.StatusInternalServerError, common.NewError("collections", errors.New("Something goes wrong")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"collection": result})
}

func CollectionsRetrieve(c *gin.Context) {
	client := common.GetDB()
	collection := client.Database("english").Collection("collections")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}

	var results []*Collection

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Can't Find any results")))
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Collection
		err := cur.Decode(&elem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewError("collections", errors.New("Invalid DB result")))
			return
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, common.NewError("collections", errors.New("Cant proceed all values")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": results, "size": len(results)})
}

func CollectionRetrieve(c *gin.Context) {
	name, errID := primitive.ObjectIDFromHex(c.Param("id"))

	if errID != nil {
		log.Println(errID)
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	client := common.GetDB()
	collection := client.Database("english").Collection("collections")

	var result Collection

	filter := bson.D{{"_id", name}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"collection": result})
}

func CollectionUpdate(c *gin.Context) {
	id, errID := primitive.ObjectIDFromHex(c.Param("id"))
	if errID != nil {
		log.Println(errID)
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	var newCollection Collection
	errBind := c.BindJSON(&newCollection)

	if errBind != nil {
		log.Println(errBind)
		c.JSON(http.StatusBadRequest, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	if newCollection.Name == "" {
		log.Println("Empty Name")
		c.JSON(http.StatusBadRequest, common.NewError("collections", errors.New("Empty Name")))
		return
	}

	client := common.GetDB()
	collection := client.Database("english").Collection("collections")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": newCollection.Name}}

	// update document
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if res.ModifiedCount == 0 {
		c.JSON(http.StatusBadRequest, common.NewError("collections", errors.New("Something went wrong")))
	}

	c.JSON(http.StatusOK, gin.H{"modified": res.ModifiedCount})
}

func CollectionCreate(c *gin.Context) {
	client := common.GetDB()
	collection := client.Database("english").Collection("collections")

	var newCollection Collection
	c.BindJSON(&newCollection)
	newCollection.Date = time.Now()

	insertResult, err := collection.InsertOne(context.TODO(), newCollection)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"collection": insertResult})
}
