package memory

import (
	"fmt"

	"lc_go/schema"
)

type ChatMessageHistory struct {
	messages []schema.ChatMessage
}

func NewChatMessageHistory(options ...NewChatMessageOption) *ChatMessageHistory {
	h := &ChatMessageHistory{
		messages: make([]schema.ChatMessage, 0),
	}
	for _, option := range options {
		option(h)
	}
	return h
}

func (h *ChatMessageHistory) AddUserMessage(value any) {
	switch value.(type) {
	case string:
		h.messages = append(h.messages, &schema.HumanChatMessage{Text: value.(string)})
	case schema.HumanChatMessage:
		h.messages = append(h.messages, &schema.HumanChatMessage{Text: value.(schema.HumanChatMessage).Text})
	default:
		fmt.Printf("Input values to buffer memory must be string. Got type %T. Input values: %v. Memory input key: %s\n", value, h.messages, value)
	}
}

func (h *ChatMessageHistory) AddAiMessage(value any) {
	switch value.(type) {
	case string:
		h.messages = append(h.messages, &schema.AiChatMessage{Text: value.(string)})
	case schema.AiChatMessage:
		h.messages = append(h.messages, &schema.AiChatMessage{Text: value.(schema.AiChatMessage).Text})
	default:
		fmt.Printf("Input values to buffer memory must be string. Got type %T. Input values: %v. Memory input key: %s\n", value, h.messages, value)
	}
}

func (h *ChatMessageHistory) AddSystemMessage(value any) {
	switch value.(type) {
	case string:
		h.messages = append(h.messages, &schema.SystemChatMessage{Text: value.(string)})
	case schema.SystemChatMessage:
		h.messages = append(h.messages, &schema.SystemChatMessage{Text: value.(schema.SystemChatMessage).Text})
	default:
		fmt.Printf("Input values to buffer memory must be string. Got type %T. Input values: %v. Memory input key: %s\n", value, h.messages, value)
	}
}

type NewChatMessageOption func(m *ChatMessageHistory)

// WithPreviousMessages adds previous messages to the history
func WithPreviousMessages(previousMessages []schema.ChatMessage) NewChatMessageOption {
	return func(m *ChatMessageHistory) {
		m.messages = append(m.messages, previousMessages...)
	}
}
