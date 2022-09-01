package Infra

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	r := gin.Default()
	db, err := ConnectDB("hihi")
	if err != nil {
		fmt.Println("Connect DB fail", err)

	}
	SetUpRouterToDo(r, db)
	return r
}
