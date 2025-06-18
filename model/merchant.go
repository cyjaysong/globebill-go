package model

// AddressInfo 地址信息
type AddressInfo struct {
	Province int    `json:"province"` // 省份
	City     int    `json:"city"`     // 城市
	District int    `json:"district"` // 地区
	Address  string `json:"address"`  // 商户地址
}

// SettleAccount 结算账户/备案账户
type SettleAccount struct {
	BankAccountType     int    `json:"bankAccountType"`               // 账户性质;1:对公;2:对私,netType为3时,只能入对私
	BankAccountNo       string `json:"bankAccountNo"`                 // 银行账号
	BankAccountName     string `json:"bankAccountName"`               // 账号名称
	BankCode            int    `json:"bankCode"`                      // 银行类型
	OpenBank            string `json:"openBank"`                      // 开户行
	BankLineNo          string `json:"bankLineNo"`                    // 联行号
	AcctMobile          string `json:"acctMobile,omitempty"`          // 预留手机,bankAccountType为2时必填
	AccountIdentityType int    `json:"accountIdentityType,omitempty"` // 证件类型,bankAccountType为2时必填
	AccountIdentityNo   string `json:"accountIdentityNo,omitempty"`   // 证件号,bankAccountType为2时必填
}

// FileItem 图片信息
type FileItem struct {
	FileTypeId int    `json:"fileTypeId"` // 图片类型
	MediaNo    string `json:"mediaNo"`    // 媒体值
}

//// 商户相关

// MerchantRegisterReq 商户进件-新增
type MerchantRegisterReq struct {
	OutMerchantNo      string         `json:"outMerchantNo"`               // 外部商户号
	NetType            int            `json:"netType"`                     // 入网类类型,1:企业商户,2:个体工商户,3:小微商户
	WxServiceNo        string         `json:"wxServiceNo,omitempty"`       // 微信服务商号,指定微信服务商号,不送则用合作伙伴关联的服务商
	WxAppId            string         `json:"wxAppId,omitempty"`           // 微信AppId,当微信服务商号不为空必填
	AliPid             string         `json:"aliPId,omitempty"`            // 支付宝PID,指定支付宝服务商号,不送则用合作伙伴关联的服务商
	MerchantName       string         `json:"merchantName"`                // 商户名称,netType为3,格式为:XX市XX区/县个体XXX(主营),如:上海市长宁区个体王小明餐饮
	ShortName          string         `json:"shortName"`                   // 商户简称,netType为3,格式为:XX市XX区/县个体XXX(主营),如:上海市长宁区个体王小明餐饮
	IsDirectConnect    bool           `json:"isDirectConnect"`             // 是否直联接入,true:是;false:否
	UnionPayInsNo      string         `json:"unionPayInsNo,omitempty"`     // 所属银联机构号,isDirectConnect为true必填
	UnionPayMrchntNo   string         `json:"unionPayMrchntNo,omitempty"`  // 银联商户号,isDirectConnect为true必填
	SettlementMode     int            `json:"settlementMode,omitempty"`    // 银联结算标识,isDirectConnect为true时默认传2
	LicenceType        int            `json:"licenceType,omitempty"`       // 执照类型,netType为1,2时必填
	LicenceNo          string         `json:"licenceNo,omitempty"`         // 执照编号,netType为1,2时必填
	LicenceExpireDate  string         `json:"licenceExpireDate,omitempty"` // 执照有效期,netType为1,2时必填,长期传2199-12-31
	ManageTypeId       int            `json:"manageTypeId"`                // 经营类型
	ManageScope        string         `json:"manageScope"`                 // 经营范围
	LepCertName        string         `json:"lepCertName"`                 // 法人名称
	LepCertType        int            `json:"lepCertType"`                 // 证件类型
	LepCertNo          string         `json:"lepCertNo"`                   // 法人证件编号
	LepCertExpireDate  string         `json:"lepCertExpireDate"`           // 法人证件有效期,长期传2199-12-31
	ContactName        string         `json:"contactName"`                 // 联系人
	ContactMobile      string         `json:"contactMobile"`               // 联系人手机号
	RegisteredCapital  int            `json:"registeredCapital"`           // 注册资金,单位:万元
	Remarks            string         `json:"remarks,omitempty"`           // 备注
	IsUseElecAgreement bool           `json:"isUseElecAgreement"`          // 是否电子协议签约
	AddressInfo        AddressInfo    `json:"addressInfo"`                 // 地址信息
	SettlementType     int            `json:"settlementType"`              // 清算方式,2:D+0,4:D+1
	AdvanceInterest    float64        `json:"advanceInterest"`             // 垫付日息,没值传0,提现费(按比例)
	WithdrawFee        float64        `json:"withdrawFee"`                 // 提现手续费,没值传0,提现费(固定值)
	FullThreshold      float64        `json:"fullThreshold"`               // 满额阀值,D+1时,没值传0;单笔出款金额,取值范围:10-1000000
	SettlementHour     int            `json:"settlementHour"`              // 结算时间点,0:0点;8:8点;12:12点;18:18点
	WithdrawType       int            `json:"withdrawType"`                // 提现方式,1:自动提现;2:手工提现
	SettleAccount      SettleAccount  `json:"settleAccount"`               // 结算账户
	FilingAccount      *SettleAccount `json:"filingAccount,omitempty"`     // 备案账户
	// netType为1,结算账户bankAccountType为2时filingAccount必填;
	// netType为2,结算账户bankAccountType为2且lepCertName与bankAccountName不一致时filingAccount必填;
	CardPay       any              `json:"cardPay,omitempty"`       // 银行卡费率信息
	QuickPassPay  any              `json:"quickPassPay,omitempty"`  // 银联扫码费率信息,1000以下优惠费率;非优惠的,用银行卡的
	AliPay        any              `json:"aliPay,omitempty"`        // 支付宝费率信息
	WxPay         any              `json:"wxPay,omitempty"`         // 微信费率信息
	WildCardList  any              `json:"wildCardList,omitempty"`  // 外卡费率信息
	TerminalList  []map[string]any `json:"terminalList"`            // 终端信息列表
	FileList      []FileItem       `json:"fileList,omitempty"`      // 图片信息
	OtherMediaNos []string         `json:"otherMediaNos,omitempty"` // 其他图片信息
}

// MerchantRegisterRes 商户进件-新增
type MerchantRegisterRes struct {
	OutMerchantNo         string `json:"outMerchantNo"`         // 外部商户号
	MerchantId            int    `json:"merchantId"`            // 平台商户号
	MrchntCode            string `json:"mrchntCode"`            // 收单机构商户号
	AuditStatus           int    `json:"auditStatus"`           // 审核状态;1:待审核;2:审核通过;3:审核拒绝
	AuditResult           string `json:"auditResult"`           // 审核结果
	MerchantStatus        int    `json:"merchantStatus"`        // 商户状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
	IsAvailableDeviceBind bool   `json:"isAvailableDeviceBind"` // 是否有效的设备绑定
}

// MerchantUpdateReq 商户进件-修改
type MerchantUpdateReq struct {
	OutMerchantNo      string         `json:"outMerchantNo,omitempty"`     // 外部商户号;与merchantId选其一
	MerchantId         int            `json:"merchantId,omitempty"`        // 平台商户号;与outMerchantNo选其一
	MerchantName       string         `json:"merchantName"`                // 商户名称,netType为3,格式为:XX市XX区/县个体XXX(主营),如:上海市长宁区个体王小明餐饮
	ShortName          string         `json:"shortName"`                   // 商户简称,netType为3,格式为:XX市XX区/县个体XXX(主营),如:上海市长宁区个体王小明餐饮
	IsDirectConnect    bool           `json:"isDirectConnect"`             // 是否直联接入,true:是;false:否
	UnionPayInsNo      string         `json:"unionPayInsNo,omitempty"`     // 所属银联机构号,isDirectConnect为true必填
	UnionPayMrchntNo   string         `json:"unionPayMrchntNo,omitempty"`  // 银联商户号,isDirectConnect为true必填
	SettlementMode     int            `json:"settlementMode,omitempty"`    // 银联结算标识,isDirectConnect为true时默认传2
	LicenceType        int            `json:"licenceType,omitempty"`       // 执照类型,netType为1,2时必填
	LicenceNo          string         `json:"licenceNo,omitempty"`         // 执照编号,netType为1,2时必填
	LicenceExpireDate  string         `json:"licenceExpireDate,omitempty"` // 执照有效期,netType为1,2时必填,长期传2199-12-31
	ManageTypeId       int            `json:"manageTypeId"`                // 经营类型
	ManageScope        string         `json:"manageScope"`                 // 经营范围
	LepCertName        string         `json:"lepCertName"`                 // 法人名称
	LepCertType        int            `json:"lepCertType"`                 // 证件类型
	LepCertNo          string         `json:"lepCertNo"`                   // 法人证件编号
	LepCertExpireDate  string         `json:"lepCertExpireDate"`           // 法人证件有效期,长期传2199-12-31
	ContactName        string         `json:"contactName"`                 // 联系人
	ContactMobile      string         `json:"contactMobile"`               // 联系人手机号
	RegisteredCapital  int            `json:"registeredCapital"`           // 注册资金,单位:万元
	Remarks            string         `json:"remarks,omitempty"`           // 备注
	IsUseElecAgreement bool           `json:"isUseElecAgreement"`          // 是否电子协议签约
	AddressInfo        AddressInfo    `json:"addressInfo"`                 // 地址信息
	SettlementType     int            `json:"settlementType"`              // 清算方式,2:D+0,4:D+1
	AdvanceInterest    float64        `json:"advanceInterest"`             // 垫付日息,没值传0,提现费(按比例)
	WithdrawFee        float64        `json:"withdrawFee"`                 // 提现手续费,没值传0,提现费(固定值)
	FullThreshold      float64        `json:"fullThreshold"`               // 满额阀值,D+1时,没值传0;单笔出款金额,取值范围:10-1000000
	SettlementHour     int            `json:"settlementHour"`              // 结算时间点,0:0点;8:8点;12:12点;18:18点
	WithdrawType       int            `json:"withdrawType"`                // 提现方式,1:自动提现;2:手工提现
	SettleAccount      SettleAccount  `json:"settleAccount"`               // 结算账户
	FilingAccount      *SettleAccount `json:"filingAccount,omitempty"`     // 备案账户
	// netType为1,结算账户bankAccountType为2时filingAccount必填;
	// netType为2,结算账户bankAccountType为2且lepCertName与bankAccountName不一致时filingAccount必填;
	CardPay       any        `json:"cardPay"`       // 银行卡费率信息
	QuickPassPay  any        `json:"quickPassPay"`  // 银联扫码费率信息,1000以下优惠费率;非优惠的,用银行卡的
	AliPay        any        `json:"aliPay"`        // 支付宝费率信息
	WxPay         any        `json:"wxPay"`         // 微信费率信息
	WildCardList  any        `json:"wildCardList"`  // 外卡费率信息
	TerminalList  any        `json:"terminalList"`  // 终端信息列表
	FileList      []FileItem `json:"fileList"`      // 图片信息
	OtherMediaNos []string   `json:"otherMediaNos"` // 其他图片信息
}

// MerchantUpdateRes 商户进件-修改
type MerchantUpdateRes struct {
	OutMerchantNo         string `json:"outMerchantNo"`         // 外部商户号
	MerchantId            int    `json:"merchantId"`            // 平台商户号
	MrchntCode            string `json:"mrchntCode"`            // 收单机构商户号
	AuditStatus           int    `json:"auditStatus"`           // 审核状态;1:待审核;2:审核通过;3:审核拒绝
	AuditResult           string `json:"auditResult"`           // 审核结果
	MerchantStatus        int    `json:"merchantStatus"`        // 商户状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
	IsAvailableDeviceBind bool   `json:"isAvailableDeviceBind"` // 是否有效的设备绑定
}

// MerchantGetReq 商户信息-获取
type MerchantGetReq struct {
	OutMerchantNo string `json:"outMerchantNo,omitempty"` // 外部商户号;与merchantId选其一
	MerchantId    int    `json:"merchantId,omitempty"`    // 平台商户号;与outMerchantNo选其一
}

// MerchantGetRes 商户信息-获取
type MerchantGetRes struct {
	OutMerchantNo  string `json:"outMerchantNo"`  // 外部商户号
	MerchantId     int    `json:"merchantId"`     // 平台商户号
	MrchntCode     string `json:"mrchntCode"`     // 收单机构商户号
	AuditStatus    int    `json:"auditStatus"`    // 审核状态;1:待审核;2:审核通过;3:审核拒绝
	AuditResult    string `json:"auditResult"`    // 审核结果
	MerchantStatus int    `json:"merchantStatus"` // 商户状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
}

// MerchantUpmchtGetReq 商户认证信息-获取
type MerchantUpmchtGetReq struct {
	OutMerchantNo string `json:"outMerchantNo"`
}

// MerchantUpmchtGetRes 商户认证信息-获取
type MerchantUpmchtGetRes struct {
	UpmchtList string `json:"upmchtList"` // 认证信息
}

type MerchantUpmchtItem struct {
	ChannelId        int    `json:"channelId"`        // 支付方式,101:微信,201:支付宝,301:银联扫码
	UpmchtNo         string `json:"upmchtNo"`         // 上游商户号
	ChannelServiceNo string `json:"channelServiceNo"` // 服务商号
	AuthorizeState   int    `json:"authorizeState"`   // 认证状态,1:认证中,2:认证成功,3:认证失败
	Enable           bool   `json:"enable"`           // 有效性,true:有效,false:失效
	CreateTime       string `json:"createTime"`       // 创建时间
}

//// 子商户相关

// SubMerchantRegisterReq  子商户进件-新增
type SubMerchantRegisterReq struct {
	OutMerchantNo     string        `json:"outMerchantNo"`     // 外部商户号
	OutSubMerchantNo  string        `json:"outSubMerchantNo"`  // 外部子商户号
	SubMerchantTypeId int           `json:"subMerchantTypeId"` // 子商户类型,2.合作商;3.员工
	SubMerchantName   string        `json:"subMerchantName"`   // 子商户名称
	LepCertName       string        `json:"lepCertName"`       // 法人名称
	Remarks           string        `json:"remarks,omitempty"` // 备注
	SettleAccount     SettleAccount `json:"settleAccount"`     // 结算账户
	FileList          []FileItem    `json:"fileList"`          // 图片信息
}

// SubMerchantRegisterRes  子商户进件-新增
type SubMerchantRegisterRes struct {
	OutMerchantNo     string `json:"outMerchantNo"`     // 外部商户号
	OutSubMerchantNo  string `json:"outSubMerchantNo"`  // 外部子商户号
	SubMerchantId     int    `json:"subMerchantId"`     // 平台子商户号
	SubMrchntCode     string `json:"subMrchntCode"`     // 收单机构子商户号
	AuditStatus       int    `json:"auditStatus"`       // 审核状态;1.待审核;2.审核通过;3.审核拒绝
	AuditResult       string `json:"auditResult"`       // 审核结果
	SubMerchantStatus int    `json:"subMerchantStatus"` // 子商户状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
}

// SubMerchantUpdateReq  子商户进件-修改
type SubMerchantUpdateReq struct {
	OutSubMerchantNo string        `json:"outSubMerchantNo"`  // 外部子商户号
	SubMerchantId    int           `json:"subMerchantId"`     // 平台子商户号
	SubMerchantName  string        `json:"subMerchantName"`   // 子商户名称
	LepCertName      string        `json:"lepCertName"`       // 法人名称
	Remarks          string        `json:"remarks,omitempty"` // 备注
	SettleAccount    SettleAccount `json:"settleAccount"`     // 结算账户
	FileList         []FileItem    `json:"fileList"`          // 图片信息
}

// SubMerchantUpdateRes  子商户进件-修改
type SubMerchantUpdateRes struct {
	OutSubMerchantNo  string `json:"outSubMerchantNo"`  // 外部子商户号
	SubMerchantId     int    `json:"subMerchantId"`     // 平台子商户号
	SubMrchntCode     string `json:"subMrchntCode"`     // 收单机构子商户号
	AuditStatus       int    `json:"auditStatus"`       // 审核状态;1.待审核;2.审核通过;3.审核拒绝
	AuditResult       string `json:"auditResult"`       // 审核结果
	SubMerchantStatus int    `json:"subMerchantStatus"` // 子商户状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
}

// SubMerchantGetReq  子商户信息-获取
type SubMerchantGetReq struct {
	OutSubMerchantNo string `json:"outSubMerchantNo"` // 外部子商户号
	SubMerchantId    int    `json:"subMerchantId"`    // 平台子商户号
}

// SubMerchantGetRes  子商户信息-获取
type SubMerchantGetRes struct {
	OutSubMerchantNo  string `json:"outSubMerchantNo"`  // 外部子商户号
	SubMerchantId     int    `json:"subMerchantId"`     // 平台子商户号
	SubMrchntCode     string `json:"subMrchntCode"`     // 收单机构子商户号
	AuditStatus       int    `json:"auditStatus"`       // 审核状态;1.待审核;2.审核通过;3.审核拒绝
	AuditResult       string `json:"auditResult"`       // 审核结果
	SubMerchantStatus int    `json:"subMerchantStatus"` // 子商户状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
}
