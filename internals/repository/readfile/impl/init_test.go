package impl

import (
	"testing"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg *model.Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "init impl readfile",
			args: args{
				cfg: &model.Config{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.cfg)
		})
	}
}
