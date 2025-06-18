package model

// TradeTerminalInfo 交易终端信息
type TradeTerminalInfo struct {
	TerminalType   string `json:"terminalType"`   // 终端设备类型,见常量 consts.TerminalTypes
	SerialNum      string `json:"serialNum"`      // 终端设备的硬件序列号
	AppVersion     string `json:"appVersion"`     // 终端应用程序的版本号
	SecretText     string `json:"secretText"`     // 密文数据
	EncryptRandNum string `json:"encryptRandNum"` // 加密随机因子
}
