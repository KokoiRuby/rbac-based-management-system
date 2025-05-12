package runtime

type UploadConfig struct {
	Avatar Avatar `json:"avatar"`
}

type Avatar struct {
	Size int64  `json:"size"`
	Dir  string `json:"dir"`
}
