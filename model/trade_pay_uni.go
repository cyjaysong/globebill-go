package model

// UniPayReq 统一支付
type UniPayReq struct {
	SN               string             `json:"sn,omitempty"`               // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId       int64              `json:"merchantId,omitempty"`       // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId       int64              `json:"terminalId,omitempty"`       // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	TradeAmount      int64              `json:"tradeAmount"`                // 交易金额(单位:分)
	PayModeId        int                `json:"payModeId"`                  // 支付方式
	OutTransId       string             `json:"outTransId"`                 // 第三方流水号(需唯一)
	AppAccessId      int64              `json:"appAccessId,omitempty"`      // POS外部接入商编号(用于APP跳转POS收银台场景,outTransId为APP方流水号)
	TradeRemark      string             `json:"tradeRemark,omitempty"`      // 交易备注
	CashierId        int64              `json:"cashierId,omitempty"`        // 收银员ID号
	TradeSplitFlag   bool               `json:"tradeSplitFlag"`             // 是否分账交易
	GoodsSubject     string             `json:"goodsSubject,omitempty"`     // 商品名称(超过30字符截断)
	UserRemark       string             `json:"userRemark,omitempty"`       // 用户备注
	ExpireTime       int64              `json:"expireTime,omitempty"`       // 超时时间(支付超时时间,单位:秒,取值范围60-1800,默认300:5分钟)
	PayCode          string             `json:"payCode,omitempty"`          // 付款凭证码(payModeId为10033/10042/10102时必填)
	NotifyUrl        string             `json:"notifyUrl,omitempty"`        // 交易结果通知地址
	SplitNotifyUrl   string             `json:"splitNotifyUrl,omitempty"`   // 分账结果通知地址(tradeSplitFlag为true时有效)
	AppId            string             `json:"appId,omitempty"`            // 微信/支付宝appid(payModeId为10032/10036/10043时必传)
	UserOpenId       string             `json:"userOpenId,omitempty"`       // 微信open_id/支付宝user_id(payModeId为10032/10036/10043时必传)
	ClientIp         string             `json:"clientIp,omitempty"`         // 客户端IP(payModeId为非10033/10042/10102/10104时必传)
	Location         string             `json:"location,omitempty"`         // 位置信息
	GoodsTag         string             `json:"goodsTag,omitempty"`         // 订单优惠标记(微信支付有效,订单优惠标记,代金券或立减优惠功能的参数)
	InstalmentNum    string             `json:"instalmentNum,omitempty"`    // 分期数(选填,payModeId为10351时必填)
	SplitReceiveList []SplitReceive     `json:"splitReceiveList,omitempty"` // 分账接收方列表,收款方不需要作为分账接收方传入,tradeSplitFlag为true时生效,若不传,可以后续调用交易分账接口进行分账
	TerminalInfo     *TradeTerminalInfo `json:"terminalInfo,omitempty"`     // 终端信息(payModeId为10033/10042/10102时传)
}

type UniPayRes struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型(1:消费;2:撤销,3:退货)
	TradeStatus     string `json:"tradeStatus"`               // 交易状态(SUCCESS:成功;FAILURE:失败;UNKNOW:未知)
	TradeTime       string `json:"tradeTime"`                 // 交易时间(yyyy-MM-dd HH:mm:ss)
	SN              string `json:"sn,omitempty"`              // 设备TUSN号
	TradeAmount     int64  `json:"tradeAmount"`               // 交易金额(单位:分)
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方流水号,第三方系统内部流水号
	TradeRemark     string `json:"tradeRemark,omitempty"`     // 交易备注
	CashierId       int64  `json:"cashierId,omitempty"`       // 收银员ID号
	SplitFlag       bool   `json:"splitFlag,omitempty"`       // 是否分账交易
	GoodsSubject    string `json:"goodsSubject,omitempty"`    // 商品名称
	UserRemark      string `json:"userRemark,omitempty"`      // 用户备注
	AcquirerTransNo string `json:"acquirerTransNO,omitempty"` // 收单交易号
	ChannelTransNo  string `json:"channelTransNo,omitempty"`  // 上游渠道交易号
	UserOpenId      string `json:"userOpenId,omitempty"`      // 微信open_id/支付宝user_id
	CodeURL         string `json:"codeUrl,omitempty"`         // 二维码链接,扫码支付时返回
	PayInfo         string `json:"payInfo,omitempty"`         // 公众号/小程序支付信息,微信公众号/小程序支付时返回,公众号/小程序支付所需参数组成的json串
	TradeNo         string `json:"tradeNo,omitempty"`         // 支付宝交易号,payModeId为10043时返回支付窗支付所需参数,10103时返回记录tn信息
	MerchantName    string `json:"merchantName"`              // 商户名称
	MrchntCode      string `json:"mrchntCode,omitempty"`      // 交易商户编号
	TerCode         string `json:"terCode,omitempty"`         // 交易终端编号
	MerchantId      int64  `json:"merchantId"`                // 平台商户号
	TerminalId      int64  `json:"terminalId"`                // 平台终端号
	Fee             int64  `json:"fee,omitempty"`             // 交易手续费
	PayTime         string `json:"payTime,omitempty"`         // 完成支付时间,目前仅支持微信/支付宝;格式:yyyy-MM-dd HH:mm:ss
	ChannelMerNo    string `json:"channelMerNo,omitempty"`    // 渠道商户号
	CardType        string `json:"cardType,omitempty"`        // 借贷标识(00:借记卡;01:贷记卡)
}
