package memory

type BufferWindowMemory struct {
	BufferMemory
	K int `json:"k"` // number of messages to keep in memory
}

func (m BufferWindowMemory) LoadMemoryVariables() (InputValues, error) {
	start := 0
	if len(m.ChatHistory.messages) > m.K {
		start = len(m.ChatHistory.messages) - m.K
	}
	if m.ReturnMessages {
		return InputValues{
			m.MemoryKey: m.ChatHistory.messages[start:],
		}, nil
	}
	return InputValues{
		m.MemoryKey: getBufferString(m.ChatHistory.messages[start:], m.HumanPrefix, m.AiPrefix),
	}, nil
}

func NewBufferWindowMemory() BufferWindowMemory {
	m := BufferWindowMemory{
		BufferMemory: NewBufferMemory(),
		K:            4,
	}
	return m
}
