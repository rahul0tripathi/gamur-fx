package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type BattleResponse struct {
	BattleId int64 `json:"battle_id"`
}
type BattleResult struct {
	Score    int `json:"score"`
	BattleId int `json:"battle_id"`
	Player   int `json:"player"`
}

func (c *controller) NewBattle(w http.ResponseWriter, r *http.Request) {
	if u := chi.URLParam(r, "user"); u != "" {
		userId, err := strconv.Atoi(u)
		if err != nil {
			c.util.InternalServerError(w, r, "invalid payload", err)
			return
		}
		err = c.db.DeductBalance(c.config.GetEntryFee(), userId, "Used For FlappyBird")
		if err != nil {
			c.util.InternalServerError(w, r, "insufficient balance", err)

		} else {
			battleId, e := c.db.NewBattle([]int{userId}, c.config.GetEntryFee(), 1)
			if e != nil {
				c.util.InternalServerError(w, r, "failed to create  battle", err)

			} else {
				c.util.JSONResponse(w, r, BattleResponse{BattleId: battleId})
			}
		}
	}
}

func (c *controller) SubmitResult(w http.ResponseWriter, r *http.Request) {
	var body BattleResult
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		c.util.InternalServerError(w, r, "invalid payload", err)
		return
	}
	err := c.db.UpdatePlayerResult(body.Player, body.Score, body.BattleId)
	if err != nil {
		c.util.InternalServerError(w, r, "failed to update battle", err)
	} else {
		c.util.JSONResponse(w, r, "succesfully updated battle")
	}
}

func (c *controller) GetPlayerBattles(w http.ResponseWriter, r *http.Request) {
	if u := chi.URLParam(r, "user"); u != "" {
		userId, err := strconv.Atoi(u)
		if err != nil {
			c.util.InternalServerError(w, r, "invalid payload", err)
			return
		}
		battles, err := c.db.GetUserBattles(userId)
		if err != nil {
			c.util.InternalServerError(w, r, "invalid payload", err)

		} else {
			c.util.JSONResponse(w, r, battles)
		}
	}
}
