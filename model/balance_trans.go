package model

// AccountSplitReq 账户分账
type AccountSplitReq struct {
	OutTransId        string `json:"outTransId"`            // 第三方系统内部流水号,需唯一
	MerchantId        int64  `json:"merchantId"`            // 被分账平台商户号
	ReceiveMerchantId int64  `json:"receiveMerchantId"`     // 分账平台商户号
	SplitAmount       int64  `json:"splitAmount"`           // 分账金额,单位:分
	SplitRemark       string `json:"splitRemark,omitempty"` // 分账备注
}

// AccountSplitQueryReq 账户分账查询
type AccountSplitQueryReq struct {
	OutTransId string `json:"outTransId,omitempty"` // 第三方系统内部流水号,与splitId选其一
	SplitId    int64  `json:"splitId,omitempty"`    // 分账流水号,与outTransId选其一
}

// AccountSplitRes 账户分账结果
type AccountSplitRes struct {
	SplitId           int64  `json:"splitId"`                     // 分账流水号
	SplitTime         string `json:"splitTime"`                   // 分账时间,格式：yyyy-MM-dd HH:mm:ss
	MrchntCode        string `json:"mrchntCode,omitempty"`        // 被分账交易商户编号
	ReceiveMrchntCode string `json:"receiveMrchntCode,omitempty"` // 分账交易子商户编号
	SplitAmount       int64  `json:"splitAmount"`                 // 分账金额,单位:分
	SplitStatus       int    `json:"splitStatus"`                 // 分账状态,1:待分账,2:分账成功,3:分账失败,4:分账处理中
	AcqTransNo        string `json:"acqTransNO,omitempty"`        // 收单交易号
}

// BalanceGetReq 余额查询
type BalanceGetReq struct {
	MerchantId int64 `json:"merchantId"` // 平台商户号
}

// BalanceGetRes 余额查询
type BalanceGetRes struct {
	MrchntCode       string `json:"mrchntCode"`       // 交易商户编号
	SettleAccount    string `json:"settleAccount"`    // 结算账户
	AllAmount        int64  `json:"allAmount"`        // 总余额,单位:分
	SettleAmount     int64  `json:"settleAmount"`     // 结算户金额,单位:分
	BalanceAmount    int64  `json:"balanceAmount"`    // 余额户金额,单位:分
	AvailWithdrawAmt int64  `json:"availWithdrawAmt"` // 可提现金额,单位:分
}

// TxnDetailGetReq 账户变动查询
type TxnDetailGetReq struct {
	MerchantId int64  `json:"merchantId"`         // 平台商户号
	AcctType   int    `json:"acctType,omitempty"` // 账户类型,1:结算账户,2:余额账户,不传两者都查询
	IncomeType int    `json:"incomeType"`         // 变动类型,1:调增,2:调减
	Date       string `json:"date"`               // 日期,格式：yyyy-MM-dd
}

// TxnDetailGetRes 账户变动查询
type TxnDetailGetRes struct {
	Data []TxnDetailItem `json:"data"`
}

type TxnDetailItem struct {
	MerchantId    int64  `json:"merchantId"`    // 平台商户号
	MrchntCode    string `json:"mrchntCode"`    // 交易商户号
	FlowTypeId    int    `json:"flowTypeId"`    // 流水类型,1:分账;2:提现;3:代付子商户,4:结算
	TxnTime       string `json:"txnTime"`       // 出入账时间,格式: yyyy-MM-dd HH:mm:ss
	AcqTransNo    string `json:"acqTransNO"`    // 收单交易号
	AcctType      int    `json:"acctType"`      // 账户类型,1:结算账户,2:余额账户
	Amount        int64  `json:"amount"`        // 金额,单位:分
	BalanceAmount int64  `json:"balanceAmount"` // 余额,单位:分
	IncomeType    int    `json:"incomeType"`    // 变动类型,1:调增,2:调减
	FlowId        int64  `json:"flowId"`        // 流水号
}

// MerchantWithdrawReq 商户提现
type MerchantWithdrawReq struct {
	OutTransId     string `json:"outTransId"`               // 第三方系统内部流水号，需唯一
	MerchantId     int64  `json:"merchantId"`               // 平台商户号
	WithdrawAmount int64  `json:"withdrawAmount"`           // 提现金额,单位:分
	WithdrawRemark string `json:"withdrawRemark,omitempty"` // 提现
}

// MerchantWithdrawQueryReq 商户提现查询
type MerchantWithdrawQueryReq struct {
	OutTransId string `json:"outTransId"` // 第三方系统内部流水号,与withdrawId选其一
	WithdrawId int64  `json:"withdrawId"` // 提现流水号,与outTransId选其一
}

// SubMerchantWithdrawReq 子商户提现
type SubMerchantWithdrawReq struct {
	OutTransId     string `json:"outTransId"`               // 第三方系统内部流水号，需唯一
	MerchantId     int64  `json:"merchantId"`               // 平台商户号
	SubMerchantId  int64  `json:"subMerchantId"`            // 平台子商户号
	WithdrawAmount int64  `json:"withdrawAmount"`           // 提现金额,单位:分
	WithdrawRemark string `json:"withdrawRemark,omitempty"` // 提现
}

// SubMerchantWithdrawQueryReq 子商户提现查询
type SubMerchantWithdrawQueryReq struct {
	OutTransId string `json:"outTransId"` // 第三方系统内部流水号,与withdrawId选其一
	WithdrawId int64  `json:"withdrawId"` // 提现流水号,与outTransId选其一
}

// MerchantWithdrawRes 商户提现
type MerchantWithdrawRes struct {
	WithdrawId     int64  `json:"withdrawId"`              // 提现流水号
	WithdrawTime   string `json:"withdrawTime"`            // 提现时间,格式: yyyy-MM-dd HH:mm:ss
	MrchntCode     string `json:"mrchntCode,omitempty"`    // 交易商户编号
	SubMrchntCode  string `json:"subMrchntCode,omitempty"` // 交易子商户编号
	WithdrawAmount int64  `json:"withdrawAmount"`          // 提现金额,单位:分
	WithdrawFee    int64  `json:"withdrawFee"`             // 提现手续费,单位:分
	WithdrawStatus int    `json:"withdrawStatus"`          // 提现状态,1:待提现,2:提现成功,3:提现失败,4:提现处理中
	AcqTransNo     string `json:"acqTransNO,omitempty"`    // 收单交易号
}

// MerchantDesigHistoryReq 出款历史查询
type MerchantDesigHistoryReq struct {
	MerchantId int64  `json:"merchantId"` // 平台商户号
	StartTime  string `json:"startTime"`  // 开始时间,格式: yyyy-MM-dd HH:mm:ss,查询时间段最长30天;
	EndTime    string `json:"endTime"`    // 结束时间,格式: yyyy-MM-dd HH:mm:ss,查询时间段最长30天;
	PageIndex  int    `json:"pageIndex"`  // 页码,从0开始
	PageSize   int    `json:"pageSize"`   // 每页记录数,取值范围50-500
}

// MerchantDesigHistoryRes 出款历史查询
type MerchantDesigHistoryRes struct {
	AmtCount int                 `json:"amtCount"` // 总记录数
	SumAmt   int64               `json:"sumAmt"`   // 交易总金额,单位:分
	Records  []MerchantDesigItem `json:"records"`  // 划付记录
}

// MerchantDesigItem 出款历史
type MerchantDesigItem struct {
	Date       string `json:"date"`               // 划付时间,格式: yyyy-MM-dd HH:mm:ss
	ReqDate    string `json:"reqDate"`            // 出款申请时间,格式: yyyy-MM-dd HH:mm:ss
	Amt        int64  `json:"amt"`                // 提现金额,单位:分
	StaCd      string `json:"staCd"`              // 状态码,00:等待出款,02:出款已受理,03:出款成功,04:出款失败
	SettleType string `json:"settleType"`         // 类型,00:提现,01:结算
	FeeAmt     int64  `json:"feeAmt"`             // 手续费,单位:分
	RespMsg    string `json:"respMsg,omitempty"`  // 响应消息
	Channel    string `json:"channel,omitempty"`  // 出款通道
	CardNo     string `json:"cardNo,omitempty"`   // 入账账号
	OrderId    string `json:"orderId,omitempty"`  // 流水号
	AcctName   string `json:"acctName,omitempty"` // 收款人姓名
	BankNo     string `json:"bankNo,omitempty"`   // 银行名称
	StaMode    string `json:"staMode,omitempty"`  // 计算模式,0:常规模式,1:特殊模式
}
