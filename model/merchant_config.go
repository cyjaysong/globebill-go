package model

type MerchantPayConfig struct {
	PayTypeId   int    `json:"payTypeId"`   // 支付类型,传1003
	ConfigType  int    `json:"configType"`  // 媒体值;1.支付目录配置;2.公众号appId配置;3.小程序appId配置
	ConfigValue string `json:"configValue"` // 配置值
}

type MerchantPayConfigResult struct {
	ConfigType  int    `json:"configType"`  // 媒体值;1.支付目录配置;2.公众号appId配置;3.小程序appId配置
	ConfigValue string `json:"configValue"` // 配置值
	InletStatus int    `json:"inletStatus"` // 进件状态;1.待进件,2.进件成功,3.进件失败,4.处理中
	InletResult string `json:"inletResult"` // 错误信息
}

// MerchantPayConfigReq 开通商户支付配置
type MerchantPayConfigReq struct {
	OutMerchantNo string              `json:"outMerchantNo"` // 外部商户号
	PayConfigList []MerchantPayConfig `json:"payConfigList"` // 支付配置
}

// MerchantPayConfigRes 开通商户支付配置
type MerchantPayConfigRes struct {
	PayConfigList []MerchantPayConfigResult `json:"payConfigList"` // 外部商户号
}

type MerchantBizConfig struct {
	BizId   int     `json:"bizId"`   // 业务类型,104:商户提现,105:子商户提现
	PerFee  float64 `json:"perFee"`  // 单笔费用
	FeeRate float64 `json:"feeRate"` // 手续费
}

// MerchantBizConfigReq 配置商户业务信息
type MerchantBizConfigReq struct {
	OutMerchantNo string              `json:"outMerchantNo"` // 外部商户号
	BizList       []MerchantBizConfig `json:"bizList"`       // 业务信息
}

// MerchantBizConfigRes 配置商户业务信息
type MerchantBizConfigRes struct{}

type MerchantBizConfigResult struct {
	BizId     int `json:"bizId"`     // 业务类型,104:商户提现,105:子商户提现
	BizStatus int `json:"bizStatus"` // 业务状态,1.待签约,2.签约成功,3.签约失败
}

// MerchantBizGetReq 获取商户业务信息
type MerchantBizGetReq struct {
	OutMerchantNo string `json:"outMerchantNo"` // 外部商户号
}

// MerchantBizGetRes 获取商户业务信息
type MerchantBizGetRes struct {
	OutMerchantNo string              `json:"outMerchantNo"` // 外部商户号
	BizList       []MerchantBizConfig `json:"bizList"`       // 业务信息
}
