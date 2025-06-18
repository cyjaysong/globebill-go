package globebill

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/globebill-go/model"
)

// UniPay 统一支付
func (t Client) UniPay(reqBody model.UniPayReq) (resBody *model.BaseRes[model.UniPayRes], err error) {
	resBody = new(model.BaseRes[model.UniPayRes])
	if err = t.commonJsonPost("/qrcode/pay", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// ScanPay 聚合码支付
func (t Client) ScanPay(reqBody model.ScanPayReq) (resBody *model.BaseRes[model.ScanPayRes], err error) {
	resBody = new(model.BaseRes[model.ScanPayRes])
	if err = t.commonJsonPost("/qrcode/create", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// AppletPay 服务商小程序支付
func (t Client) AppletPay(reqBody model.AppletPayReq) (resBody *model.BaseRes[model.AppletPayRes], err error) {
	resBody = new(model.BaseRes[model.AppletPayRes])
	if err = t.commonJsonPost("/qrcode/minipPayArgs", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// TradeQuery 交易查询
func (t Client) TradeQuery(reqBody model.TradeQueryReq) (resBody *model.BaseRes[model.TradeQueryRes], err error) {
	resBody = new(model.BaseRes[model.TradeQueryRes])
	if err = t.commonJsonPost("/qrcode/query", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// TradeRefund 交易退款
func (t Client) TradeRefund(reqBody model.TradeRefundReq) (resBody *model.BaseRes[model.TradeRefundRes], err error) {
	resBody = new(model.BaseRes[model.TradeRefundRes])
	if err = t.commonJsonPost("/qrcode/refund", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// TradeSplitNotifyVerify 交易分账通知验签
func (t Client) TradeSplitNotifyVerify(AccessId, SignType, SignValue string, httpBody []byte) (notifyReq *model.TradeSplitNotifyReq, err error) {
	if AccessId != t.accessId {
		return nil, errors.New("响应内容接入商编号不一致")
	}
	if SignType != "SHA256withRSA" {
		return nil, errors.New("响应内容签名算法不一致")
	}
	if !t.Sha256WithRsaVerify(httpBody, SignValue) {
		return nil, errors.New("响应内容验签失败")
	}
	notifyReq = new(model.TradeSplitNotifyReq)
	if err = sonic.Unmarshal(httpBody, notifyReq); err != nil {
		return nil, err
	}
	return
}
