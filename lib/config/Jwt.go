package config

type LoginAuthSettings struct {
	PrivateKey    string `mapstructure:"private_key"`
	PublicKey     string `mapstructure:"public_key"`
	Issuer        string
	Expire        int64
	RefreshExpire int64 `mapstructure:"refresh_expire"`
}

type ActivateAuthSettings struct {
	PrivateKey string `mapstructure:"private_key"`
	PublicKey  string `mapstructure:"public_key"`
	Issuer     string
	Expire     int64
}
