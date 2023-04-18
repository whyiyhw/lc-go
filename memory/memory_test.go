package memory

import (
	"reflect"
	"testing"
)

func Test_getInputValue(t *testing.T) {
	type args struct {
		values InputValues
		key    string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				values: InputValues{
					"test": "test",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				values: InputValues{
					"test": "test",
				},
				key: "test",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				values: InputValues{
					"test": "test",
				},
				key: "test1",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getInputValue(tt.args.values, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getInputValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInputValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
