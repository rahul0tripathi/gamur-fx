package controllers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (c *controller) GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	if u := chi.URLParam(r, "user"); u != "" {
		userId, err := strconv.Atoi(u)
		if err != nil {
			c.util.InternalServerError(w, r, "invalid payload", err)
			return
		}

		transactions, err := c.db.GetAllTransactions(userId)
		if err != nil {
			c.util.InternalServerError(w, r, "failed to get user transactions", err)
		} else {
			c.util.JSONResponse(w, r, transactions)
		}
	}
}
