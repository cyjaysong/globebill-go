package model

// TradeQueryReq 交易查询
type TradeQueryReq struct {
	TradeId     int64  `json:"tradeId,omitempty"`     // 交易流水号
	OutTransId  string `json:"outTransId,omitempty"`  // 第三方流水号,需唯一
	AppAccessId int64  `json:"appAccessId,omitempty"` // POS外部接入商编号,用于APP跳转POS收银台场景,outTransId为APP方流水号
}

// TradeQueryRes 交易查询
type TradeQueryRes struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;1:消费;2:撤销;3:退货
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS|FAILURE|UNKNOW
	TradeTime       string `json:"tradeTime"`                 // 交易时间;yyyy-MM-dd HH:mm:ss
	SourceTradeId   int64  `json:"sourceTradeId,omitempty"`   // 源交易流水号(撤销/退货交易才有)
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int64  `json:"tradeAmount"`               // 交易金额(单位:分)
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方流水号
	CashierId       int64  `json:"cashierId,omitempty"`       // 收银员ID号
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
	MerchantId      int64  `json:"merchantId"`                // 平台商户号
	TerminalId      int64  `json:"terminalId"`                // 平台终端号
	Fee             int64  `json:"fee,omitempty"`             // 交易手续费
	PayTime         string `json:"payTime,omitempty"`         // 完成支付时间;目前仅支持微信/支付宝;格式:yyyy-MM-dd HH:mm:ss
	ChannelMerNo    string `json:"channelMerNo,omitempty"`    // 渠道商户号
	CardType        string `json:"cardType,omitempty"`        // 借贷标识;00:借记卡;01:贷记卡
}
