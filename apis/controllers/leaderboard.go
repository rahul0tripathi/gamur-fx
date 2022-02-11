package controllers

import "net/http"

func (c *controller) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	players, err := c.db.GetTopPlayers()
	if err != nil {
		c.util.InternalServerError(w, r, "", err)
	} else {
		c.util.JSONResponse(w, r, players)
	}
}
