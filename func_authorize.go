package globebill

import "github.com/cyjaysong/globebill-go/model"

// FacePayAuthInfo 获取刷脸支付授权信息
func (t Client) FacePayAuthInfo(reqBody model.FacePayAuthInfoReq) (resBody *model.BaseRes[model.FacePayAuthInfoRes], err error) {
	resBody = new(model.BaseRes[model.FacePayAuthInfoRes])
	if err = t.commonJsonPost("/facePay/getAuthInfo", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// WechatScanPayGetOpenId 微信付款码获取OpenId
func (t Client) WechatScanPayGetOpenId(reqBody model.WechatScanPayGetOpenIdReq) (resBody *model.BaseRes[model.WechatScanPayGetOpenIdRes], err error) {
	resBody = new(model.BaseRes[model.WechatScanPayGetOpenIdRes])
	if err = t.commonJsonPost("/wechat/openIdQuery", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}

// UnionPayGetUserId 获取银联UserId
func (t Client) UnionPayGetUserId(reqBody model.UnionPayGetUserIdReq) (resBody *model.BaseRes[model.UnionPayGetUserIdRes], err error) {
	resBody = new(model.BaseRes[model.UnionPayGetUserIdRes])
	if err = t.commonJsonPost("/unionPay/getUserId", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}
