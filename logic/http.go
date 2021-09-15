package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/weiwenchong/quant-api/api/order"
	. "github.com/weiwenchong/quant-api/common"
	"net/http"
)

func RegisterHttp(router *gin.Engine) {
	router.POST("/quant/test/", ApiTest)
	router.POST("/quant/order/grid/create/", order.ApiCreateGridOrder)
	router.POST("/quant/order/close/", order.ApiCloseOrder)
	router.POST("/quant/order/list/", order.ApiGetOrdersByUid)
	router.POST("/quant/order/grid/trial", order.ApiGridTrial)
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