package user

import (
	"clevergo.tech/clevergo"
	"github.com/alexedwards/scs/v2"
)

type Handler struct {
	sessionManager *scs.SessionManager
}

func New(sessionManager *scs.SessionManager) *Handler {
	return &Handler{
		sessionManager: sessionManager,
	}
}

func (h *Handler) Register(router clevergo.Router) {
	router.Get("/login", h.login)
	router.Get("/callback", h.callback)
	router.Get("/logout", h.logout)
}
