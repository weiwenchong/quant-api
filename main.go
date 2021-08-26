package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchong-wei/quant-api/logic"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	logic.RegisterHttp(router)
	logic.InitLogic()

	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Panicf("quantapi router.Run err:%v", err)
	}
}
