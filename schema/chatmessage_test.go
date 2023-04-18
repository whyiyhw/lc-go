package schema

import "testing"

func TestAiChatMessage_GetRole(t *testing.T) {
	type fields struct {
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				Text: "AI",
			},
			want: "assistant",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &AiChatMessage{
				Text: tt.fields.Text,
			}
			if got := h.GetRole(); got != tt.want {
				t.Errorf("GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAiChatMessage_GetText(t *testing.T) {
	type fields struct {
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				Text: "AI",
			},
			want: "AI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &AiChatMessage{
				Text: tt.fields.Text,
			}
			if got := h.GetText(); got != tt.want {
				t.Errorf("GetText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHumanChatMessage_GetRole(t *testing.T) {
	type fields struct {
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				Text: "Human",
			},
			want: "user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HumanChatMessage{
				Text: tt.fields.Text,
			}
			if got := h.GetRole(); got != tt.want {
				t.Errorf("GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHumanChatMessage_GetText(t *testing.T) {
	type fields struct {
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				Text: "Human",
			},
			want: "Human",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HumanChatMessage{
				Text: tt.fields.Text,
			}
			if got := h.GetText(); got != tt.want {
				t.Errorf("GetText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemChatMessage_GetRole(t *testing.T) {
	type fields struct {
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				Text: "System",
			},
			want: "system",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SystemChatMessage{
				Text: tt.fields.Text,
			}
			if got := m.GetRole(); got != tt.want {
				t.Errorf("GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemChatMessage_GetText(t *testing.T) {
	type fields struct {
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				Text: "System",
			},
			want: "System",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SystemChatMessage{
				Text: tt.fields.Text,
			}
			if got := m.GetText(); got != tt.want {
				t.Errorf("GetText() = %v, want %v", got, tt.want)
			}
		})
	}
}
