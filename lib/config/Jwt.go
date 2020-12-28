package config

type LoginAuthSettings struct {
	JwtSettings
	RefreshExpire int32 `mapstructure:"refresh_expire"`
}

type ActivateAuthSettings struct {
	JwtSettings
}

type JwtSettings struct {
	Secret string
	Key string
	Issuer string
	Expire int32
}
