package impl

import (
	"context"
	"encoding/json"
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

func Test_usecase_reconciliationData(t *testing.T) {
	rReadFileMock := new(readfile.MockReadFile)
	conf := &model.Config{}
	// init data bank case success
	BankTransactionSuccess := model.DataBank{}
	dataBankCSVSuccess := make(map[string][]model.DataBankCSV)
	dataBankCSVSuccess["bank_testing"] = []model.DataBankCSV{
		model.DataBankCSV{
			UniqueIdentifier: "bank-testing-001",
			Amount:           10000,
			DateString:       "2024-01-01",
			Type:             model.Debit,
		},
		model.DataBankCSV{
			UniqueIdentifier: "bank-testing-002",
			Amount:           -10000,
			DateString:       "2024-01-01",
			Type:             model.Credit,
		},
	}
	BankTransactionSuccess.TotalData = 2
	BankTransactionSuccess.DataBankCSV = dataBankCSVSuccess

	// init data bank case not found several
	BankTransactionNotFoundSeveral := model.DataBank{}
	dataBankCSVNotFoundSeveral := make(map[string][]model.DataBankCSV)
	dataBankCSVNotFoundSeveral["bank_testing"] = []model.DataBankCSV{
		model.DataBankCSV{
			UniqueIdentifier: "bank-testing-001",
			Amount:           10000,
			DateString:       "2024-01-01",
			Type:             model.Debit,
		},
		model.DataBankCSV{
			UniqueIdentifier: "bank-testing-002",
			Amount:           -10000,
			DateString:       "2024-01-01",
			Type:             model.Credit,
		},
		model.DataBankCSV{
			UniqueIdentifier: "bank-testing-003",
			Amount:           -22222,
			DateString:       "2024-01-01",
			Type:             model.Credit,
		},
	}
	BankTransactionNotFoundSeveral.TotalData = 2
	BankTransactionNotFoundSeveral.DataBankCSV = dataBankCSVNotFoundSeveral

	// init detail match tansaction response case success
	detailMatchTransactionsSuccess := make(map[string]model.DetailMatchedTransaction)
	detailMatchTransactionsSuccess["bank_testing"] = model.DetailMatchedTransaction{
		TotalNumberMatchedTransactions: 2,
		DetailTransaction: []model.DetailTransaction{
			model.DetailTransaction{
				TrxID:            "trx001",
				UniqueIdentifier: "bank-testing-001",
				Amount:           10000,
				Date:             "2024-01-01",
				TransactionTime:  "2024-01-01 10:00:00",
			},
			model.DetailTransaction{
				TrxID:            "trx002",
				UniqueIdentifier: "bank-testing-002",
				Amount:           -10000,
				Date:             "2024-01-01",
				TransactionTime:  "2024-01-01 10:00:01",
			},
		},
	}
	// init detail match tansaction response case not found several
	detailMatchTransactionsNotFoundSeveral := make(map[string]model.DetailMatchedTransaction)
	detailMatchTransactionsNotFoundSeveral["bank_testing"] = model.DetailMatchedTransaction{
		TotalNumberMatchedTransactions: 2,
		DetailTransaction: []model.DetailTransaction{
			model.DetailTransaction{
				TrxID:            "trx001",
				UniqueIdentifier: "bank-testing-001",
				Amount:           10000,
				Date:             "2024-01-01",
				TransactionTime:  "2024-01-01 10:00:00",
			},
			model.DetailTransaction{
				TrxID:            "trx002",
				UniqueIdentifier: "bank-testing-002",
				Amount:           -10000,
				Date:             "2024-01-01",
				TransactionTime:  "2024-01-01 10:00:01",
			},
		},
	}

	// init response success case
	detailOfUnmatchedTransactionsSuccess := make(map[string]model.DetailUnmatchedTransaction)
	wantResponseSuccess := model.ReconciliationResponse{
		TotalTranscationsProcessed:       4,
		TotalNumberMatchedTransactions:   4,
		DetailOfMatchedTransactions:      detailMatchTransactionsSuccess,
		TotalNumberUnmatchedTransactions: 0,
		DetailOfUnmatchedTransactions:    detailOfUnmatchedTransactionsSuccess,
		TotalDiscrepanciesAmount:         0,
	}

	// init response notfound severeal case
	detailOfUnmatchedTransactionsNotFoundSeveral := make(map[string]model.DetailUnmatchedTransaction)
	detailOfUnmatchedTransactionsNotFoundSeveral["bank_testing"] = model.DetailUnmatchedTransaction{
		Info: "bank statement not found on any system transaction",
		DetailTransaction: []model.DetailTransaction{
			model.DetailTransaction{
				UniqueIdentifier: "bank-testing-003",
				Amount:           -22222,
				Date:             "2024-01-01",
				Type:             model.Debit,
			},
		},
	}
	detailOfUnmatchedTransactionsNotFoundSeveral["system"] = model.DetailUnmatchedTransaction{
		Info: "system transaction not found on any bank statement",
		DetailTransaction: []model.DetailTransaction{
			model.DetailTransaction{
				TrxID:           "trx003",
				Amount:          9999,
				TransactionTime: "2024-01-01 10:00:01",
				Type:            model.Debit,
			},
		},
	}

	wantResponseNotFoundSeveral := model.ReconciliationResponse{
		TotalTranscationsProcessed:       4,
		TotalNumberMatchedTransactions:   4,
		DetailOfMatchedTransactions:      detailMatchTransactionsNotFoundSeveral,
		TotalNumberUnmatchedTransactions: 2,
		DetailOfUnmatchedTransactions:    detailOfUnmatchedTransactionsNotFoundSeveral,
		TotalDiscrepanciesAmount:         0,
	}
	type args struct {
		systemTransaction *model.DataSystem
		bankTransaction   *model.DataBank
	}
	tests := []struct {
		name string
		args args
		want model.ReconciliationResponse
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				systemTransaction: &model.DataSystem{
					DataSystemCSV: []model.DataSystemCSV{
						model.DataSystemCSV{
							TrxID:                 "trx001",
							Amount:                10000,
							Type:                  model.Debit,
							TransactionTimeString: "2024-01-01 10:00:00",
						},
						model.DataSystemCSV{
							TrxID:                 "trx002",
							Amount:                -10000,
							Type:                  model.Credit,
							TransactionTimeString: "2024-01-01 10:00:01",
						},
					},
					TotalData: 2,
				},
				bankTransaction: &BankTransactionSuccess,
			},
			want: wantResponseSuccess,
		}, {
			name: "some transaction not match",
			args: args{
				systemTransaction: &model.DataSystem{
					DataSystemCSV: []model.DataSystemCSV{
						model.DataSystemCSV{
							TrxID:                 "trx001",
							Amount:                10000,
							Type:                  model.Debit,
							TransactionTimeString: "2024-01-01 10:00:00",
						},
						model.DataSystemCSV{
							TrxID:                 "trx002",
							Amount:                -10000,
							Type:                  model.Credit,
							TransactionTimeString: "2024-01-01 10:00:01",
						},
						model.DataSystemCSV{
							TrxID:                 "trx003",
							Amount:                9999,
							Type:                  model.Credit,
							TransactionTimeString: "2024-01-01 10:00:01",
						},
					},
					TotalData: 2,
				},
				bankTransaction: &BankTransactionNotFoundSeveral,
			},
			want: wantResponseNotFoundSeveral,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				cfg:      conf,
				readFile: rReadFileMock,
			}
			if got := u.reconciliationData(tt.args.systemTransaction, tt.args.bankTransaction); !reflect.DeepEqual(got, tt.want) {
				resp, _ := json.Marshal(got)
				expect, _ := json.Marshal(tt.want)
				t.Errorf("usecase.getMatchData() = %s, \n\nwant %s", resp, expect)
			}
		})
	}
}
