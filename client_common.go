package globebill

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/deatil/go-cryptobin/cryptobin/rsa"
	"time"
)

func (t Client) getCommonApiUrl(path string) string {
	if t.prodEnv {
		return commonProdApiUrl + path
	}
	return commonTestApiUrl + path
}

func (t Client) getMerchantApiUrl(path string) string {
	if t.prodEnv {
		return merchantProdApiUrl + path
	}
	return merchantTestApiUrl + path
}

// 公共相关接口请求
func (t Client) commonJsonPost(path string, reqBody, resBody any) (err error) {
	bodyJsonBytes, _ := sonic.Marshal(reqBody)
	//if t.devMode {
	//	log.Println("data:", string(bodyJsonBytes))
	//}
	header := map[string]string{"AccessId": t.accessId, "Timestamp": time.Now().Format(time.DateTime),
		"SignType": "SHA256withRSA", "SignValue": t.Sha256WithRsaSign(bodyJsonBytes)}
	response, err := t.reqClient.R().SetHeaders(header).SetBodyJsonBytes(bodyJsonBytes).Post(t.getCommonApiUrl(path))
	if err != nil {
		return err
	}
	var bytes []byte
	if bytes, err = response.ToBytes(); err != nil {
		return err
	}
	if response.GetHeader("AccessId") != t.accessId {
		return errors.New("响应内容接入商编号不一致")
	}
	if response.GetHeader("SignType") != "SHA256withRSA" {
		return errors.New("响应内容签名算法不一致")
	}
	if !t.Sha256WithRsaVerify(bytes, response.GetHeader("SignValue")) {
		return errors.New("响应内容验签失败")
	}
	return sonic.Unmarshal(bytes, resBody)
}

// 商户相关接口请求
func (t Client) merchantJsonPost(path string, reqBody, resBody any) (err error) {
	bodyJsonBytes, _ := sonic.Marshal(reqBody)
	//if t.devMode {
	//	log.Println("data:", string(bodyJsonBytes))
	//}
	header := map[string]string{"AccessId": t.accessId, "Timestamp": time.Now().Format(time.DateTime),
		"SignType": "SHA256withRSA", "SignValue": t.Sha256WithRsaSign(bodyJsonBytes)}
	response, err := t.reqClient.R().SetHeaders(header).SetBodyJsonBytes(bodyJsonBytes).Post(t.getMerchantApiUrl(path))
	if err != nil {
		return err
	}
	var bytes []byte
	if bytes, err = response.ToBytes(); err != nil {
		return err
	}
	if response.GetHeader("AccessId") != t.accessId {
		return errors.New("响应内容接入商编号不一致")
	}
	if response.GetHeader("SignType") != "RSA256" {
		return errors.New("响应内容签名算法不一致")
	}
	if !t.Sha256WithRsaVerify(bytes, response.GetHeader("SignValue")) {
		return errors.New("响应内容验签失败")
	}
	return sonic.Unmarshal(bytes, resBody)
}

// Sha256WithRsaSign 签名
func (t Client) Sha256WithRsaSign(bytes []byte) (sign string) {
	obj := rsa.New().FromPrivateKey(t.accessPrivateKey).SetSignHash("SHA256")
	return obj.FromBytes(bytes).Sign().ToBase64String()
}

// Sha256WithRsaVerify 验签
func (t Client) Sha256WithRsaVerify(bytes []byte, sign string) (pass bool) {
	obj := rsa.New().FromPublicKey(t.platformPublicKey).SetSignHash("SHA256")
	return obj.FromBase64String(sign).Verify(bytes).ToVerify()
}
