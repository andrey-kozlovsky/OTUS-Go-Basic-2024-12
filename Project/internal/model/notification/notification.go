package notification

import (
	"time"
)

// Notification представляет сущность уведомления.
type Notification struct {
	id      int
	Message string
	SentAt  time.Time
	isRead  bool
}

// NewNotification создает новое уведомление.
func NewNotification(id int, message string) *Notification {
	return &Notification{
		id:      id,
		Message: message,
		SentAt:  time.Now(),
		isRead:  false,
	}
}

// GetID возвращает ID уведомления.
func (n *Notification) GetID() int {
	return n.id
}

// MarkAsRead отмечает уведомление как прочитанное.
func (n *Notification) MarkAsRead() {
	n.isRead = true
}

// IsRead возвращает статус уведомления (прочитано/не прочитано).
func (n *Notification) IsRead() bool {
	return n.isRead
}
