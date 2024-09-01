package impl

import (
	"bytes"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/bagusandrian/reconciliation-service/internals/model"
	ucReconciliation "github.com/bagusandrian/reconciliation-service/internals/usecase/reconciliation"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
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
							CSVFilePath: "expect_error.csv",
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

func Test_handler_Reconciliation(t *testing.T) {
	app := fiber.New()
	conf := &model.Config{}
	usecaseReadFileMock := new(ucReconciliation.MockUsecase)
	handlerReconciliation := New(conf, usecaseReadFileMock)
	app.Post("/reconciliation", handlerReconciliation.Reconciliation)
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		data         []byte // data
		expectedCode int    // expected HTTP status code
	}{
		// TODO: Add test cases.
		{
			description:  "POST HTTP status 400 bad request",
			route:        "/reconciliation",
			expectedCode: 400,
		},
		{
			description:  "POST HTTP status 400 missing request (file not found)",
			route:        "/reconciliation",
			expectedCode: 400,
			data:         []byte("{\"system_transaction_csv_file_path\": \"/etc/csv/system.csv\",\"bank_statements\": [{\"bank_name\": \"bca\",\"csv_file_path\": \"/etc/csv/bca.csv\"},{\"bank_name\": \"danamon\",\"csv_file_path\": \"/etc/csv/danamon.csv\"},{\"bank_name\": \"bri\",\"csv_file_path\": \"/etc/csv/bri.csv\"}],\"reconciliaton_start_date\": \"2024-01-20\",\"reconciliaton_end_date\": \"2024-01-22\"}"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(tt.data))
			req.Header.Set("Content-Type", "application/json")
			// Perform the request plain with the app,
			// the second argument is a request latency
			// (set to -1 for no latency)
			resp, _ := app.Test(req, 1)
			if body, err := io.ReadAll(resp.Body); err == nil {
				log.Println(string(body)) // "testing"
			}
			// Verify, if the status code is as expected
			assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)

		})
	}
}
