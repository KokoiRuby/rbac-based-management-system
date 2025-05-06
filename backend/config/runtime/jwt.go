package runtime

type JWT struct {
	Expire        int    `yaml:"expire"`
	RefreshExpire int    `yaml:"refreshExpire"`
	Issuer        string `yaml:"issuer"`
	SecretKey     string `yaml:"secretKey"`
}
