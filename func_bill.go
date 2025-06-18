package globebill

import "github.com/cyjaysong/globebill-go/model"

// TradeBillSummary 对账单汇总
func (t Client) TradeBillSummary(reqBody model.TradeBillSummaryReq) (resBody *model.BaseRes[model.TradeBillSummaryRes], err error) {
	resBody = new(model.BaseRes[model.TradeBillSummaryRes])
	if err = t.commonJsonPost("/bill/getBillSummary", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// TradeBillDetail 对账单列表
func (t Client) TradeBillDetail(reqBody model.TradeBillDetailReq) (resBody *model.BaseRes[model.TradeBillDetailRes], err error) {
	resBody = new(model.BaseRes[model.TradeBillDetailRes])
	if err = t.commonJsonPost("/bill/getBillDetail", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}
