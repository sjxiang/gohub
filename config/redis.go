package config


type Redis struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	Database    int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
}
