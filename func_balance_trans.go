package globebill

import "github.com/cyjaysong/globebill-go/model"

// AccountSplit 账户分账
func (t Client) AccountSplit(reqBody model.AccountSplitReq) (resBody *model.BaseRes[model.AccountSplitRes], err error) {
	resBody = new(model.BaseRes[model.AccountSplitRes])
	if err = t.commonJsonPost("/balanceTrans/accountSplit", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// AccountSplitQuery 账户分账查询
func (t Client) AccountSplitQuery(reqBody model.AccountSplitQueryReq) (resBody *model.BaseRes[model.AccountSplitRes], err error) {
	resBody = new(model.BaseRes[model.AccountSplitRes])
	if err = t.commonJsonPost("/balanceTrans/accountSplit", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// BalanceGet 余额查询
func (t Client) BalanceGet(reqBody model.BalanceGetReq) (resBody *model.BaseRes[model.BalanceGetRes], err error) {
	resBody = new(model.BaseRes[model.BalanceGetRes])
	if err = t.commonJsonPost("/balanceTrans/getBalance", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// TxnDetailGet 账户变动查询
func (t Client) TxnDetailGet(reqBody model.TxnDetailGetReq) (resBody *model.BaseRes[model.TxnDetailGetRes], err error) {
	resBody = new(model.BaseRes[model.TxnDetailGetRes])
	if err = t.commonJsonPost("/balanceTrans/getTxnDetail", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantWithdraw 商户提现
func (t Client) MerchantWithdraw(reqBody model.MerchantWithdrawReq) (resBody *model.BaseRes[model.MerchantWithdrawRes], err error) {
	resBody = new(model.BaseRes[model.MerchantWithdrawRes])
	if err = t.commonJsonPost("/balanceTrans/merchantWithdraw", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantWithdrawQuery 商户提现查询
func (t Client) MerchantWithdrawQuery(reqBody model.MerchantWithdrawQueryReq) (resBody *model.BaseRes[model.MerchantWithdrawRes], err error) {
	resBody = new(model.BaseRes[model.MerchantWithdrawRes])
	if err = t.commonJsonPost("/balanceTrans/merchantWithdrawQuery", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// SubMerchantWithdraw 子商户提现
func (t Client) SubMerchantWithdraw(reqBody model.SubMerchantWithdrawReq) (resBody *model.BaseRes[model.MerchantWithdrawRes], err error) {
	resBody = new(model.BaseRes[model.MerchantWithdrawRes])
	if err = t.commonJsonPost("/balanceTrans/subMerchantWithdraw", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// SubMerchantWithdrawQuery 子商户提现查询
func (t Client) SubMerchantWithdrawQuery(reqBody model.SubMerchantWithdrawQueryReq) (resBody *model.BaseRes[model.MerchantWithdrawRes], err error) {
	resBody = new(model.BaseRes[model.MerchantWithdrawRes])
	if err = t.commonJsonPost("/balanceTrans/subMerchantWithdrawQuery", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantDesigHistory 出款历史查询
func (t Client) MerchantDesigHistory(reqBody model.MerchantDesigHistoryReq) (resBody *model.BaseRes[model.MerchantDesigHistoryRes], err error) {
	resBody = new(model.BaseRes[model.MerchantDesigHistoryRes])
	if err = t.commonJsonPost("/balanceTrans/getDesigHistory", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}
