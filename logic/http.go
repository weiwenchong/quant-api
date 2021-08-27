package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchong-wei/quant-api/api/order"
	. "github.com/wenchong-wei/quant-api/common"
	"net/http"
)

func RegisterHttp(router *gin.Engine) {
	router.POST("/quant/test/", ApiTest)
	router.POST("/quant/order/grid/create/", order.ApiCreateGridOrder)
	router.POST("/quant/order/close/", order.ApiCloseOrder)
	router.POST("/quant/order/list/", order.ApiGetOrdersByUid)
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