package schema

type ChatMessage interface {
	GetText() string
	GetRole() string
}

// HumanChatMessage is a message from a human
type HumanChatMessage struct {
	Text string
}

func (h *HumanChatMessage) GetText() string { return h.Text }
func (h *HumanChatMessage) GetRole() string { return "user" }

// AiChatMessage is a message from an AI
type AiChatMessage struct {
	Text string
}

func (h *AiChatMessage) GetText() string { return h.Text }
func (h *AiChatMessage) GetRole() string { return "assistant" }

// SystemChatMessage is a message from the system
type SystemChatMessage struct {
	Text string
}

func (m *SystemChatMessage) GetText() string { return m.Text }
func (m *SystemChatMessage) GetRole() string { return "system" }
