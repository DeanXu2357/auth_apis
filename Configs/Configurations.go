package Configs

type Configurations struct {
	Server ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBName string
	DBUser string
	DBPassword string
}
