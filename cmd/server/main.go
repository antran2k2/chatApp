package main

import (
	"chatApp/internal/auth"
	"chatApp/internal/chat"
	"log"
	"net/http"
)

func main() {
	// Khởi tạo WebSocket hub
	chat.StartHub()

	// Routes
	http.HandleFunc("/", auth.IndexHandler) // Trang chủ
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/auth/callback", auth.CallbackHandler)
	http.HandleFunc("/chat", chat.ChatHandler)
	http.HandleFunc("/ws", chat.WebSocketHandler)
	http.HandleFunc("/logout", auth.LogoutHandler)

	// Chạy server
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
