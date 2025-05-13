package runtime

type AWS struct {
	S3 S3 `yaml:"s3"`
}

type S3 struct {
	Region    string `yaml:"region"`
	KeyID     string `yaml:"keyID"`
	AccessKey string `yaml:"accessKey"`
	Bucket    string `yaml:"bucket"`
	Avatar    Avatar `yaml:"avatar"`
}
