package globebill

import "github.com/cyjaysong/globebill-go/model"

// PreTradeFreezing 预授权交易请求
func (t Client) PreTradeFreezing(reqBody model.PreTradeFreezingReq) (resBody *model.BaseRes[model.PreTradeFreezingRes], err error) {
	resBody = new(model.BaseRes[model.PreTradeFreezingRes])
	if err = t.commonJsonPost("/webauth/freezing", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// PreTradeUnfreezing 预授权交易撤销
func (t Client) PreTradeUnfreezing(reqBody model.PreTradeUnfreezingReq) (resBody *model.BaseRes[model.PreTradeUnfreezingRes], err error) {
	resBody = new(model.BaseRes[model.PreTradeUnfreezingRes])
	if err = t.commonJsonPost("/webauth/unfreezing", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// PreTradeDeduct 预授权完成
func (t Client) PreTradeDeduct(reqBody model.PreTradeDeductReq) (resBody *model.BaseRes[model.PreTradeDeductRes], err error) {
	resBody = new(model.BaseRes[model.PreTradeDeductRes])
	if err = t.commonJsonPost("/webauth/deduct", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// PreTradeDeductRefund 预授权完成退款
func (t Client) PreTradeDeductRefund(reqBody model.PreTradeDeductRefundReq) (resBody *model.BaseRes[model.PreTradeDeductRefundRes], err error) {
	resBody = new(model.BaseRes[model.PreTradeDeductRefundRes])
	if err = t.commonJsonPost("/webauth/deductRefund", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// PreTradeQuery 预授权查询
func (t Client) PreTradeQuery(reqBody model.PreTradeQueryReq) (resBody *model.BaseRes[model.PreTradeQueryRes], err error) {
	resBody = new(model.BaseRes[model.PreTradeQueryRes])
	if err = t.commonJsonPost("/webauth/query", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}
