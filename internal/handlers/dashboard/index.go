package dashboard

import (
	"net/http"

	"clevergo.tech/clevergo"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Register(router clevergo.Router) {
	router.Get("/dashboard", h.index)
}

func (h *Handler) index(c *clevergo.Context) error {
	return c.Render(http.StatusOK, "dashboard/index.tmpl", nil)
}
