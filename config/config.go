package config

import (
	"crypto/rand"
	"math/big"
)

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
	max := int64(10)
	min := int64(2)
	bg := big.NewInt(max - min)

	// get big.Int between 0 and bg
	// in this case 0 to 20
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	// add n to min to support the passed in range
	return float64(n.Int64() + min)

	//return c.EntryFee
}
