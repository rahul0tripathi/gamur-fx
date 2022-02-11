package controllers

import (
	"github.com/rahul0tripathi/gamur/config"
	"github.com/rahul0tripathi/gamur/database"
	"github.com/rahul0tripathi/gamur/util"
	"net/http"
)

type controller struct {
	db     database.Database
	util   util.Util
	config config.Config
}

type Controllers interface {
	CreateNewUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	SignInUser(http.ResponseWriter, *http.Request)
	NewBattle(w http.ResponseWriter, r *http.Request)
	SubmitResult(w http.ResponseWriter, r *http.Request)
	GetPlayerBattles(w http.ResponseWriter, r *http.Request)
	GetUserTransactions(w http.ResponseWriter, r *http.Request)
	GetLeaderboard(w http.ResponseWriter, r *http.Request)
}

func NewController(db database.Database, u util.Util, cfg config.Config) Controllers {
	c := &controller{}
	c.db = db
	c.util = u
	c.config = cfg
	return c
}
