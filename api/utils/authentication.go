package utils

import (
	"github.com/alexsuslov/gobb/api/models"
	"github.com/alexsuslov/gobb/pkg/config"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"net/http"
)

var Store *sessions.CookieStore

func GetCookieStore(r *http.Request) *sessions.CookieStore {
	if Store == nil {
		cookie_key, _ := config.Config.GetString("gobb", "cookie_key")
		Store = sessions.NewCookieStore([]byte(cookie_key))
	}

	return Store
}

func GetCurrentUser(r *http.Request) *models.User {
	cached := context.Get(r, "user")
	if cached != nil {
		return cached.(*models.User)
	}

	session, _ := GetCookieStore(r).Get(r, "sirsid")

	if session.Values["username"] == nil || session.Values["password"] == nil {
		return nil
	}
	err, current_user := models.AuthenticateUser(session.Values["username"].(string), session.Values["password"].(string))

	if err != nil {
		return nil
	}

	context.Set(r, "user", current_user)
	return current_user
}
