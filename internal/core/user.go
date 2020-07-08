package core

import (
	"encoding/gob"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func init() {
	gob.Register(User{})
}

type User struct {
	jwt.StandardClaims
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Nickname      string `json:"nickname"`
	Picture       string `json:"picture"`
	AccessToken   string `json:"-"`
	IDToken       string `json:"-"`
}

func (u User) GetID() string {
	return u.Subject
}

func (u User) Expired() bool {
	return time.Now().Unix() > u.ExpiresAt
}
