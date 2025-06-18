package model

// FacePayAuthInfoReq 获取刷脸支付授权信息
type FacePayAuthInfoReq struct {
	SN         string `json:"sn,omitempty"`         // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId int    `json:"merchantId,omitempty"` // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId int    `json:"terminalId,omitempty"` // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	RawData    string `json:"rawData"`              // 设备信息
}

// FacePayAuthInfoRes 获取刷脸支付授权信息
type FacePayAuthInfoRes struct {
	AuthInfo  string `json:"authInfo"`            // 授权信息,授权信息请在客户端做缓存
	ExpiresIn int    `json:"expiresIn,omitempty"` // 授权信息有效时间,单位:秒,在有效时间内,同一台终端设备可以重复使用authInfo
	AppId     int    `json:"appid"`               // 服务商AppId
	MchId     string `json:"mchId"`               // 服务商编号
	SubAppId  string `json:"subAppid"`            // 子商户AppId
	SubMchId  string `json:"subMchId"`            // 子商户编号
}

// WechatScanPayGetOpenIdReq 微信付款码获取OpenId
type WechatScanPayGetOpenIdReq struct {
	SN         string `json:"sn,omitempty"`         // 设备TUSN号,有固定TUSN号可使用sn参数交易
	MerchantId int    `json:"merchantId,omitempty"` // 平台商户号,渠道对接可使用merchantId和terminalId参数交易
	TerminalId int    `json:"terminalId,omitempty"` // 平台终端号,渠道对接可使用merchantId和terminalId参数交易
	AuthCode   string `json:"authCode"`             // 	付款码
}

// WechatScanPayGetOpenIdRes 微信付款码获取OpenId
type WechatScanPayGetOpenIdRes struct {
	OpenId string `json:"openId"` // 微信OpenId
}

// UnionPayGetUserIdReq 银联获取UserId
type UnionPayGetUserIdReq struct {
	AppId    string `json:"appId,omitempty"` // 银联支付标识,HTTP请求User Agent中包含银联支付标识,格式为"UnionPay/<版本号>",注意APP标识仅支持字母和数字;示例:UnionPay/1.0 ICBCeLife
	UserCode string `json:"userCode"`        // 授权码
}

// UnionPayGetUserIdRes 银联获取UserId
type UnionPayGetUserIdRes struct {
	UserId string `json:"userId"` // 银联UserId
}
