package runtime

type SMTPConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Code       string `yaml:"code"`
	SkipVerify bool   `yaml:"skipVerify"`
}
