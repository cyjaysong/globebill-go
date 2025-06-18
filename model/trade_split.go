package model

// SplitReceive 分账接收方
type SplitReceive struct {
	MerchantId int64 `json:"merchantId"` // 分账商户号
	Amount     int64 `json:"amount"`     // 分账金额 单位分
}

// TradeSplitReq 交易分账
type TradeSplitReq struct {
	SN               string         `json:"sn,omitempty"`           // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId       int64          `json:"merchantId,omitempty"`   // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId       int64          `json:"terminalId,omitempty"`   // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	TradeId          int64          `json:"tradeId"`                // 交易流水号
	SplitReceiveList []SplitReceive `json:"splitReceiveList"`       // 分账接收方列表
	SplitRemarks     string         `json:"splitRemarks,omitempty"` // 分账备注
}

// TradeSplitQueryReq 交易分账查询
type TradeSplitQueryReq struct {
	SN           string `json:"sn,omitempty"`         // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId   int64  `json:"merchantId,omitempty"` // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId   int64  `json:"terminalId,omitempty"` // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	SplitOrderId int64  `json:"splitOrderId"`         // 交易分账流水号
}

// TradeSplitRes 交易分账
type TradeSplitRes struct {
	SplitOrderId    int64  `json:"splitOrderId"`          // 交易分账流水号
	SplitOrderStaus int    `json:"splitOrderStaus"`       // 分账状态;2:成功,3:失败,其它值:请调用查询接口查询状态
	SplitResult     string `json:"splitResult,omitempty"` // 分账结果描述
}
