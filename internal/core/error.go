package core

import (
	"net/http"

	"clevergo.tech/clevergo"
)

func ErrorHandler(next clevergo.Handle) clevergo.Handle {
	return func(c *clevergo.Context) error {
		if err := next(c); err != nil {
			e, ok := err.(clevergo.Error)
			if !ok {
				e = clevergo.NewError(http.StatusInternalServerError, err)
			}
			return c.Render(e.Status(), "home/error.tmpl", clevergo.Map{
				"err": e,
			})
		}

		return nil
	}
}
