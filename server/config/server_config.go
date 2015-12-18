package config

type ServerConfig struct {
	BindAddress string `json:"bind_address" yaml:"bind_address"`
	HttpPort    int    `json:"http_port" yaml:"http_port"`
	UiDir       string `json:"ui_dir" yaml:"ui_dir"`
}

func newServerConfig() *ServerConfig {
	return &ServerConfig{
		BindAddress: "0.0.0.0",
		HttpPort:    8080,
		UiDir:       "/opt/gotrack/share",
	}
}
