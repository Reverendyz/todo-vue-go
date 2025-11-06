package insert

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reverendyz/todo-vue-web/backend-service/types"
	"github.com/reverendyz/todo-vue-web/backend-service/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InsertTask(c *gin.Context) {
	var payload types.Document
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	client, _ := mongo.Connect(options.Client().ApplyURI(utils.GetEnvOrFallback("TODO_BACKEND_DATABASE_URI", "mongodb://admin:password@localhost:27017")))
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("taskdb").Collection("tasks")

	parsed, err := bson.Marshal(payload)
	if err != nil {
		fmt.Printf("Could not parse the object to bson. err: %v", err)
	}

	result, err := collection.InsertOne(ctx, parsed)
	if err != nil {
		fmt.Printf("Could not insert. err: %v", err)
	}

	c.IndentedJSON(http.StatusCreated,
		gin.H{
			"message": result.InsertedID,
		})
}
