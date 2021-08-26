package logic

import order "github.com/wenchong-wei/quant-order/pub"

func InitLogic() {
	// init调用服务
	order.InitClient()
}
