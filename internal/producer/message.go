package producer

type MessageType uint

const (
	Create MessageType = iota
	Update
	Remove
)

type Message struct {
	Type string `json:"type"`
	ID   uint64 `json:"id"`
}

// String реализация метода MessageType
func (msgType MessageType) String() string {
	switch msgType {
	case Create:
		return "Create"
	case Update:
		return "Update"
	case Remove:
		return "Remove"
	default:
		return "Unknown MessageType"
	}
}

// NewMessage создаёт новое сообщение для брокера
func NewMessage(suggestionID uint64, msgType MessageType) *Message {
	return &Message{
		Type: msgType.String(),
		ID:   suggestionID,
	}
}
