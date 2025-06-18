package globebill

import "github.com/cyjaysong/globebill-go/model"

// MerchantRegister 商户进件-新增
func (t Client) MerchantRegister(reqBody model.MerchantRegisterReq) (resBody *model.MerchantBaseRes[model.MerchantRegisterRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantRegisterRes])
	if err = t.merchantJsonPost("/Api/Access/Merchant/Register", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantUpdate 商户进件-修改
func (t Client) MerchantUpdate(reqBody model.MerchantUpdateReq) (resBody *model.MerchantBaseRes[model.MerchantUpdateRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantUpdateRes])
	if err = t.merchantJsonPost("/Api/Access/Merchant/Update", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantGet 商户信息-获取
func (t Client) MerchantGet(reqBody model.MerchantGetReq) (resBody *model.MerchantBaseRes[model.MerchantGetRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantGetRes])
	if err = t.merchantJsonPost("/Api/Access/Merchant/Get", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// SubMerchantRegister 子商户进件-新增
func (t Client) SubMerchantRegister(reqBody model.SubMerchantRegisterReq) (resBody *model.MerchantBaseRes[model.SubMerchantRegisterRes], err error) {
	resBody = new(model.MerchantBaseRes[model.SubMerchantRegisterRes])
	if err = t.merchantJsonPost("/Api/Access/SubMerchant/Register", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// SubMerchantUpdate 子商户进件-修改
func (t Client) SubMerchantUpdate(reqBody model.SubMerchantUpdateReq) (resBody *model.MerchantBaseRes[model.SubMerchantUpdateRes], err error) {
	resBody = new(model.MerchantBaseRes[model.SubMerchantUpdateRes])
	if err = t.merchantJsonPost("/Api/Access/SubMerchant/Update", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// SubMerchantGet 子商户信息-获取
func (t Client) SubMerchantGet(reqBody model.SubMerchantGetReq) (resBody *model.MerchantBaseRes[model.SubMerchantGetRes], err error) {
	resBody = new(model.MerchantBaseRes[model.SubMerchantGetRes])
	if err = t.merchantJsonPost("/Api/Access/SubMerchant/Get", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantPayConfig 商户支付配置
func (t Client) MerchantPayConfig(reqBody model.MerchantPayConfigReq) (resBody *model.MerchantBaseRes[model.MerchantPayConfigRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantPayConfigRes])
	if err = t.merchantJsonPost("/Api/Access/Merchant/PayConfig/Apply", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantBizConfig 商户业务配置
func (t Client) MerchantBizConfig(reqBody model.MerchantBizConfigReq) (resBody *model.MerchantBaseRes[model.MerchantBizConfigRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantBizConfigRes])
	if err = t.merchantJsonPost("/Api/Access/Merchant/Biz/Apply", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantBizGet 商户业务配置获取
func (t Client) MerchantBizGet(reqBody model.MerchantBizGetReq) (resBody *model.MerchantBaseRes[model.MerchantBizGetRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantBizGetRes])
	if err = t.merchantJsonPost("/Api/Access/Merchant/Biz/Query", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantUpmchtGet 商户认证信息-获取
func (t Client) MerchantUpmchtGet(reqBody model.MerchantUpmchtGetReq) (resBody *model.MerchantBaseRes[model.MerchantUpmchtGetRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantUpmchtGetRes])
	if err = t.merchantJsonPost("/Api/Access/Upmcht/Get", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}
