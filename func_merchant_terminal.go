package globebill

import "github.com/cyjaysong/globebill-go/model"

// MerchantTerminalRegister 终端进件-新增
func (t Client) MerchantTerminalRegister(reqBody model.MerchantTerminalRegisterReq) (resBody *model.MerchantBaseRes[model.MerchantTerminalRegisterRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantTerminalRegisterRes])
	if err = t.merchantJsonPost("/Api/Access/Terminal/Register", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantTerminalUpdate 终端进件-修改
func (t Client) MerchantTerminalUpdate(reqBody model.MerchantTerminalUpdateReq) (resBody *model.MerchantBaseRes[model.MerchantTerminalUpdateRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantTerminalUpdateRes])
	if err = t.merchantJsonPost("/Api/Access/Terminal/Update", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantTerminalGet 终端信息-获取
func (t Client) MerchantTerminalGet(reqBody model.MerchantTerminalGetReq) (resBody *model.MerchantBaseRes[model.MerchantTerminalGetRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantTerminalGetRes])
	if err = t.merchantJsonPost("/Api/Access/Terminal/Get", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantBindDevice 绑定设备
func (t Client) MerchantBindDevice(reqBody model.MerchantBindDeviceReq) (resBody *model.MerchantBaseRes[model.MerchantBindDeviceRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantBindDeviceRes])
	if err = t.merchantJsonPost("/Api/Access/Terminal/BindDevice", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// MerchantUnbindDevice 解绑设备
func (t Client) MerchantUnbindDevice(reqBody model.MerchantUnbindDeviceReq) (resBody *model.MerchantBaseRes[model.MerchantUnbindDeviceRes], err error) {
	resBody = new(model.MerchantBaseRes[model.MerchantUnbindDeviceRes])
	if err = t.merchantJsonPost("/Api/Access/Terminal/UnbindDevice", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}
