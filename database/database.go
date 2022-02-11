package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rahul0tripathi/gamur/config"
	"go.uber.org/zap"
)

type Database interface {
	SetupDatabase()
	CreateUser(string, string) error
	GetUser(int) (User, error)
	DeductBalance(float64, int, string) error
	GetAllTransactions(int) ([]Txn, error)
	NewBattle([]int, float64, int) error
	GetUserBattles(int) ([]Battle, error)
	UpdatePlayerResult(int, int, int) error
	GetTopPlayers() ([]Leaderboard, error)
	VerifyUserPassword(int, string) bool
}
type database struct {
	db     *sql.DB
	schema *sql.Stmt
	logger *zap.SugaredLogger
}

func NewDatabase(c config.Config, l *zap.SugaredLogger) (Database, error) {
	dbConfig := c.GetDatabaseConfig()
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	d := &database{}
	d.logger = l
	db, err := sql.Open("mysql", connString)
	if err != nil {
		l.Errorf("failed to connect to database %s", dbConfig.Database)
		l.Error(err)
		return nil, err
	}
	db.SetMaxOpenConns(60)
	db.SetMaxIdleConns(5)
	d.db = db
	return d, nil
}
func (d *database) SetupDatabase() {
	for _, v := range SchemaList {
		x, err := d.db.Prepare(v)
		if err != nil {
			d.logger.Error(err)
			return
		}
		_, err = x.Exec()
		if err != nil {
			d.logger.Error("failed to setup database", err)
		}
	}
}
