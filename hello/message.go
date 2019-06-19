package hello

// Message type
type Message struct {
	msg string
}

// NewMessage construct a Message
func NewMessage(msg string) Message {
	return Message{msg: msg}
}

// Append text to a Message
func (m *Message) Append(text string) {
	m.msg += text
}

// String for Message
func (m *Message) String() string {
	return m.msg
}
