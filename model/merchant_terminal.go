package model

// MerchantTerminalRegisterReq 终端进件-新增
type MerchantTerminalRegisterReq struct {
	OutMerchantNo        string     `json:"outMerchantNo"`        // 外部商户号
	OutTerminalNo        string     `json:"outTerminalNo"`        // 外部终端号
	TerminalName         string     `json:"terminalName"`         // 签过单名称
	DeviceSn             string     `json:"deviceSn,omitempty"`   // 设备号
	ContactTel           string     `json:"contactTel"`           // 联系电话
	InstallAddress       string     `json:"installAddress"`       // 安装地址
	InstallProvince      int64      `json:"installProvince"`      // 安装地址-省份
	InstallCity          int64      `json:"installCity"`          // 安装地址-城市
	InstallDistrict      int64      `json:"installDistrict"`      // 安装地址-区县
	IsOpenConsume        bool       `json:"isOpenConsume"`        // 交易权限-查询
	IsOpenRefund         bool       `json:"isOpenRefund"`         // 交易权限-退货
	IsOpenQueryTrade     bool       `json:"isOpenQueryTrade"`     // 交易权限-交易权限-消费类
	IsOpenPreauthorize   bool       `json:"isOpenPreauthorize"`   // 交易权限-预授权类
	IsOpenDebitCard      bool       `json:"isOpenDebitCard"`      // 卡类型权限-借记卡
	IsOpenCreditCard     bool       `json:"isOpenCreditCard"`     // 卡类型权限-贷记卡
	IsOpenSemiCreditCard bool       `json:"isOpenSemiCreditCard"` // 卡类型权限-准贷记卡
	IsOpenPrepaidCard    bool       `json:"isOpenPrepaidCard"`    // 卡类型权限-预付费卡
	IsOpenPreauthorizeQr bool       `json:"isOpenPreauthorizeQr"` // 交易权限-二维码预授权类
	FileList             []FileItem `json:"fileList"`             // 图片信息
}

// MerchantTerminalRegisterRes 终端进件-新增
type MerchantTerminalRegisterRes struct {
	OutTerminalNo         string `json:"outTerminalNo"`         // 外部终端号
	TerminalId            int64  `json:"terminalId"`            // 平台终端号
	TerCode               string `json:"terCode"`               // 收单机构终端号
	AuditStatus           int    `json:"auditStatus"`           // 审核状态;1:待审核;2:审核通过;3:审核拒绝
	AuditResult           string `json:"auditResult"`           // 审核结果
	TerminalStatus        int    `json:"terminalStatus"`        // 终端状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
	IsAvailableDeviceBind bool   `json:"isAvailableDeviceBind"` // 是否有效的设备绑定
}

// MerchantTerminalUpdateReq 终端进件-修改
type MerchantTerminalUpdateReq struct {
	OutTerminalNo        string     `json:"outTerminalNo,omitempty"` // 外部终端号;与terminalId选其一
	TerminalId           int64      `json:"terminalId,omitempty"`    // 平台终端号;与outTerminalNo选其一
	TerminalName         string     `json:"terminalName"`            // 签过单名称
	ContactTel           string     `json:"contactTel"`              // 联系电话
	InstallAddress       string     `json:"installAddress"`          // 安装地址
	InstallProvince      int64      `json:"installProvince"`         // 安装地址-省份
	InstallCity          int64      `json:"installCity"`             // 安装地址-城市
	InstallDistrict      int64      `json:"installDistrict"`         // 安装地址-区县
	IsOpenConsume        bool       `json:"isOpenConsume"`           // 交易权限-查询
	IsOpenRefund         bool       `json:"isOpenRefund"`            // 交易权限-退货
	IsOpenQueryTrade     bool       `json:"isOpenQueryTrade"`        // 交易权限-交易权限-消费类
	IsOpenPreauthorize   bool       `json:"isOpenPreauthorize"`      // 交易权限-预授权类
	IsOpenDebitCard      bool       `json:"isOpenDebitCard"`         // 卡类型权限-借记卡
	IsOpenCreditCard     bool       `json:"isOpenCreditCard"`        // 卡类型权限-贷记卡
	IsOpenSemiCreditCard bool       `json:"isOpenSemiCreditCard"`    // 卡类型权限-准贷记卡
	IsOpenPrepaidCard    bool       `json:"isOpenPrepaidCard"`       // 卡类型权限-预付费卡
	IsOpenPreauthorizeQr bool       `json:"isOpenPreauthorizeQr"`    // 交易权限-二维码预授权类
	FileList             []FileItem `json:"fileList"`                // 图片信息
}

// MerchantTerminalUpdateRes 终端进件-修改
type MerchantTerminalUpdateRes struct {
	OutTerminalNo  string `json:"outTerminalNo"`  // 外部终端号
	TerminalId     int64  `json:"terminalId"`     // 平台终端号
	TerCode        string `json:"terCode"`        // 收单机构终端号
	AuditStatus    int    `json:"auditStatus"`    // 审核状态;1:待审核;2:审核通过;3:审核拒绝
	AuditResult    string `json:"auditResult"`    // 审核结果
	TerminalStatus int    `json:"terminalStatus"` // 终端状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
}

// MerchantTerminalGetReq 终端信息-获取
type MerchantTerminalGetReq struct {
	OutTerminalNo string `json:"outTerminalNo,omitempty"` // 外部终端号;与terminalId选其一
	TerminalId    int64  `json:"terminalId,omitempty"`    // 平台终端号;与outTerminalNo选其一
}

// MerchantTerminalGetRes 终端信息-获取
type MerchantTerminalGetRes struct {
	OutTerminalNo         string `json:"outTerminalNo"`         // 外部终端号
	TerminalId            int64  `json:"terminalId"`            // 平台终端号
	TerCode               string `json:"terCode"`               // 收单机构终端号
	AuditStatus           int    `json:"auditStatus"`           // 审核状态;1:待审核;2:审核通过;3:审核拒绝
	AuditResult           string `json:"auditResult"`           // 审核结果
	TerminalStatus        int    `json:"terminalStatus"`        // 终端状态;1.未开通;2.已开通;3.冻结;4.注销;5.拒绝;6.恢复
	IsAvailableDeviceBind bool   `json:"isAvailableDeviceBind"` // 是否有效的设备绑定
}

// MerchantBindDeviceReq 绑定设备
type MerchantBindDeviceReq struct {
	OutTerminalNo string `json:"outTerminalNo"` // 外部终端号
	DeviceSn      int64  `json:"deviceSn"`      // 设备号
}

// MerchantBindDeviceRes 绑定设备
type MerchantBindDeviceRes struct{}

// MerchantUnbindDeviceReq 解绑设备
type MerchantUnbindDeviceReq struct {
	OutTerminalNo string `json:"outTerminalNo"` // 外部终端号
	DeviceSn      int64  `json:"deviceSn"`      // 设备号
}

// MerchantUnbindDeviceRes 解绑设备
type MerchantUnbindDeviceRes struct{}
