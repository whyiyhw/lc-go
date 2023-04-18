package memory

import (
	"reflect"
	"testing"

	"lc_go/schema"
)

func TestChatMessageHistory_AddAiMessage(t *testing.T) {
	type fields struct {
		messages []schema.ChatMessage
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test1",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: "test",
			},
		},
		{
			name: "test2",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: schema.AiChatMessage{
					Text: "test",
				},
			},
		},
		{
			name: "test3",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: schema.SystemChatMessage{
					Text: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ChatMessageHistory{
				messages: tt.fields.messages,
			}
			h.AddAiMessage(tt.args.value)
		})
	}
}

func TestChatMessageHistory_AddUserMessage(t *testing.T) {
	type fields struct {
		messages []schema.ChatMessage
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test1",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: "test",
			},
		},
		{
			name: "test2",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: schema.HumanChatMessage{
					Text: "test",
				},
			},
		},
		{
			name: "test2",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: schema.SystemChatMessage{
					Text: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ChatMessageHistory{
				messages: tt.fields.messages,
			}
			h.AddUserMessage(tt.args.value)
		})
	}
}

func TestNewChatMessageHistory(t *testing.T) {
	tests := []struct {
		name string
		want *ChatMessageHistory
	}{
		{
			name: "test1",
			want: &ChatMessageHistory{
				messages: []schema.ChatMessage{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChatMessageHistory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChatMessageHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatMessageHistory_AddSystemMessage(t *testing.T) {
	type fields struct {
		messages []schema.ChatMessage
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test1",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: "test",
			},
		},
		{
			name: "test2",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: schema.SystemChatMessage{
					Text: "test",
				},
			},
		},
		{
			name: "test2",
			fields: fields{
				messages: []schema.ChatMessage{},
			},
			args: args{
				value: schema.AiChatMessage{
					Text: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ChatMessageHistory{
				messages: tt.fields.messages,
			}
			h.AddSystemMessage(tt.args.value)
		})
	}
}

func TestNewChatMessageHistory1(t *testing.T) {
	type args struct {
		options []NewChatMessageOption
	}
	tests := []struct {
		name string
		args args
		want *ChatMessageHistory
	}{
		{
			name: "test1",
			args: args{
				options: []NewChatMessageOption{
					WithPreviousMessages([]schema.ChatMessage{
						&schema.HumanChatMessage{
							Text: "test",
						},
					}),
				},
			},
			want: &ChatMessageHistory{
				messages: []schema.ChatMessage{
					&schema.HumanChatMessage{
						Text: "test",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChatMessageHistory(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChatMessageHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}
