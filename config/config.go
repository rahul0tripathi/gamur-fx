package config

type Config interface {
	GetDatabaseConfig() *DatabaseConfig
	GetEntryFee() float64
}
type config struct {
	DatabaseConfig *DatabaseConfig
	EntryFee       float64
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
	}, EntryFee: 2.50}
	return c
}

func (c *config) GetDatabaseConfig() *DatabaseConfig {
	return c.DatabaseConfig
}
func (c *config) GetEntryFee() float64 {
	return c.EntryFee
}
