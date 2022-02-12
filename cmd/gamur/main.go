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
				l.Info("setup successful")

				go s.Run()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				l.Info("KILLING")
				//TODO: Add method to gracefully close server
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
