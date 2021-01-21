package config

type TracerSettings struct {
	AgentHost   string `mapstructure:"agent_host"`
	AgentPort   int    `mapstructure:"agent_port"`
	SamplerHost string `mapstructure:"sampler_host"`
	SamplerPort int    `mapstructure:"sampler_port"`
}
