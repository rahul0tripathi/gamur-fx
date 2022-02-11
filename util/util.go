package util

import (
	"github.com/go-chi/render"
	"github.com/rahul0tripathi/gamur/util/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type util struct {
	l *zap.SugaredLogger
}
type Util interface {
	InternalServerError(http.ResponseWriter, *http.Request, string, error)
	JSONResponse(http.ResponseWriter, *http.Request, interface{})
}

func NewUtil(l *zap.SugaredLogger) Util {
	u := &util{}
	u.l = l
	return u
}
func (u *util) JSONResponse(w http.ResponseWriter, r *http.Request, message interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, message)
}
func (u *util) InternalServerError(w http.ResponseWriter, r *http.Request, message string, err error) {
	if err != nil {
		u.l.Error(err)
	}
	if message == "" {
		message = "Internal Server Error"
	}
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, message)
}

var Module = fx.Provide(func(l *zap.SugaredLogger) Util {
	return NewUtil(l)
})
var Modules = fx.Options(
	logger.Module,
	Module,
)
