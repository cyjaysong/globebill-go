package model

// PreTradeFreezingReq 预授权交易请求
type PreTradeFreezingReq struct {
	SN           string             `json:"sn,omitempty"`           // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId   int64              `json:"merchantId,omitempty"`   // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId   int64              `json:"terminalId,omitempty"`   // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	TradeAmount  int64              `json:"tradeAmount"`            // 交易金额(单位:分)
	PayModeId    int                `json:"payModeId"`              // 支付方式
	OutTransId   string             `json:"outTransId"`             // 第三方系统内部流水号
	AppAccessId  int64              `json:"appAccessId,omitempty"`  // POS外部接入商编号,用于APP跳转POS收银台场景,outTransId为APP方流水号
	TradeRemark  string             `json:"tradeRemark,omitempty"`  // 交易备注
	CashierId    int64              `json:"cashierId,omitempty"`    // 收银员ID号
	GoodsSubject string             `json:"goodsSubject,omitempty"` // 商品名称(超过30字符截断)
	UserRemark   string             `json:"userRemark,omitempty"`   // 用户备注
	ExpireTime   int64              `json:"expireTime,omitempty"`   // 超时时间(支付超时时间,单位:秒,取值范围60-1800,默认300:5分钟)
	PayCode      string             `json:"payCode,omitempty"`      // 付款凭证码(payModeId为10033/10042/10102时必填)
	NotifyURL    string             `json:"notifyUrl,omitempty"`    // 交易结果通知地址
	ClientIP     string             `json:"clientIp,omitempty"`     // 客户端IP(payModeId为非10033/10042/10102/10104时必传)
	Location     string             `json:"location,omitempty"`     // 位置信息
	TerminalInfo *TradeTerminalInfo `json:"terminalInfo,omitempty"` // 终端信息,payModeId为10033、10042、10102时传
}

// PreTradeFreezingRes 预授权交易请求
type PreTradeFreezingRes struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;4:预授权请求;5:预授权完成;6:预授权撤销;8:预授权完成退款
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS:成功;FAILURE:失败;UNKNOW:未知
	TradeTime       string `json:"tradeTime"`                 // 交易时间;格式: yyyy-MM-dd HH:mm:ss
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int64  `json:"tradeAmount"`               // 交易金额,单位:分
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方系统内部流水号
	TradeRemark     string `json:"tradeRemark,omitempty"`     // 交易备注
	CashierId       int64  `json:"cashierId,omitempty"`       // 收银员ID号
	GoodsSubject    string `json:"goodsSubject,omitempty"`    // 商品名称
	AcquirerTransNo string `json:"acquirerTransNO,omitempty"` // 收单交易号
	ChannelTransNo  string `json:"channelTransNo,omitempty"`  // 上游渠道交易号
	UserOpenId      string `json:"userOpenId,omitempty"`      // 微信open_id/支付宝user_id
	CodeURL         string `json:"codeUrl,omitempty"`         // 二维码链接,扫码支付返回
	MerchantName    string `json:"merchantName"`              // 商户名称
	MrchntCode      string `json:"mrchntCode,omitempty"`      // 交易商户编号
	TerCode         string `json:"terCode,omitempty"`         // 交易终端编号
	MerchantId      int64  `json:"merchantId"`                // 平台商户号
	TerminalId      int64  `json:"terminalId"`                // 平台终端号
	Fee             int64  `json:"fee,omitempty"`             // 交易手续费,单位:分
	PayTime         string `json:"payTime,omitempty"`         // 完成支付时间,目前仅支持微信/支付宝;格式:yyyy-MM-dd HH:mm:ss
}

// PreTradeUnfreezingReq 预授权交易撤销
type PreTradeUnfreezingReq struct {
	SN            string `json:"sn,omitempty"`          // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId    int64  `json:"merchantId,omitempty"`  // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId    int64  `json:"terminalId,omitempty"`  // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	SourceTradeId int64  `json:"sourceTradeId"`         // 原交易流水号
	OutTransId    string `json:"outTransId"`            // 第三方系统内部流水号
	AppAccessId   int64  `json:"appAccessId,omitempty"` // POS外部接入商编号,用于APP跳转POS收银台场景,outTransId为APP方流水号
	TradeRemark   string `json:"tradeRemark,omitempty"` // 交易备注
	CashierId     int64  `json:"cashierId,omitempty"`   // 收银员ID号
}

// PreTradeUnfreezingRes 预授权交易撤销
type PreTradeUnfreezingRes struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;4:预授权请求;5:预授权完成;6:预授权撤销;8:预授权完成退款
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS:成功;FAILURE:失败;UNKNOW:未知
	TradeTime       string `json:"tradeTime"`                 // 交易时间;格式: yyyy-MM-dd HH:mm:ss
	SourceTradeId   int64  `json:"sourceTradeId"`             // 原交易流水号
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int64  `json:"tradeAmount"`               // 交易金额,单位:分
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方系统内部流水号
	CashierId       int64  `json:"cashierId,omitempty"`       // 收银员ID号
	AcquirerTransNo string `json:"acquirerTransNO,omitempty"` // 收单交易号
	ChannelTransNo  string `json:"channelTransNo,omitempty"`  // 上游渠道交易号
	MerchantName    string `json:"merchantName"`              // 商户名称
	MrchntCode      string `json:"mrchntCode,omitempty"`      // 交易商户编号
	TerCode         string `json:"terCode,omitempty"`         // 交易终端编号
	MerchantId      int64  `json:"merchantId"`                // 平台商户号
	TerminalId      int64  `json:"terminalId"`                // 平台终端号
}

// PreTradeDeductReq 预授权交易完成
type PreTradeDeductReq struct {
	SN            string `json:"sn,omitempty"`           // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId    int64  `json:"merchantId,omitempty"`   // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId    int64  `json:"terminalId,omitempty"`   // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	TradeAmount   int64  `json:"tradeAmount"`            // 交易金额(单位:分)
	SourceTradeId int64  `json:"sourceTradeId"`          // 原交易流水号
	OutTransId    string `json:"outTransId"`             // 第三方系统内部流水号
	AppAccessId   int64  `json:"appAccessId,omitempty"`  // POS外部接入商编号,用于APP跳转POS收银台场景,outTransId为APP方流水号
	CashierId     int64  `json:"cashierId,omitempty"`    // 收银员ID号
	GoodsSubject  string `json:"goodsSubject,omitempty"` // 商品名称
	ExpireTime    int64  `json:"expireTime,omitempty"`   // 超时时间(支付超时时间,单位:秒,取值范围60-1800,默认300:5分钟)
	PayCode       string `json:"payCode,omitempty"`      // 付款凭证码(payModeId为10033/10042/10102时必填)
	NotifyURL     string `json:"notifyUrl,omitempty"`    // 交易结果通知地址
	ClientIP      string `json:"clientIp,omitempty"`     // 客户端IP(payModeId为非10033/10042/10102/10104时必传)
	Location      string `json:"location,omitempty"`     // 位置信息
}

// PreTradeDeductRes 预授权交易完成
type PreTradeDeductRes struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;4:预授权请求;5:预授权完成;6:预授权撤销;8:预授权完成退款
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS:成功;FAILURE:失败;UNKNOW:未知
	TradeTime       string `json:"tradeTime"`                 // 交易时间;格式: yyyy-MM-dd HH:mm:ss
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int64  `json:"tradeAmount"`               // 交易金额,单位:分
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方系统内部流水号
	TradeRemark     string `json:"tradeRemark,omitempty"`     // 交易备注
	CashierId       int64  `json:"cashierId,omitempty"`       // 收银员ID号
	GoodsSubject    string `json:"goodsSubject,omitempty"`    // 商品名称
	UserRemark      string `json:"userRemark,omitempty"`      // 用户备注
	AcquirerTransNo string `json:"acquirerTransNO,omitempty"` // 收单交易号
	ChannelTransNo  string `json:"channelTransNo,omitempty"`  // 上游渠道交易号
	UserOpenId      string `json:"userOpenId,omitempty"`      // 微信open_id/支付宝user_id
	MerchantName    string `json:"merchantName"`              // 商户名称
	MrchntCode      string `json:"mrchntCode,omitempty"`      // 交易商户编号
	TerCode         string `json:"terCode,omitempty"`         // 交易终端编号
	MerchantId      int64  `json:"merchantId"`                // 平台商户号
	TerminalId      int64  `json:"terminalId"`                // 平台终端号
	Fee             int64  `json:"fee,omitempty"`             // 交易手续费,单位:分
	PayTime         string `json:"payTime,omitempty"`         // 完成支付时间,目前仅支持微信/支付宝;格式:yyyy-MM-dd HH:mm:ss
}

// PreTradeDeductRefundReq 预授权交易完成退款
type PreTradeDeductRefundReq struct {
	SN            string `json:"sn,omitempty"`          // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId    int64  `json:"merchantId,omitempty"`  // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId    int64  `json:"terminalId,omitempty"`  // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	SourceTradeId int64  `json:"sourceTradeId"`         // 原交易流水号
	OutTransId    string `json:"outTransId"`            // 第三方系统内部流水号
	AppAccessId   int64  `json:"appAccessId,omitempty"` // POS外部接入商编号,用于APP跳转POS收银台场景,outTransId为APP方流水号
	TradeAmount   int64  `json:"tradeAmount"`           // 退款金额(单位:分)
	TradeRemark   string `json:"tradeRemark,omitempty"` // 交易备注
	CashierId     int64  `json:"cashierId,omitempty"`   // 收银员ID号
}

// PreTradeDeductRefundRes 预授权交易完成退款
type PreTradeDeductRefundRes struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;4:预授权请求;5:预授权完成;6:预授权撤销;8:预授权完成退款
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS:成功;FAILURE:失败;UNKNOW:未知
	TradeTime       string `json:"tradeTime"`                 // 交易时间;格式:yyyy-MM-dd HH:mm:ss
	SourceTradeId   int64  `json:"sourceTradeId"`             // 原交易流水号
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int64  `json:"tradeAmount"`               // 交易金额,单位:分
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方系统内部流水号
	CashierId       int64  `json:"cashierId,omitempty"`       // 收银员ID号
	AcquirerTransNo string `json:"acquirerTransNO,omitempty"` // 收单交易号
	ChannelTransNo  string `json:"channelTransNo,omitempty"`  // 上游渠道交易号
	MerchantName    string `json:"merchantName"`              // 商户名称
	MrchntCode      string `json:"mrchntCode,omitempty"`      // 交易商户编号
	TerCode         string `json:"terCode,omitempty"`         // 交易终端编号
	MerchantId      int64  `json:"merchantId"`                // 平台商户号
	TerminalId      int64  `json:"terminalId"`                // 平台终端号
}

// PreTradeQueryReq 预授权交易查询
type PreTradeQueryReq struct {
	TradeId     int64  `json:"tradeId,omitempty"`     // 交易流水号,交易流水号和第三方流水号二选一,建议优先使用交易流水号
	OutTransId  string `json:"outTransId,omitempty"`  // 第三方系统内部流水号,交易流水号和第三方流水号二选一,建议优先使用交易流水号
	AppAccessId int64  `json:"appAccessId,omitempty"` // POS外部接入商编号,用于APP跳转POS收银台场景,outTransId为APP方流水号
}

// PreTradeQueryRes 预授权交易查询
type PreTradeQueryRes struct {
	TradeId         int64  `json:"tradeId"`                   // 交易流水号
	TradeType       int    `json:"tradeType"`                 // 交易类型;4:预授权请求;5:预授权完成;6:预授权撤销;8:预授权完成退款
	TradeStatus     string `json:"tradeStatus"`               // 交易状态;SUCCESS:成功;FAILURE:失败;UNKNOW:未知
	TradeTime       string `json:"tradeTime"`                 // 交易时间;格式:yyyy-MM-dd HH:mm:ss
	SourceTradeId   int64  `json:"sourceTradeId"`             // 原交易流水号
	SN              string `json:"sn"`                        // 设备TUSN号
	TradeAmount     int64  `json:"tradeAmount"`               // 交易金额,单位:分
	PayModeId       int    `json:"payModeId"`                 // 支付方式
	OutTransId      string `json:"outTransId"`                // 第三方系统内部流水号
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
	Fee             int64  `json:"fee,omitempty"`             // 交易手续费,单位:分
	PayTime         string `json:"payTime,omitempty"`         // 完成支付时间,目前仅支持微信/支付宝;格式:yyyy-MM-dd HH:mm:ss
}
