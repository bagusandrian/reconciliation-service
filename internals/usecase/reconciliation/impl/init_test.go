package impl

import (
	"testing"

	"github.com/bagusandrian/reconciliation-service/internals/model"
	"github.com/bagusandrian/reconciliation-service/internals/repository/readfile"
)

func TestNew(t *testing.T) {
	readFileMock := new(readfile.MockReadFile)
	type args struct {
		cfg      *model.Config
		readFile readfile.ReadFile
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success init usecase reconciliation",
			args: args{
				cfg:      &model.Config{},
				readFile: readFileMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.cfg, tt.args.readFile)
		})
	}
}
