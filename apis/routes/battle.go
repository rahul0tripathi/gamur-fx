package routes

import "github.com/go-chi/chi/v5"

func (r route) BattleRouter(router chi.Router) {
	router.Post("/create/{user}", r.c.NewBattle)
	router.Post("/result", r.c.SubmitResult)
	router.Get("/{user}", r.c.GetPlayerBattles)
}
