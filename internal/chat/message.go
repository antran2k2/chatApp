package chat

type Message struct {
	User    string `json:"user"`    // Tên người gửi
	Content string `json:"content"` // Nội dung tin nhắn
}
