package middleware

import (
	"net/http"

	"clevergo.tech/authmidware"
	"clevergo.tech/clevergo"
)

// IsAuthenticated redirects to the given URL if user isn't login.
func IsAuthenticated(url string, skipper clevergo.Skipper) clevergo.MiddlewareFunc {
	return func(next clevergo.Handle) clevergo.Handle {
		return func(c *clevergo.Context) error {
			if user := authmidware.GetIdentity(c.Context()); user == nil && !skipper(c) {
				return c.Redirect(http.StatusTemporaryRedirect, url)
			}
			return next(c)
		}
	}
}
