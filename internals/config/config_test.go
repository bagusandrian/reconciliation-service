package config

import (
	"testing"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

func TestNew(t *testing.T) {
	type args struct {
		repoName string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success create new config",
			args: args{
				repoName: "reconciliation-service",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.repoName)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
