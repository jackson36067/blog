package ws

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WsHandler(c *gin.Context) {
	userIdStr := c.Query("userId")
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// 建立连接
	client := &Client{
		UserId: uint(userId),
		Conn:   conn,
	}
	WsManager.AddClient(client)
	// 心跳 + 监听关闭
	go handleConn(client)
}

func handleConn(client *Client) {
	defer func() {
		client.Conn.Close()
		WsManager.RemoveClient(client.UserId)
	}()

	client.Conn.SetReadLimit(512)
	client.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	client.Conn.SetPongHandler(func(string) error {
		client.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, _, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
