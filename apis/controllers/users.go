package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type signInUserResponse struct {
	Success bool `json:"success"`
	UserId  int  `json:"user_id"`
}

func (c *controller) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var body createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		c.util.InternalServerError(w, r, "invalid payload", err)
		return
	}
	err := c.db.CreateUser(body.Username, body.Password)
	if err != nil {
		c.util.InternalServerError(w, r, "failed to create user", err)
		return
	}
	c.util.JSONResponse(w, r, "successfully created user")
}

func (c *controller) SignInUser(w http.ResponseWriter, r *http.Request) {
	var body createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		c.util.InternalServerError(w, r, "invalid payload", err)
		return
	}
	user, err := c.db.GetUserByUserName(body.Username)
	if err != nil {
		c.util.InternalServerError(w, r, "invalid payload", err)
		return
	}
	h := md5.Sum([]byte(body.Password))
	verified := c.db.VerifyUserPassword(user, hex.EncodeToString(h[:]))
	if !verified {
		c.util.InternalServerError(w, r, "failed to verify user", err)
	} else {
		c.util.JSONResponse(w, r, signInUserResponse{
			Success: true,
			UserId:  user,
		})
	}
}
func (c *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	if u := chi.URLParam(r, "user"); u != "" {
		i, err := strconv.Atoi(u)
		if err != nil {
			c.util.InternalServerError(w, r, "invalid payload", err)
			return
		}
		user, err := c.db.GetUser(i)
		if err != nil {
			c.util.InternalServerError(w, r, "invalid payload", err)
			return
		}
		c.util.JSONResponse(w, r, user)
	} else {
		c.util.InternalServerError(w, r, "invalid payload", nil)
	}
}
