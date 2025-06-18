package model

// TradeBillSummaryReq 对账单汇总
type TradeBillSummaryReq struct {
	BillDate string `json:"billDate"` // 对账日期,格式: yyyy-MM-dd
}

// TradeBillSummaryRes 对账单汇总
type TradeBillSummaryRes struct {
	TradeTotalAmount   int64 `json:"tradeTotalAmount"`   // 交易总金额;单位:分
	TradeCount         int   `json:"tradeCount"`         // 交易总笔数
	IndirectTradeCount int64 `json:"indirectTradeCount"` // 间联交易总笔数
	DirectTradeCount   int64 `json:"directTradeCount"`   // 直联交易总笔数
}

// TradeBillDetailReq 对账单列表
type TradeBillDetailReq struct {
	BillDate  string `json:"billDate"`  // 对账日期,格式: yyyy-MM-dd
	PageIndex int    `json:"pageIndex"` // 页码,从1开始
	PageSize  int    `json:"pageSize"`  // 每页交易笔数,取值范围10-500
}

// TradeBillDetailRes 对账单列表
type TradeBillDetailRes struct {
	Datas []TradeBillSummaryRes `json:"datas"`
}

type TradeBillDetailItem struct {
	BillId                int64  `json:"billId"`                          // 账单号
	AcquirerTransNO       string `json:"acquirerTransNO"`                 // 钱宝流水号
	BatchNO               string `json:"batchNO,omitempty"`               // 批次号
	TradeTypeId           int    `json:"tradeTypeId"`                     // 交易类型;1:消费;2:撤销;3:退货;4:预授权请求;5:预授权完成;6:预授权撤销;8:预授权完成退款
	MrchntCode            string `json:"mrchntCode"`                      // 钱宝商户号
	TerCode               string `json:"terCode"`                         // 钱宝终端号
	TradeAmount           int64  `json:"tradeAmount"`                     // 交易金额(单位:分)
	ReceiveAmount         int64  `json:"receiveAmount"`                   // 实收金额(已减去手续费)
	TradeFee              int    `json:"tradeFee,omitempty"`              // 手续费
	ReceiveFee            int    `json:"receiveFee,omitempty"`            // 实收手续费
	PayTime               string `json:"payTime"`                         // 支付时间;格式:yyyy-MM-dd HH:mm:ss
	PayTypeId             int    `json:"payTypeId"`                       // 支付类型:1006:银行卡;1003:微信;1004:支付宝;1010:云闪付;1034:数字人民币
	PayAccount            string `json:"payAccount,omitempty"`            // 支付账号
	CardType              int    `json:"cardType,omitempty"`              // 卡类型;0:未知;1:借记卡;2:贷记卡;3:准贷记卡;4:预付费卡
	ReferenceNO           string `json:"referenceNO,omitempty"`           // 参考号
	AuthCode              string `json:"authCode,omitempty"`              // 授权码
	SourceAcquirerTransNO string `json:"sourceAcquirerTransNO,omitempty"` // 源流水号;直联时该值为空
	SourceBatchNO         string `json:"sourceBatchNO,omitempty"`         // 源批次号
	SourceTradePayDate    string `json:"sourceTradePayDate,omitempty"`    // 源交易日期;格式:yyyy-MM-dd HH:mm:ss
	AgencyId              int    `json:"agencyId"`                        // 代理号
	MerchantId            int    `json:"merchantId"`                      // 平台商户号
	TerminalId            int    `json:"terminalId"`                      // 平台终端号
	TradeId               int64  `json:"tradeId"`                         // 交易流水号
	IsDirect              bool   `json:"isDirect"`                        // 是否银联直联交易
	MerchantName          string `json:"merchantName"`                    // 商户名称
	OutTransNO            string `json:"outTransNO,omitempty"`            // 第三方交易流水号
	UnionPayMrchntCode    string `json:"unionPayMrchntCode,omitempty"`    // 银联直联商户编号
	UnionPayTerCode       string `json:"unionPayTerCode,omitempty"`       // 银联直联终端编号
}
