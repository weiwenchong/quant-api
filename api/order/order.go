package order

import (
	"github.com/gin-gonic/gin"
	orderAdapter "github.com/wenchong-wei/quant-order/pub"
	. "github.com/wenchong-wei/quant-api/common"
	"net/http"
)

func ApiCreateGridOrder(c *gin.Context) {
	req := new(orderAdapter.CreateGridOrderReq)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, ApiReturn{
			Ret : -1,
			Msg: "request paramenters format error",
		})
		return
	}

	res, err := orderAdapter.Client.CreateGridOrder(c.Request.Context(), req)
	if c.IsAborted() {
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, ApiReturn{
			Ret: -1,
			Msg: err.Error(),
		})
	}

	c.JSON(http.StatusOK, ApiReturn{
		Ret: 1,
		Data: &ApiData{
			Ext: res,
			Ent: nil,
		},
	})
}

