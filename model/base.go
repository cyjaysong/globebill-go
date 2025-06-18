package model

// BaseRes 响应公共结构体
type BaseRes[T any] struct {
	Code   string `json:"code"`             // 响应码
	Msg    string `json:"msg"`              // 响应消息
	Result T      `json:"result,omitempty"` // 业务结果,JSON格式
}

// MerchantBaseRes 商户相关响应公共结构体
type MerchantBaseRes[T any] struct {
	Code   int    `json:"code"`             // 响应码
	Msg    string `json:"msg"`              // 响应消息
	Result T      `json:"result,omitempty"` // 业务结果,JSON格式
}

// NotifyReq 回调公共结构体
type NotifyReq[T any] struct {
	Code            string `json:"code"`                      // 响应码
	Msg             string `json:"msg"`                       // 返回信息
	SignatureMethod string `json:"signatureMethod,omitempty"` // 签名方式,本次接口请求使用的签名算法,固定:SM3WITHSM2,不参与签名
	Sign            string `json:"sign,omitempty"`            // 签名,响应参数data明文转成json格式后使用SM3WITHSM2的签名值,不参与签名
	Data            T      `json:"data,omitempty"`            // 响应实体
	MerchantId      string `json:"merchantId"`                // 商户编号,由系统生成的唯一商户标识,不参与签名
}
