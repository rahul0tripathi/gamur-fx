package config

type Config interface {
	GetDatabaseConfig() *DatabaseConfig
}
type config struct {
	DatabaseConfig *DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Username string `json:"username"`
	Database string `json:"database"`
}

func NewConfig() Config {
	c := &config{DatabaseConfig: &DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		Password: "hello1234",
		Username: "root",
		Database: "gamur",
	}}
	return c
}

func (c *config) GetDatabaseConfig() *DatabaseConfig {
	return c.DatabaseConfig
}
