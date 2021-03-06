package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weiwenchong/quant-api/logic"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	logic.RegisterHttp(router)
	logic.InitLogic()

	err := router.Run("0.0.0.0:10001")
	if err != nil {
		log.Panicf("quantapi router.Run err:%v", err)
	}
}
