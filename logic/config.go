package logic

import orderAdapter "github.com/wenchong-wei/quant-order/adapter"

func InitLogic() {
	// init调用服务
	orderAdapter.InitClient()
}
