package stock

import "net/http"
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{}

type HubS struct {
	clients      map[*Client]bool
	broadcast    chan []byte
	addClient    chan *Client
	removeClient chan *Client
}

var Hub = HubS{
	broadcast:    make(chan []byte),
	addClient:    make(chan *Client),
	removeClient: make(chan *Client),
	clients:      make(map[*Client]bool),
}

func (Hub *HubS) Start() {
	for {
		select {
		case conn := <-Hub.addClient:
			Hub.clients[conn] = true
		case conn := <-Hub.removeClient:
			if _, ok := Hub.clients[conn]; ok {
				delete(Hub.clients, conn)
				close(conn.send)
			}
		case message := <-Hub.broadcast:
			for conn := range Hub.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(Hub.clients, conn)
				}
			}
		}
	}
}

type Client struct {
	ws   *websocket.Conn
	send chan []byte
}

func (c *Client) write() {
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.ws.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) read() {
	defer func() {
		Hub.removeClient <- c
		c.ws.Close()
	}()

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			Hub.removeClient <- c
			c.ws.Close()
			break
		}

		Hub.broadcast <- message
	}
}

func Socket(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)

	if err != nil {
		http.NotFound(res, req)
		return
	}

	client := &Client{
		ws:   conn,
		send: make(chan []byte),
	}

	Hub.addClient <- client

	go client.write()
	go client.read()
}
