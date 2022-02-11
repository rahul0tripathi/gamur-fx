package routes

import "github.com/go-chi/chi/v5"

func (r route) UserRouter(router chi.Router) {
	router.Post("/create", r.c.CreateNewUser)
	router.Get("/{user}", r.c.GetUser)
	router.Post("/signin", r.c.SignInUser)
	router.Get("/transactions/{user}", r.c.GetUserTransactions)
}
