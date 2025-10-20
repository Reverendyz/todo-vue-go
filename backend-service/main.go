package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/reverendyz/todo-vue-web/backend-service/handlers/get"
	"github.com/reverendyz/todo-vue-web/backend-service/handlers/insert"
	"github.com/reverendyz/todo-vue-web/backend-service/handlers/patch"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/", get.FindAll)
	r.POST("/", insert.InsertTask)
	r.PATCH("/:id", patch.UpdateTask)

	if err := r.Run(); err != nil {
		fmt.Printf("Error while spinning up the backend service. err: %v", err)
	}
}
