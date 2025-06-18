package model

// TradePushReq 交易推送
type TradePushReq struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;1:消费;2:撤销;3:退货;4:预授权请求;5:预授权完成;6:预授权撤销;8:预授权完成退款
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS:成功;FAILURE:失败;UNKNOW:未知
	TradeTime       string `json:"tradeTime"`                 // 交易时间;格式:yyyy-MM-dd HH:mm:ss
	SourceTradeId   int64  `json:"sourceTradeId,omitempty"`   // 源交易流水号;撤销交易和退货交易才有返回
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int    `json:"tradeAmount"`               // 交易金额(单位:分)
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方流水号,第三方系统内部流水号
	CashierId       int    `json:"cashierId,omitempty"`       // 收银员ID号
	SplitFlag       bool   `json:"splitFlag,omitempty"`       // 是否分账交易
	GoodsSubject    string `json:"goodsSubject,omitempty"`    // 商品名称
	TradeRemark     string `json:"tradeRemark,omitempty"`     // 交易备注
	UserRemark      string `json:"userRemark,omitempty"`      // 用户备注
	AcquirerTransNo string `json:"acquirerTransNO,omitempty"` // 收单交易号
	ChannelTransNo  string `json:"channelTransNo,omitempty"`  // 上游渠道交易号
	UserOpenId      string `json:"userOpenId,omitempty"`      // 微信open_id/支付宝user_id
	MerchantName    string `json:"merchantName"`              // 商户名称
	MrchntCode      string `json:"mrchntCode,omitempty"`      // 交易商户编号
	TerCode         string `json:"terCode,omitempty"`         // 交易终端编号
	MerchantId      int    `json:"merchantId"`                // 平台商户号
	TerminalI       int    `json:"terminalId"`                // 平台终端号
	Fee             int    `json:"fee,omitempty"`             // 交易手续费
	PayTime         string `json:"payTime,omitempty"`         // 完成支付时间;目前仅支持微信/支付宝
}

// TradeNotifyReq 交易通知
type TradeNotifyReq struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;1:消费;2:撤销;3:退货
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS:成功;FAILURE:失败;UNKNOW:未知
	TradeTime       string `json:"tradeTime"`                 // 交易时间;格式:yyyy-MM-dd HH:mm:ss
	SourceTradeId   int64  `json:"sourceTradeId,omitempty"`   // 源交易流水号;撤销交易和退货交易才有返回
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int    `json:"tradeAmount"`               // 交易金额(单位:分)
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方流水号
	CashierId       int    `json:"cashierId,omitempty"`       // 收银员ID号
	SplitFlag       bool   `json:"splitFlag,omitempty"`       // 是否分账交易
	GoodsSubject    string `json:"goodsSubject,omitempty"`    // 商品名称
	TradeRemark     string `json:"tradeRemark,omitempty"`     // 交易备注
	UserRemark      string `json:"userRemark,omitempty"`      // 用户备注
	AcquirerTransNo string `json:"acquirerTransNO,omitempty"` // 收单交易号
	ChannelTransNo  string `json:"channelTransNo,omitempty"`  // 上游渠道交易号
	UserOpenId      string `json:"userOpenId,omitempty"`      // 微信open_id/支付宝user_id
	MerchantName    string `json:"merchantName"`              // 商户名称
	MrchntCode      string `json:"mrchntCode,omitempty"`      // 交易商户编号
	TerCode         string `json:"terCode,omitempty"`         // 交易终端编号
	MerchantId      int    `json:"merchantId"`                // 平台商户号
	TerminalId      int    `json:"terminalId"`                // 平台终端号
	Fee             int    `json:"fee,omitempty"`             // 交易手续费
	PayTime         string `json:"payTime,omitempty"`         // 完成支付时间;格式:yyyy-MM-dd HH:mm:ss
	ChannelMerNo    string `json:"channelMerNo,omitempty"`    // 渠道商户号
	CardType        string `json:"cardType,omitempty"`        // 借贷标识;00:借记卡;01:贷记卡
}

// TradeSplitNotifyReq 交易分账通知
type TradeSplitNotifyReq struct {
	TradeId          int64  `json:"tradeId"`               // 交易流水号
	SplitOrderStatus int    `json:"splitOrderStatus"`      // 分账状态;2:成功;3:失败
	SplitAmount      int    `json:"splitAmount"`           // 分账金额,单位:分
	SplitFee         int    `json:"splitFee,omitempty"`    // 分账手续费,单位:分
	OutTransId       string `json:"outTransId"`            // 第三方系统内部流水号
	SplitResult      string `json:"splitResult,omitempty"` // 分账结果描述
	MerchantId       int    `json:"merchantId"`            // 平台商户号
	TerminalId       int    `json:"terminalId"`            // 平台终端号
}
