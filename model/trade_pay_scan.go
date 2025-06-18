package model

// ScanPayReq 聚合码支付
type ScanPayReq struct {
	SN           string `json:"sn,omitempty"`           // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId   int64  `json:"merchantId,omitempty"`   // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId   int64  `json:"terminalId,omitempty"`   // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	TradeAmount  int64  `json:"tradeAmount"`            // 交易金额(单位:分)
	OutTransId   string `json:"outTransId"`             // 第三方流水号,需唯一
	AppAccessId  int64  `json:"appAccessId,omitempty"`  // POS外部接入商编号,用于APP跳转POS收银台场景,outTransId为APP方流水号
	TradeRemark  string `json:"tradeRemark,omitempty"`  // 交易备注
	CashierId    int64  `json:"cashierId,omitempty"`    // 收银员ID号
	GoodsSubject string `json:"goodsSubject,omitempty"` // 商品名称
	ExpireTime   int64  `json:"expireTime,omitempty"`   // 超时时间(支付超时时间,单位:秒,取值范围60-1800,默认300:5分钟)
	NotifyURL    string `json:"notifyUrl,omitempty"`    // 交易结果通知地址
	GoodsTag     string `json:"goodsTag,omitempty"`     // 订单优惠标记,微信支付有效,代金券或立减优惠功能的参数
}

// ScanPayRes 聚合码支付
type ScanPayRes struct {
	OutTransId string `json:"outTransId"` // 第三方流水号
	CodeUrl    string `json:"codeUrl"`    // 二维码链接
}
