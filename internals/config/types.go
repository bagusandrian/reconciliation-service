package config

type Server struct {
	HostName string
	HostIP   string
	HTTP     ServiceConf `yaml:"http"`
	GRPC     ServiceConf `yaml:"grpc"`
	NSQ      ServiceConf `yaml:"nsq"`
	Cron     ServiceConf `yaml:"cron"`
}

// HTTP defines server config for http server
type ServiceConf struct {
	Port string `yaml:"port"`
}
