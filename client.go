package globebill

import (
	reqclient "github.com/imroc/req/v3"
	"os"
	"time"
)

const alphabet = "0123456789ABCDEF"

//const commonProdApiUrl = "http://119.23.124.194:60002"
//const merchantProdApiUrl = "http://119.23.124.194:19073"

const commonProdApiUrl = "https://ppapi.globebill.com"
const merchantProdApiUrl = "https://ppaccess.globebill.com"

type Client struct {
	accessId          string // 接入商编号
	accessPrivateKey  []byte // 接入商私钥
	platformPublicKey []byte // 平台公钥
	reqClient         *reqclient.Client
	devMode           bool
}

func NewClient(accessId, accessPrivateKeyPath, platformPublicKeyPath string, devMode bool) (client *Client, err error) {
	privateBytes, err := os.ReadFile(accessPrivateKeyPath)
	if err != nil {
		return nil, err
	}
	publicFileBytes, err := os.ReadFile(platformPublicKeyPath)
	if err != nil {
		return nil, err
	}

	client = &Client{accessId: accessId, accessPrivateKey: privateBytes, platformPublicKey: publicFileBytes, devMode: devMode}
	client.reqClient = reqclient.C().SetTimeout(time.Second * 10).SetCommonRetryCount(1)
	client.reqClient.SetUserAgent("")
	if devMode {
		client.reqClient.DevMode()
	}
	return
}
