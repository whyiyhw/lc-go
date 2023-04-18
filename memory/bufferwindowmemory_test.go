package memory

import (
	"reflect"
	"testing"

	"lc_go/schema"
)

func TestBufferWindowMemory_LoadMemoryVariables(t *testing.T) {
	type fields struct {
		BufferMemory BufferMemory
		K            int
	}
	tests := []struct {
		name    string
		fields  fields
		want    InputValues
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				BufferMemory: BufferMemory{
					MemoryKey:      "history",
					ReturnMessages: true,
					ChatHistory: &ChatMessageHistory{
						messages: []schema.ChatMessage{
							&schema.HumanChatMessage{
								Text: "test1",
							},
							&schema.AiChatMessage{
								Text: "test2",
							},
							&schema.HumanChatMessage{
								Text: "test3",
							},
							&schema.AiChatMessage{
								Text: "test4",
							},
							&schema.SystemChatMessage{
								Text: "system",
							},
						},
					},
				},
				K: 3,
			},
			want: InputValues{
				"history": []schema.ChatMessage{
					&schema.HumanChatMessage{
						Text: "test3",
					},
					&schema.AiChatMessage{
						Text: "test4",
					},
					&schema.SystemChatMessage{
						Text: "system",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				BufferMemory: BufferMemory{
					MemoryKey:      "history",
					ReturnMessages: false,
					HumanPrefix:    "Human",
					ChatHistory: &ChatMessageHistory{
						messages: []schema.ChatMessage{
							&schema.HumanChatMessage{
								Text: "test1",
							},
						},
					},
				},
				K: 1,
			},
			want: InputValues{
				"history": "Human: test1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := BufferWindowMemory{
				BufferMemory: tt.fields.BufferMemory,
				K:            tt.fields.K,
			}
			got, err := m.LoadMemoryVariables()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadMemoryVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadMemoryVariables() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBufferWindowMemory(t *testing.T) {
	tests := []struct {
		name string
		want BufferWindowMemory
	}{
		{
			name: "test1",
			want: BufferWindowMemory{
				BufferMemory: BufferMemory{
					MemoryKey:      "history",
					HumanPrefix:    "Human",
					AiPrefix:       "AI",
					ReturnMessages: false,
					ChatHistory:    NewChatMessageHistory(),
				},
				K: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBufferWindowMemory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBufferWindowMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
