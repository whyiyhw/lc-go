package memory

import (
	"fmt"
	"strings"

	"lc_go/schema"
)

type BufferMemory struct {
	ChatHistory    *ChatMessageHistory // 保存的聊天记录
	ReturnMessages bool                // 是否返回聊天记录
	InputKey       string              // 输入的键
	OutputKey      string              // 输出的键
	HumanPrefix    string              // 人类的前缀
	AiPrefix       string              // AI 的前缀
	MemoryKey      string              // 保存聊天记录的键
}

// NewBufferMemory 创建一个 BufferMemory
func NewBufferMemory() BufferMemory {
	m := BufferMemory{
		ChatHistory:    NewChatMessageHistory(),
		ReturnMessages: false,
		InputKey:       "",
		OutputKey:      "",
		HumanPrefix:    "Human",
		AiPrefix:       "AI",
		MemoryKey:      "history",
	}

	return m
}

// SaveContext 保存聊天记录
func (m BufferMemory) SaveContext(inputValues InputValues, outputValues InputValues) error {
	userInputValue, err := getInputValue(inputValues, m.InputKey)
	if err != nil {
		return err
	}

	m.ChatHistory.AddUserMessage(userInputValue)

	aiOutputValue, err := getInputValue(outputValues, m.OutputKey)
	if err != nil {
		return nil
	}

	m.ChatHistory.AddAiMessage(aiOutputValue)

	return nil
}

// LoadMemoryVariables 加载聊天记录
func (m BufferMemory) LoadMemoryVariables() (InputValues, error) {
	if m.ReturnMessages {
		return InputValues{
			m.MemoryKey: m.ChatHistory.messages,
		}, nil
	}

	return InputValues{
		m.MemoryKey: getBufferString(m.ChatHistory.messages, m.HumanPrefix, m.AiPrefix),
	}, nil
}

func getBufferString(messages []schema.ChatMessage, humanPrefix string, aiPrefix string) string {
	var buffer []string
	for _, message := range messages {
		switch message.(type) {
		case *schema.HumanChatMessage:
			buffer = append(buffer, humanPrefix+": "+message.GetText())
		case *schema.AiChatMessage:
			buffer = append(buffer, aiPrefix+": "+message.GetText())
		case *schema.SystemChatMessage:
			buffer = append(buffer, "system: "+message.GetText())
		default:
			fmt.Println("Unknown message type")
		}
	}
	return strings.Join(buffer, "\n")
}
