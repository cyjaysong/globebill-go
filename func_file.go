package globebill

import "github.com/cyjaysong/globebill-go/model"

// FileUpload 上传媒体图片
func (t Client) FileUpload(reqBody model.FileUploadReq) (resBody *model.MerchantBaseRes[model.FileUploadRes], err error) {
	resBody = new(model.MerchantBaseRes[model.FileUploadRes])
	if err = t.merchantJsonPost("/Api/Access/FileInfo/Upload", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}
