package main

import (
	"context"
	"github.com/rahul0tripathi/gamur/apis/controllers"
	"github.com/rahul0tripathi/gamur/apis/routes"
	"github.com/rahul0tripathi/gamur/config"
	"github.com/rahul0tripathi/gamur/database"
	"github.com/rahul0tripathi/gamur/server"
	"github.com/rahul0tripathi/gamur/util"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func register(lifecycle fx.Lifecycle, l *zap.SugaredLogger, d database.Database, s server.Server) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				l.Info("starting up server")
				d.SetupDatabase()
				//d.CreateUser("test","helloiw")
				//d.GetUser(0)
				//d.GetUser(1)
				//err := d.DeductBalance(500.35,1,"user for game x")
				//fmt.Println(d.GetAllTransactions(1))
				//fmt.Println(d.NewBattle([]int{1},5.00,1))
				//fmt.Println(d.GetUserBattles(1))
				//fmt.Println(d.UpdatePlayerResult(1,199,1))
				//fmt.Println(d.GetTopPlayers())
				l.Info("setup successful")
				go s.Run()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				l.Info("server kiled")
				return nil
			},
		},
	)

}
func main() {
	fx.New(
		util.Modules,
		fx.Provide(
			config.NewConfig,
			database.NewDatabase,
			controllers.NewController,
			routes.NewRouter,
			server.NewServer,
		),
		fx.Invoke(register),
	).Run()
}
