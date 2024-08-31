package impl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

func Test_validationReconciliationRequest(t *testing.T) {
	repoName := "reconciliation-service"
	repoPath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/bagusandrian", repoName)
	type args struct {
		req *model.ReconciliationRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid request",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: filepath.Join(repoPath, "files/csv", "system.csv"),
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "bank_1",
							CSVFilePath: filepath.Join(repoPath, "files/csv", "bca.csv"),
						},
					},
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "2024-01-02",
				},
			},
		},
		{
			name: "error parsing ReconciliationStartDate request",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: filepath.Join(repoPath, "files/csv", "system.csv"),
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "bank_1",
							CSVFilePath: filepath.Join(repoPath, "files/csv", "bca.csv"),
						},
					},
					ReconciliationStartDateString: "expect error",
					ReconciliationEndDateString:   "2024-01-02",
				},
			},
			wantErr: true,
		},
		{
			name: "error parsing ReconciliationEndDateString request",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: filepath.Join(repoPath, "files/csv", "system.csv"),
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "bank_1",
							CSVFilePath: filepath.Join(repoPath, "files/csv", "bca.csv"),
						},
					},
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "expect error",
				},
			},
			wantErr: true,
		},
		{
			name: "error extention csv system",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: filepath.Join(repoPath, "files/csv", "system.mpv"),
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "bank_1",
							CSVFilePath: filepath.Join(repoPath, "files/csv", "bca.csv"),
						},
					},
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "2024-01-02",
				},
			},
			wantErr: true,
		},
		{
			name: "error file path system",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: "expect error",
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "bank_1",
							CSVFilePath: filepath.Join(repoPath, "files/csv", "bca.csv"),
						},
					},
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "2024-01-02",
				},
			},
			wantErr: true,
		},
		{
			name: "error bank ext",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: filepath.Join(repoPath, "files/csv", "system.csv"),
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "bank_1",
							CSVFilePath: filepath.Join(repoPath, "files/csv", "bca.mp3"),
						},
					},
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "2024-01-02",
				},
			},
			wantErr: true,
		},
		{
			name: "error file path bank",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: filepath.Join(repoPath, "files/csv", "system.csv"),
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "bank_1",
							CSVFilePath: "expect error",
						},
					},
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "2024-01-02",
				},
			},
			wantErr: true,
		},
		{
			name: "end date lower than start date",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: filepath.Join(repoPath, "files/csv", "system.csv"),
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "bank_1",
							CSVFilePath: filepath.Join(repoPath, "files/csv", "bca.csv"),
						},
					},
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "2023-01-02",
				},
			},
			wantErr: true,
		},
		{
			name: "data bank empty",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath:  filepath.Join(repoPath, "files/csv", "system.csv"),
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "2024-01-02",
				},
			},
			wantErr: true,
		},
		{
			name: "bank name empty",
			args: args{
				req: &model.ReconciliationRequest{
					SystemTransactionCSVFilePath: filepath.Join(repoPath, "files/csv", "system.csv"),
					BankStatements: []model.BankCSCVFilePath{
						model.BankCSCVFilePath{
							BankName:    "",
							CSVFilePath: filepath.Join(repoPath, "files/csv", "bca.csv"),
						},
					},
					ReconciliationStartDateString: "2024-01-01",
					ReconciliationEndDateString:   "2024-01-02",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validationReconciliationRequest(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validationReconciliationRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
