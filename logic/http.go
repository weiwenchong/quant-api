package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchong-wei/quant-api/api/order"
	. "github.com/wenchong-wei/quant-api/common"
	"net/http"
)

func RegisterHttp(router *gin.Engine) {
	router.POST("/quant/order/grid/create/", order.ApiCreateGridOrder)
	router.POST("/quant/test/", ApiTest)
}


func ApiTest(c *gin.Context) {
	req := new(string)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, ApiReturn{
			Ret : -1,
			Msg: "request paramenters format error",
		})
		return
	}

	*req += "is ok"

	c.JSON(http.StatusOK, ApiReturn{
		Ret: 1,
		Data: &ApiData{
			Ext: req,
			Ent: nil,
		},
	})
}