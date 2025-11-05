package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/reverendyz/todo-vue-web/backend-service/handlers/get"
	"github.com/reverendyz/todo-vue-web/backend-service/handlers/insert"
	"github.com/reverendyz/todo-vue-web/backend-service/handlers/patch"
	"github.com/reverendyz/todo-vue-web/backend-service/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	server := fmt.Sprintf("%s:%s",
		utils.GetEnvOrFallback("TODO_BACKEND_HOST", "0.0.0.0"),
		utils.GetEnvOrFallback("TODO_BACKEND_PORT", "8080"))

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"PATCH",
			"GET",
			"POST",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
		AllowOrigins: []string{
			"*",
		},
	}))
	r.GET("/live", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/ready", readyHandler)
	r.GET("/", get.FindAll)
	r.POST("/", insert.InsertTask)
	r.PATCH("/:id", patch.UpdateTask)

	if err := r.Run(server); err != nil {
		fmt.Printf("Error while spinning up the backend service. err: %v", err)
	}
}

func readyHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	options.Client().ApplyURI(utils.GetEnvOrFallback("TODO_BACKEND_DATABASE_URI", "mongodb://admin:password@localhost:27017"))

	client, err := mongo.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(ctx)

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		c.IndentedJSON(http.StatusRequestTimeout, fmt.Sprintf("Could not ping the server: %v", err))
		return
	}
}
