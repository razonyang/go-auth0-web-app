package core

import (
	"errors"
	"net/http"

	"clevergo.tech/auth"
	"github.com/alexedwards/scs/v2"
)

type SessionAuthenticator struct {
	sessionManager *scs.SessionManager
}

func NewSessionAuthenticator(sessionManager *scs.SessionManager) *SessionAuthenticator {
	return &SessionAuthenticator{
		sessionManager: sessionManager,
	}
}

// Authenticates the current user.
func (a *SessionAuthenticator) Authenticate(r *http.Request, w http.ResponseWriter) (auth.Identity, error) {
	ctx := r.Context()
	user, ok := a.sessionManager.Get(ctx, "auth_user").(User)
	if !ok {
		return nil, errors.New("no logged on user")
	}
	if user.Expired() {
		return nil, errors.New("user token expired")
	}
	return user, nil
}

// Challenge generates challenges upon authentication failure.
func (a *SessionAuthenticator) Challenge(*http.Request, http.ResponseWriter) {
}
