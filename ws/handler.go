package ws

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"blog/dto/response"
	"blog/enum"
	"blog/global"
	"blog/models"

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

	client := &Client{
		UserId: uint(userId),
		Conn:   conn,
	}

	WsManager.AddClient(client)
	go handleConn(client)
}

func handleConn(client *Client) {
	// 使用 defer 确保清理
	defer func() {
		WsManager.RemoveClient(client.UserId, client.Conn)
	}()

	client.Conn.SetReadLimit(1024) // 适当调大，防止消息被截断
	// 设置读取截止时间，如果 60s 内没收到任何消息，ReadJSON 会报错并退出循环
	_ = client.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	// 注意：删除了后端的 Ticker Ping，因为前端已经在发 Ping 了。
	// 这样可以减少锁的争抢频率，降低“心跳超时”报错的概率。

	for {
		var msg response.SendMessageResponse
		if err := client.Conn.ReadJSON(&msg); err != nil {
			log.Println("[WS] 读取循环退出:", err)
			break
		}

		// 只要有消息进来，就更新读取超时
		_ = client.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		switch msg.Type {
		case "ping":
			// 响应前端文本 Ping：立刻加锁回 Pong
			client.mu.Lock()
			_ = client.Conn.WriteJSON(response.SendMessageResponse{Type: "pong"})
			client.mu.Unlock()

		case "check_unread":
			// 【关键优化】异步执行数据库操作，防止阻塞心跳响应
			go func(uid uint) {
				var count int64
				global.MysqlDB.Model(&models.Message{}).
					Where("receive_user_id = ? AND type = ? AND status = ? AND is_read = ?",
						uid, enum.PrivateMessage, enum.Normal, false).
					Count(&count)

				_ = WsManager.SendToUser(uid, response.SendMessageResponse{
					Type: "has_unread",
					Data: count > 0,
				})
			}(client.UserId)
		}
	}
}
