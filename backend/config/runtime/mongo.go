package runtime

type MongoConfig struct {
	User              string `yaml:"user"`
	Pass              string `yaml:"pass"`
	Host              string `yaml:"host"`
	App               string `yaml:"app"`
	ConnectionTimeout int    `yaml:"connectionTimeout"`
	PingTimeout       int    `yaml:"pingTimeout"`
	PingInterval      int    `yaml:"pingInterval"`
}
