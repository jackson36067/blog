package ws

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserId uint
	Conn   *websocket.Conn
	close  sync.Once
	mu     sync.Mutex // 保证只有一个客户端在写数据
}

func (c *Client) Close() {
	c.close.Do(func() {
		_ = c.Conn.Close()
	})
}

type Manager struct {
	clients map[uint]*Client
	lock    sync.RWMutex
}

var WsManager = &Manager{
	clients: make(map[uint]*Client),
}

func (m *Manager) AddClient(client *Client) {
	m.lock.Lock()
	defer m.lock.Unlock()
	// 1. 检查是否存在旧连接
	if oldClient, ok := m.clients[client.UserId]; ok {
		_ = oldClient.Conn.Close() // 强制关闭旧连接
	}
	// 2. 注册新连接
	m.clients[client.UserId] = client
}

func (m *Manager) RemoveClient(userId uint, conn *websocket.Conn) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if client, ok := m.clients[userId]; ok {
		if client.Conn == conn {
			delete(m.clients, userId)
			log.Printf("用户 %d 连接已移除", userId)
		}
	}
}

func (m *Manager) GetClient(userId uint) *Client {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.clients[userId]
}

func (m *Manager) SendToUser(userId uint, data any) error {
	client := m.GetClient(userId)
	if client == nil {
		return nil
	}
	client.mu.Lock()
	defer client.mu.Unlock()
	_ = client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return client.Conn.WriteJSON(data)
}
