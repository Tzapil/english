package words

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
func WordsRegister(router *gin.RouterGroup) {
	router.POST("/word", WordCreate)
	router.GET("/word/:id", WordRetrieve)
	router.GET("/words/random", RandomWordRetrieve)
	router.PUT("/word/:id", WordUpdate)
	router.GET("/collection/:id/words", WordsRetrieve)
	router.GET("/collection/:id/words/random", RandomCollectionWordRetrieve)
	router.DELETE("/word/:id", CollectionDelete)
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

	var result Word

	filter := bson.D{{"_id", name}}
	err := collection.FindOneAndDelete(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"collection": result})
}

func RandomWordRetrieve(c *gin.Context) {
	client := common.GetDB()
	collection := client.Database("english").Collection("words")

	count, err := collection.CountDocuments(context.TODO(), bson.D{{}})

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	if count == 0 {
		c.JSON(http.StatusBadRequest, common.NewError("words", errors.New("Empty collection")))
		return
	}

	var result Word

	findOptions := options.FindOne()
	findOptions.SetSkip(rand.Int63n(count))

	errR := collection.FindOne(context.TODO(), bson.D{{}}, findOptions).Decode(&result)

	if errR != nil {
		log.Println(errR)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"word": result})
}

func RandomCollectionWordRetrieve(c *gin.Context) {
	name, errID := primitive.ObjectIDFromHex(c.Param("id"))

	if errID != nil {
		log.Println(errID)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	client := common.GetDB()
	collection := client.Database("english").Collection("words")

	count, err := collection.CountDocuments(context.TODO(), bson.D{{"collection_id", name}})

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	if count == 0 {
		c.JSON(http.StatusBadRequest, common.NewError("words", errors.New("Empty collection")))
		return
	}

	var result Word

	findOptions := options.FindOne()
	findOptions.SetSkip(rand.Int63n(count))

	errR := collection.FindOne(context.TODO(), bson.D{{"collection_id", name}}, findOptions).Decode(&result)

	if errR != nil {
		log.Println(errR)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"word": result})
}

func WordsRetrieve(c *gin.Context) {
	name, errID := primitive.ObjectIDFromHex(c.Param("id"))

	if errID != nil {
		log.Println(errID)
		c.JSON(http.StatusBadRequest, common.NewError("collections", errors.New("Invalid param")))
		return
	}

	client := common.GetDB()
	collection := client.Database("english").Collection("words")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Can't Find any results")))
		return
	}

	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}

	results := make([]*Word, 0)

	cur, err := collection.Find(context.TODO(), bson.D{{"collection_id", name}}, findOptions)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Can't Find any results")))
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Word
		err := cur.Decode(&elem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewError("words", errors.New("Invalid DB result")))
			return
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, common.NewError("words", errors.New("Can't proceed all values")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": results, "size": len(results)})
}

func WordRetrieve(c *gin.Context) {
	name, errID := primitive.ObjectIDFromHex(c.Param("id"))

	if errID != nil {
		log.Println(errID)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	client := common.GetDB()
	collection := client.Database("english").Collection("words")

	var result Word

	filter := bson.D{{"_id", name}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"word": result})
}

func WordUpdate(c *gin.Context) {
	id, errID := primitive.ObjectIDFromHex(c.Param("id"))
	if errID != nil {
		log.Println(errID)
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	var newCollection Word
	errBind := c.BindJSON(&newCollection)

	if errBind != nil {
		log.Println(errBind)
		c.JSON(http.StatusBadRequest, common.NewError("words", errors.New("Invalid param")))
		return
	}

	if newCollection.Word == "" {
		log.Println("Empty Name")
		c.JSON(http.StatusBadRequest, common.NewError("words", errors.New("Empty Name")))
		return
	}

	client := common.GetDB()
	collection := client.Database("english").Collection("words")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"word": newCollection.Word}}

	// update document
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if res.ModifiedCount == 0 {
		c.JSON(http.StatusBadRequest, common.NewError("words", errors.New("Something went wrong")))
	}

	c.JSON(http.StatusOK, gin.H{"modified": res.ModifiedCount})
}

func WordCreate(c *gin.Context) {
	client := common.GetDB()
	collection := client.Database("english").Collection("words")

	var newWord Word
	c.BindJSON(&newWord)
	newWord.Date = time.Now()

	// 422 validation

	insertResult, err := collection.InsertOne(context.TODO(), newWord)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("words", errors.New("Invalid param")))
		return
	}

	// insertResult -> {InsertedID: "abbababababbabab"}
	c.JSON(http.StatusCreated, gin.H{"collection": insertResult})
}
