package hello

// Message type
type Message struct {
	msg string
}

// CreateMessage construct a Message
func CreateMessage(msg string) Message {
	return Message{msg: msg}
}

// ToString format a Message to string
func (m *Message) ToString() string {
	return m.msg
}

// Append text to a Message
func (m *Message) Append(text string) {
	m.msg += text
}
