package wechat

//Config example config
// type Config struct {
// 	Listen string `yaml:"listen"`
// 	Redis  struct {
// 		Host        string `yaml:"host"`
// 		Password    string `yaml:"password"`
// 		Database    int    `yaml:"database"`
// 		MaxActive   int    `yaml:"maxActive"`
// 		MaxIdle     int    `yaml:"maxIdle"`
// 		IdleTimeout int    `yaml:"idleTimeout"`
// 	} `yaml:"redis"`
// 	*OfficialAccountConfig `yaml:"officialAccountConfig"`
// }

// OfficialAccountConfig 公众号相关配置
type WeChatClientConfig struct {
	AppID          string `json:"app_id" yaml:"app_id"`
	AppSecret      string `json:"app_secret" yaml:"app_secret"`
	Token          string `json:"token" yaml:"token"`
	EncodingAESKey string `json:"encoding_aes_key" yaml:"encoding_aes_key"`
}
