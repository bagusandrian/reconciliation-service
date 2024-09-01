package impl

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/bagusandrian/reconciliation-service/internals/model"
	"github.com/bagusandrian/reconciliation-service/internals/repository/readfile"
)

func Test_usecase_ReconciliationComparition(t *testing.T) {
	rReadFileMock := new(readfile.MockReadFile)
	reqFailedFileSystem := model.ReconciliationRequest{
		SystemTransactionCSVFilePath: "testing",
	}
	respSuccessGetFile := model.ReconciliationResponse{}
	respSuccessGetFile.DetailOfMatchedTransactions = make(map[string]model.DetailMatchedTransaction)
	respSuccessGetFile.DetailOfUnmatchedTransactions = make(map[string]model.DetailUnmatchedTransaction)
	type args struct {
		ctx context.Context
		req model.ReconciliationRequest
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantResp model.ReconciliationResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "failed get system file",
			args: args{
				ctx: context.Background(),
				req: reqFailedFileSystem,
			},
			mockFunc: func() {
				rReadFileMock.On("GetSystemReconciliationCSV", reqFailedFileSystem).
					Return(model.DataSystem{}, errors.New("expect error")).Once()
			},
			wantErr: true,
		},
		{
			name: "failed get bank file",
			args: args{
				ctx: context.Background(),
				req: reqFailedFileSystem,
			},
			mockFunc: func() {
				rReadFileMock.On("GetSystemReconciliationCSV", reqFailedFileSystem).
					Return(model.DataSystem{}, nil).Once()
				rReadFileMock.On("GetBankReconciliationCSV", reqFailedFileSystem).
					Return(model.DataBank{}, errors.New("expect error")).Once()
			},
			wantErr: true,
		},
		{
			name: "success get file",
			args: args{
				ctx: context.Background(),
				req: model.ReconciliationRequest{},
			},
			mockFunc: func() {
				rReadFileMock.On("GetSystemReconciliationCSV", model.ReconciliationRequest{}).
					Return(model.DataSystem{}, nil).Once()
				rReadFileMock.On("GetBankReconciliationCSV", model.ReconciliationRequest{}).
					Return(model.DataBank{}, nil).Once()
			},
			wantResp: respSuccessGetFile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			u := &usecase{
				cfg:      &model.Config{},
				readFile: rReadFileMock,
			}
			gotResp, err := u.ReconciliationComparition(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.ReconciliationComparition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("usecase.ReconciliationComparition() = %+v, want %+v", gotResp, tt.wantResp)
			}
		})
	}
}
