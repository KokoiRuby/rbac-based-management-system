package runtime

type RedisConfig struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
	// TODO: Validation?
	DB           int `yaml:"db"`
	PingTimeout  int `yaml:"pingTimeout"`
	PingInterval int `yaml:"pingInterval"`
}
