package routes

import "github.com/go-chi/chi/v5"

func (r route) LeaderboardRouter(router chi.Router) {
	router.Get("/", r.c.GetLeaderboard)
}
