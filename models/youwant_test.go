package models

import (
	"testing"
)

func TestYouwant_ToJsonString(t *testing.T) {
	type fields struct {
		Label    string
		Detail   string
		Type     string
		Template Template
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "You want to",
			fields: fields{
				Label:  "一个标准 debian 构建系统结构",
				Detail: "基于 x86_64 的 rootfs 构建结构",
				Type:   "linux|debian|build-system",
				Template: Template{
					Action: "shell",
					Shell: Shell{
						Delay: 0,
						Commands: []string{
							"wget http://127.0.0.1/8000/linux/fileuploader/rootfs.tar",
							"sleep 1",
							"echo 1 >> log.txt",
							"echo 2 >> log.txt",
						},
					},
					VSCommands: []string{
						RefreshFilesExplorer,
					},
				},
			},
			want: "Done",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			youwant := &Youwant{
				Label:    tt.fields.Label,
				Detail:   tt.fields.Detail,
				Type:     tt.fields.Type,
				Template: tt.fields.Template,
			}
			// if got := youwant.ToJsonString(); got != tt.want {
			// t.Errorf("Youwant.ToJsonString() = \n%v, want %v", got, tt.want)
			// }

			t.Log(youwant.ToJsonString())
		})
	}
}

func TestNewYouwant(t *testing.T) {
	tests := []struct {
		name string
		want Youwant
	}{
		// TODO: Add test cases.
		{
			name: "you want",
			want: NewYouwant(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.want.ToJsonString())
		})
	}
}
