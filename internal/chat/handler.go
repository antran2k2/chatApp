package chat

import (
	"chatApp/internal/session"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection", http.StatusInternalServerError)
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan Message),
	}

	hub.Register <- client
	go client.ReadPump()
	go client.WritePump()
}
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	user, err := session.GetUserFromSession(r)
	if err != nil || user == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	data := struct {
		User string
	}{
		User: user,
	}

	tmpl, err := template.ParseFiles("web/static/chat.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}
