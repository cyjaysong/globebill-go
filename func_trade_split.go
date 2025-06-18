package globebill

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/globebill-go/model"
)

// TradeSplit 交易分账
func (t Client) TradeSplit(reqBody model.TradeSplitReq) (resBody *model.BaseRes[model.TradeSplitRes], err error) {
	resBody = new(model.BaseRes[model.TradeSplitRes])
	if err = t.commonJsonPost("/split/tradeSplit", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// TradeSplitQuery 交易分账查询
func (t Client) TradeSplitQuery(reqBody model.TradeSplitQueryReq) (resBody *model.BaseRes[model.TradeSplitRes], err error) {
	resBody = new(model.BaseRes[model.TradeSplitRes])
	if err = t.commonJsonPost("/split/tradeSplitQuery", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// TradePushVerify 交易推送验签
func (t Client) TradePushVerify(AccessId, SignType, SignValue string, httpBody []byte) (notifyReq *model.TradePushReq, err error) {
	if AccessId != t.accessId {
		return nil, errors.New("响应内容接入商编号不一致")
	}
	if SignType != "SHA256withRSA" {
		return nil, errors.New("响应内容签名算法不一致")
	}
	if !t.Sha256WithRsaVerify(httpBody, SignValue) {
		return nil, errors.New("响应内容验签失败")
	}
	notifyReq = new(model.TradePushReq)
	if err = sonic.Unmarshal(httpBody, notifyReq); err != nil {
		return nil, err
	}
	return
}

// TradeNotifyVerify 交易通知验签
func (t Client) TradeNotifyVerify(AccessId, SignType, SignValue string, httpBody []byte) (notifyReq *model.TradeNotifyReq, err error) {
	if AccessId != t.accessId {
		return nil, errors.New("响应内容接入商编号不一致")
	}
	if SignType != "SHA256withRSA" {
		return nil, errors.New("响应内容签名算法不一致")
	}
	if !t.Sha256WithRsaVerify(httpBody, SignValue) {
		return nil, errors.New("响应内容验签失败")
	}
	notifyReq = new(model.TradeNotifyReq)
	if err = sonic.Unmarshal(httpBody, notifyReq); err != nil {
		return nil, err
	}
	return
}
