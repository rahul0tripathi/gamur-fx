package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rahul0tripathi/gamur/apis/controllers"
)

type route struct {
	c          controllers.Controllers
	baseRouter chi.Router
}
type Router interface {
	GetBaseRouter() chi.Router
}

func NewRouter(c controllers.Controllers) Router {
	r := &route{}
	r.baseRouter = chi.NewRouter()
	r.baseRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "ws://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))
	r.c = c
	r.BootstrapRouters()
	return r
}
func (r route) GetBaseRouter() chi.Router {
	return r.baseRouter
}
func (r route) BootstrapRouters() {
	r.baseRouter.Route("/user", r.UserRouter)
	r.baseRouter.Route("/battles", r.BattleRouter)
	r.baseRouter.Route("/leaderboard", r.LeaderboardRouter)
}
