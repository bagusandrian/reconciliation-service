package impl

import (
	"reflect"
	"testing"
)

func Test_repoReadFile_openFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name        string
		args        args
		wantHeader  []string
		wantRecords [][]string
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name: "failed read file",
			args: args{
				filepath: ".wrong.csv",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotRecords, err := openFileCSV(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("repoReadFile.openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("repoReadFile.openFile() gotHeader = %v, want %v", gotHeader, tt.wantHeader)
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("repoReadFile.openFile() gotRecords = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}
