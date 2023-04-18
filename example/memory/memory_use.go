package main

import (
	"fmt"

	"lc_go/memory"
	"lc_go/schema"
)

func main() {
	// use buffer memory
	m := memory.NewBufferMemory()
	m.InputKey = m.HumanPrefix
	m.OutputKey = m.AiPrefix
	_ = m.SaveContext(memory.InputValues{m.HumanPrefix: "TXX"}, memory.InputValues{m.AiPrefix: "ok"})

	m.ReturnMessages = false
	data, _ := m.LoadMemoryVariables()
	fmt.Println(data[m.MemoryKey])
	fmt.Println("------------------------")

	m.ReturnMessages = true
	res, _ := m.LoadMemoryVariables()
	dto := res[m.MemoryKey]
	if rec, ok := dto.([]schema.ChatMessage); ok {
		for _, i2 := range rec {
			fmt.Println(i2.GetRole(), i2.GetText())
		}
	}

	// use buffer window memory
	bwm := memory.NewBufferWindowMemory()
	bwm.InputKey = m.HumanPrefix
	bwm.OutputKey = m.AiPrefix
	bwm.K = 2
	_ = bwm.SaveContext(memory.InputValues{bwm.HumanPrefix: "1"}, memory.InputValues{bwm.AiPrefix: "ok"})
	_ = bwm.SaveContext(memory.InputValues{bwm.HumanPrefix: "2"}, memory.InputValues{bwm.AiPrefix: "ok"})

	bwm.ReturnMessages = false
	data, _ = bwm.LoadMemoryVariables()
	fmt.Println(data[bwm.MemoryKey])
	fmt.Println("------------------------")

	bwm.ReturnMessages = true
	res, _ = bwm.LoadMemoryVariables()
	dto = res[bwm.MemoryKey]
	if rec, ok := dto.([]schema.ChatMessage); ok {
		for _, i2 := range rec {
			fmt.Println(i2.GetRole(), i2.GetText())
		}
	}
}
