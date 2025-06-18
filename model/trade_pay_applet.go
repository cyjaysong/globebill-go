package model

// AppletPayReq 服务商小程序支付
type AppletPayReq struct {
	SN           string `json:"sn,omitempty"`           // 设备TUSN号;有固定TUSN号时使用
	MerchantId   int    `json:"merchantId,omitempty"`   // 平台商户号;渠道对接可使用merchantId和terminalId参数交易
	TerminalId   int    `json:"terminalId,omitempty"`   // 平台终端号;渠道对接可使用merchantId和terminalId参数交易
	TradeAmount  int    `json:"tradeAmount"`            // 交易金额;单位:分
	OutTransId   string `json:"outTransId"`             // 第三方系统内部流水号,需唯一
	AppAccessId  int    `json:"appAccessId,omitempty"`  // POS外部接入商编号;用于APP跳转POS收银台场景,outTransId为APP方流水号
	TradeRemark  string `json:"tradeRemark,omitempty"`  // 交易备注
	CashierId    int    `json:"cashierId,omitempty"`    // 收银员ID号
	SplitFlag    bool   `json:"splitFlag,omitempty"`    // 是否分账交易
	GoodsSubject string `json:"goodsSubject,omitempty"` // 商品名称
	ExpireTime   int    `json:"expireTime,omitempty"`   // 超时时间(支付超时时间,单位:秒,取值范围60-1800,默认300:5分钟)
	NotifyURL    string `json:"notifyUrl,omitempty"`    // 交易结果通知地址 (长度200)
	GoodsTag     string `json:"goodsTag,omitempty"`     // 订单优惠标记（微信支付专用）
}

// AppletPayRes 服务商小程序支付
type AppletPayRes struct {
	OutTransId string `json:"outTransId"` // 第三方系统内部流水号,需唯一
	TradeId    string `json:"tradeId"`    // 交易流水号
	AppId      string `json:"appId"`      // 服务商appId
	Path       string `json:"path"`       // 小程序支付路径
}
