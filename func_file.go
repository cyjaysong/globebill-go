package globebill

import (
	"fmt"
	"github.com/cyjaysong/globebill-go/model"
	"github.com/deatil/go-cryptobin/tool/encoding"
)

// FileUpload 上传媒体图片
func (t Client) FileUpload(reqBody model.FileUploadReq) (resBody *model.MerchantBaseRes[model.FileUploadRes], err error) {
	if reqBody.FileData == "" && reqBody.FileUrl != "" {
		if response, err := t.reqClient.R().Get(reqBody.FileUrl); err != nil {
			return nil, err
		} else if response.IsErrorState() {
			return nil, fmt.Errorf("load file use url err %s", response.Status)
		} else if len(response.Bytes()) == 0 {
			return nil, fmt.Errorf("load file use url fail")
		} else {
			reqBody.FileData = encoding.Base64Encode(response.Bytes())
		}
	}
	resBody = new(model.MerchantBaseRes[model.FileUploadRes])
	if err = t.merchantJsonPost("/Api/Access/FileInfo/Upload", reqBody, resBody); err != nil {
		return nil, err
	}
	return
}
