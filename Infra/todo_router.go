package Infra

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todoApi/handler"
)

func SetUpRouterToDo(r *gin.Engine, db *mongo.Client) {
	handlerTodo := handler.NewToDoHandler(db)
	router := r.Group("/todo")
	{
		router.GET("/", handlerTodo.GetToDo)
		router.GET("/all", handlerTodo.GetAllToDo)
		router.POST("/add", handlerTodo.AddNewToDo)
		router.DELETE("/delete", handlerTodo.DeleteToDo)
		router.PUT("/update", handlerTodo.UpdateToDo)
	}

}
