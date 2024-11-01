package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func GetUserFromSession(r *http.Request) (string, error) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return "", err
	}
	user, ok := session.Values["user"].(string)
	if !ok {
		return "", nil
	}
	return user, nil
}

func SaveUserToSession(w http.ResponseWriter, r *http.Request, user string) {
	session, _ := store.Get(r, "session-name")
	session.Values["user"] = user
	err := session.Save(r, w)
	if err != nil {
		return
	}
}

func ClearUserSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	delete(session.Values, "user")
	err := session.Save(r, w)
	if err != nil {
		return
	}
}
