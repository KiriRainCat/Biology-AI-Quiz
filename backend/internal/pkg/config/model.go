package config

type Configuration struct {
	Server     Server     `mapstructure:"server" json:"server" yaml:"server"`
	Postgresql Postgresql `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
	OpenAI     OpenAI     `mapstructure:"openai" json:"openai" yaml:"openai"`
}

type Server struct {
	Port        string `mapstructure:"port" json:"port,omitempty" yaml:"port"`
	RequestAuth string `mapstructure:"request_auth" json:"request_auth" yaml:"request_auth"`
}

type Postgresql struct {
	Host     string `mapstructure:"host" json:"host,omitempty" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port,omitempty" yaml:"port"`
	Db       string `mapstructure:"db" json:"db,omitempty" yaml:"db"`
	User     string `mapstructure:"user" json:"user,omitempty" yaml:"user"`
	Password string `mapstructure:"password" json:"password,omitempty" yaml:"password"`
}

type OpenAI struct {
	Token   string `mapstructure:"token" json:"token" yaml:"token"`
	BaseUrl string `mapstructure:"base_url" json:"base_url" yaml:"base_url"`
}
