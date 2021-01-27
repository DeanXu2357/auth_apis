package config

type SwaggerSettings struct {
	Host        string
	Title       string
	Description string
	Version     string
	BasePath    string `mapstructure:"base_path"`
}
