package model

// FileUploadReq 上传媒体图片
type FileUploadReq struct {
	FileData string `json:"fileData"` // 图片Base64
	FileUrl  string `json:"-"`        // 用于使用网络资源图片上传
}

// FileUploadRes 上传媒体图片
type FileUploadRes struct {
	MediaNo string `json:"mediaNo"` // 图片媒体ID
}
