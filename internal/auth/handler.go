package auth

import (
	"chatApp/config"
	"chatApp/internal/session"
	"html/template"
	"log"
	"net/http"
)

var state = "random-state"

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := config.GoogleOAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	userInfo, err := GetGoogleUserInfo(r.Context(), token)

	//client := config.GoogleOAuthConfig.Client(r.Context(), token)
	//resp, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	//if err != nil {
	//	http.Error(w, "Failed to get user info", http.StatusInternalServerError)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//var user map[string]interface{}
	//if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
	//	http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
	//	return
	//}
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	session.SaveUserToSession(w, r, userInfo.Email)

	http.Redirect(w, r, "/chat", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOAuthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	user, err := session.GetUserFromSession(r)
	if err != nil {
		http.Error(w, "Failed to get user from session", http.StatusInternalServerError)
		return
	}

	if user != "" {
		http.Redirect(w, r, "/chat", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("web/static/index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session.ClearUserSession(w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
