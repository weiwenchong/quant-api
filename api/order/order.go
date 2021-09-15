package order

import (
	"github.com/gin-gonic/gin"
	orderAdapter "github.com/weiwenchong/quant-order/adapter"
	order "github.com/weiwenchong/quant-order/pub"
	. "github.com/weiwenchong/quant-api/common"
	"net/http"
)

func ApiCreateGridOrder(c *gin.Context) {
	req := new(order.CreateGridOrderReq)
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

func ApiCloseOrder(c *gin.Context) {
	req := new(order.CloseOrderReq)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, ApiReturn{
			Ret : -1,
			Msg: "request paramenters format error",
		})
		return
	}

	res, err := orderAdapter.Client.CloseOrder(c.Request.Context(), req)
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

func ApiGetOrdersByUid(c *gin.Context) {
	req := new(order.GetOrdersByUidReq)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, ApiReturn{
			Ret : -1,
			Msg: "request paramenters format error",
		})
		return
	}

	res, err := orderAdapter.Client.GetOrdersByUid(c.Request.Context(), req)
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

func ApiGridTrial(c *gin.Context) {
	req := new(order.CreateGridOrderReq)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, ApiReturn{
			Ret : -1,
			Msg: "request paramenters format error",
		})
		return
	}

	res, err := orderAdapter.Client.GetGridTrial(c.Request.Context(), req)
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
