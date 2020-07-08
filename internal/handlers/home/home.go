package home

import (
	"clevergo.tech/clevergo"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Register(router clevergo.Router) {
	router.Get("/", h.index)
}
