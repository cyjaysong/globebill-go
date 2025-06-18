package model

// TradeRefundReq 交易退款
type TradeRefundReq struct {
	SN                   string         `json:"sn,omitempty"`                   // 设备TUSN号;有固定TUSN号可使用sn参数交易
	MerchantId           int            `json:"merchantId,omitempty"`           // 平台商户号;渠道对接可使用merchantId和terminalId参数交易
	TerminalId           int            `json:"terminalId,omitempty"`           // 平台终端号;渠道对接可使用merchantId和terminalId参数交易
	SourceTradeId        int64          `json:"sourceTradeId,omitempty"`        // 原交易流水号(选其一填)
	SourceOutTransId     string         `json:"sourceOutTransId,omitempty"`     // 第三方原交易流水号(选其一填)
	SourceChannelTransNo string         `json:"sourceChannelTransNO,omitempty"` // 渠道原交易流水号(选其一填)
	TradeAmount          int            `json:"tradeAmount"`                    // 交易金额(单位:分)
	OutTransId           string         `json:"outTransId"`                     // 第三方退款流水号;第三方系统内部流水号，需唯一
	AppAccessId          int            `json:"appAccessId,omitempty"`          // POS外部接入商编号;用于APP跳转POS收银台场景，outTransId为APP方流水号
	TradeRemark          string         `json:"tradeRemark,omitempty"`          // 交易备注(退款)
	CashierId            int            `json:"cashierId,omitempty"`            // 收银员ID号
	SplitReceiveList     []SplitReceive `json:"splitReceiveList,omitempty"`     // 退款时分账接收方承担列表
}

// TradeRefundRes 交易退款
type TradeRefundRes struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;1:消费;2:撤销;3:退货
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS/FAILURE/UNKNOW
	TradeTime       string `json:"tradeTime"`                 // 交易时间;格式:yyyy-MM-dd HH:mm:ss
	SourceTradeId   int64  `json:"sourceTradeId"`             // 原交易流水号
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int    `json:"tradeAmount"`               // 交易金额(单位:分)
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方流水号
	CashierId       int    `json:"cashierId,omitempty"`       // 收银员ID号
	AcquirerTransNo string `json:"acquirerTransNO,omitempty"` // 收单交易号
	ChannelTransNo  string `json:"channelTransNo,omitempty"`  // 上游渠道交易号
	MerchantName    string `json:"merchantName"`              // 商户名称
	MrchntCode      string `json:"mrchntCode,omitempty"`      // 交易商户编号
	TerCode         string `json:"terCode,omitempty"`         // 交易终端编号
	MerchantId      int    `json:"merchantId"`                // 平台商户号
	TerminalId      int    `json:"terminalId"`                // 平台终端号
	Fee             int    `json:"fee,omitempty"`             // 交易手续费
}
