package directory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateIfPathExists(t *testing.T) {
	dir := "./"
	validation := ValidateIfPathExists(dir)
	assert.True(t, validation)
}

func TestCreateFile(t *testing.T) {
	file := "/tmp/gogpttestcheckfile"
	err := CreateFile(file)
	assert.Empty(t, err)
}

func TestPathWalk(t *testing.T) {
	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				folderPath: "pkg/adapter2/",
			}, wantErr: true,
		},
		{
			args: args{
				folderPath: "pkg/adapter/",
			}, wantErr: false, want: "../../pkg/adapter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathWalk(tt.args.folderPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathWalk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PathWalk() = %v, want %v", got, tt.want)
			}
		})
	}
}
