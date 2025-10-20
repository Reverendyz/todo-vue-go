package patch

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reverendyz/todo-vue-web/backend-service/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UpdateTask(c *gin.Context) {
	objID, err := bson.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(ctx)

	collection := client.Database("taskdb").Collection("tasks")

	filter := bson.M{"_id": objID}
	task := &types.Document{}
	collection.FindOne(ctx, filter).Decode(&task)
	task.Done = !task.Done
	update := bson.M{"$set": bson.M{"done": task.Done}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"matched":  result.MatchedCount,
		"modified": result.ModifiedCount,
		"upserted": result.UpsertedID,
	})
}
