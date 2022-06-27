package routes

import (
	"github.com/gin-gonic/gin"
	"context"
	"net/http"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var entryCollection *mongo.Collection = openCollection(Client, "calories")
// c *gin.Context has all request params
func AddEntry(c *gin.Context) {
}

func GetEntries(c *gin.Context) {
	var ctx, cancel = Context.WithTimeout(context.Background(), 100*time.second)

	var entries []bson.M 
	cursor, err := entryCollection.Find(ctx, bson.M{})
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	if err= cursor.All(ctx, &entries); err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOk, entries)
}

func GetEntryById(c *gin.Context) {
}

func GetEntryByIngredient(c *gin.Context) {
}

func UpdateEntry(c *gin.Context) {
}

func UpdateIngredient(c *gin.Context) {
}

func DeleteEntry(c *gin.Context) {
	entryId := c.Params.ByName("id")
	docId, _ := primitive.ObjectIDFromHex(entryId)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	entryCollection.DeleteOne(ctx, bson.M{"_id":docId})

	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOk, result.DeletedCount)
}
