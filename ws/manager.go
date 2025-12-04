package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserId uint
	Conn   *websocket.Conn
}

type Manager struct {
	clients map[uint]*Client
	lock    sync.RWMutex
}

var WsManager = &Manager{
	clients: make(map[uint]*Client),
}

// AddClient 添加客户端
func (m *Manager) AddClient(client *Client) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.clients[client.UserId] = client
}

// RemoveClient 删除客户端
func (m *Manager) RemoveClient(userId uint) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.clients, userId)
}

// GetClient 获取客户端
func (m *Manager) GetClient(userId uint) *Client {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.clients[userId]
}

// SendToUser 推送消息方法
func (m *Manager) SendToUser(userId uint, data interface{}) error {
	m.lock.RLock()
	client, ok := m.clients[userId]
	m.lock.RUnlock()

	if !ok {
		return nil // 用户不在线
	}
	return client.Conn.WriteJSON(data)
}
