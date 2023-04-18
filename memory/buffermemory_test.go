package memory

import (
	"reflect"
	"testing"

	"lc_go/schema"
)

func TestBufferMemory_SaveContext(t *testing.T) {
	type fields struct {
		ChatHistory    *ChatMessageHistory
		ReturnMessages bool
		InputKey       string
		OutputKey      string
		HumanPrefix    string
		AiPrefix       string
		MemoryKey      string
	}
	type args struct {
		inputValues  InputValues
		outputValues InputValues
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				ChatHistory:    NewChatMessageHistory(),
				ReturnMessages: false,
				InputKey:       "",
				OutputKey:      "",
				HumanPrefix:    "Human",
				AiPrefix:       "AI",
				MemoryKey:      "history",
			},
			args: args{
				inputValues: InputValues{
					"test": "test",
				},
				outputValues: InputValues{
					"test": "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := BufferMemory{
				ChatHistory:    tt.fields.ChatHistory,
				ReturnMessages: tt.fields.ReturnMessages,
				InputKey:       tt.fields.InputKey,
				OutputKey:      tt.fields.OutputKey,
				HumanPrefix:    tt.fields.HumanPrefix,
				AiPrefix:       tt.fields.AiPrefix,
				MemoryKey:      tt.fields.MemoryKey,
			}
			if err := m.SaveContext(tt.args.inputValues, tt.args.outputValues); (err != nil) != tt.wantErr {
				t.Errorf("SaveContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewBufferMemory(t *testing.T) {
	tests := []struct {
		name string
		want BufferMemory
	}{
		{
			name: "test1",
			want: BufferMemory{
				ChatHistory:    NewChatMessageHistory(),
				ReturnMessages: false,
				InputKey:       "",
				OutputKey:      "",
				HumanPrefix:    "Human",
				AiPrefix:       "AI",
				MemoryKey:      "history",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBufferMemory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBufferMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBufferString(t *testing.T) {
	type args struct {
		messages []schema.ChatMessage
		prefix   string
		prefix2  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				messages: []schema.ChatMessage{
					&schema.HumanChatMessage{
						Text: "test",
					},
				},
				prefix:  "Human",
				prefix2: "AI",
			},
			want: "Human: test",
		},
		{
			name: "test2",
			args: args{
				messages: []schema.ChatMessage{
					&schema.AiChatMessage{
						Text: "test",
					},
				},
				prefix:  "Human",
				prefix2: "AI",
			},
			want: "AI: test",
		},
		{
			name: "test2",
			args: args{
				messages: []schema.ChatMessage{
					&schema.SystemChatMessage{
						Text: "test",
					},
				},
				prefix:  "Human",
				prefix2: "AI",
			},
			want: "system: test",
		},
		{
			name: "test2",
			args: args{
				messages: []schema.ChatMessage{
					nil,
				},
				prefix:  "Human",
				prefix2: "AI",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBufferString(tt.args.messages, tt.args.prefix, tt.args.prefix2); got != tt.want {
				t.Errorf("getBufferString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBufferMemory_LoadMemoryVariables(t *testing.T) {
	type fields struct {
		ChatHistory    *ChatMessageHistory
		ReturnMessages bool
		InputKey       string
		OutputKey      string
		HumanPrefix    string
		AiPrefix       string
		MemoryKey      string
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
				ChatHistory:    NewChatMessageHistory(),
				ReturnMessages: false,
				InputKey:       "",
				OutputKey:      "",
				HumanPrefix:    "Human",
				AiPrefix:       "AI",
				MemoryKey:      "history",
			},
			want: InputValues{
				"history": "",
			},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				ChatHistory:    NewChatMessageHistory(),
				ReturnMessages: true,
				InputKey:       "input",
				OutputKey:      "output",
				HumanPrefix:    "Human",
				AiPrefix:       "AI",
				MemoryKey:      "history",
			},
			want: InputValues{
				"history": []schema.ChatMessage{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := BufferMemory{
				ChatHistory:    tt.fields.ChatHistory,
				ReturnMessages: tt.fields.ReturnMessages,
				InputKey:       tt.fields.InputKey,
				OutputKey:      tt.fields.OutputKey,
				HumanPrefix:    tt.fields.HumanPrefix,
				AiPrefix:       tt.fields.AiPrefix,
				MemoryKey:      tt.fields.MemoryKey,
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
