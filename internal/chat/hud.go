package chat

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}

var hub = Hub{
	Clients:    make(map[*Client]bool),
	Broadcast:  make(chan Message),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

// StartHub bắt đầu hub và xử lý tin nhắn
func StartHub() {
	go func() {
		for {
			select {
			case client := <-hub.Register:
				hub.Clients[client] = true
			case client := <-hub.Unregister:
				delete(hub.Clients, client)
			case message := <-hub.Broadcast:
				for client := range hub.Clients {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(hub.Clients, client)
					}
				}
			}
		}
	}()
}
