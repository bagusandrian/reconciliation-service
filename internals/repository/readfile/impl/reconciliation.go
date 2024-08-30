package impl

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	cons "github.com/bagusandrian/reconciliation-service/internals/constant"
	"github.com/bagusandrian/reconciliation-service/internals/model"
)

func (r *repoReadFile) GetSystemReconciliationCSV(req model.ReconciliationRequest) (resp []model.DataSystemCSV, err error) {
	// open file
	headerCSV, records, err := r.openFile(req.SystemTransactionCSVFilePath)
	if err != nil {
		return resp, err
	}
	// validate header
	if headerValid := reflect.DeepEqual(headerCSV, cons.ValidateHeaderSystemCSV); !headerValid {
		return resp, fmt.Errorf("file system reconciliation csv header is invalid: %s, expectation: %s",
			strings.Join(headerCSV, ","), strings.Join(cons.ValidateHeaderSystemCSV, ","))
	}
	countRow := 1
	DataSystemCSV := []model.DataSystemCSV{}
	// Loop and storing data
	for _, eachrecord := range records {
		// validate every rows have 4 field
		if len(eachrecord) != cons.LenRowSystem {
			return resp, fmt.Errorf("row %d is invalid data", countRow)
		}
		// validate amount
		amount, err := strconv.ParseFloat(eachrecord[1], 64)
		if err != nil {
			return resp, fmt.Errorf("row %d amount is not valid %s err:%+v", countRow, eachrecord[1], err)
		}
		// validate type
		typeTransaction, err := strconv.Atoi(eachrecord[2])
		if err != nil {
			return resp, fmt.Errorf("row %d type is not valid %s err:%+v", countRow, eachrecord[2], err)
		}
		transactionTime, err := time.Parse("2006-01-02 15:04:05", eachrecord[3])
		if err != nil {
			return resp, fmt.Errorf("row %d transaction_time is not valid %s err:%+v", countRow, eachrecord[3], err)
		}
		DataSystemCSV = append(DataSystemCSV, model.DataSystemCSV{
			TrxID:                 eachrecord[0],
			Amount:                amount,
			Type:                  model.TypeTransaction(typeTransaction),
			TransactionTimeString: eachrecord[3],
			TransactionTime:       transactionTime,
		})
		countRow++
	}
	return DataSystemCSV, err
}
func (r *repoReadFile) GetBankReconciliationCSV(req model.ReconciliationRequest) (resp map[string][]model.DataBankCSV, err error) {
	resp = make(map[string][]model.DataBankCSV)
	for _, v := range req.BankStatements {
		resp[v.BankName] = []model.DataBankCSV{}
		// open file
		headerCSV, records, err := r.openFile(v.CSVFilePath)
		if err != nil {
			return resp, err
		}
		if headerValid := reflect.DeepEqual(headerCSV, cons.ValidateHeaderBankCSV); !headerValid {
			return resp, fmt.Errorf("file bank %s reconciliation csv header is invalid: %s, expectation: %s",
				v.BankName, strings.Join(headerCSV, ","), strings.Join(cons.ValidateHeaderBankCSV, ","))
		}
		countRow := 1
		DataBankCSV := []model.DataBankCSV{}
		// Loop and storing data
		for _, eachrecord := range records {
			// validate every rows have 3 field
			if len(eachrecord) != cons.LenRowBank {
				return resp, fmt.Errorf("bank %s row %d is invalid data", v.BankName, countRow)
			}
			// validate amount
			amount, err := strconv.ParseFloat(eachrecord[1], 64)
			if err != nil {
				return resp, fmt.Errorf("bank %s row %d amount is not valid %s err:%+v", v.BankName, countRow, eachrecord[1], err)
			}

			transactionDate, err := time.Parse("2006-01-02", eachrecord[2])
			if err != nil {
				return resp, fmt.Errorf("row %d transaction_time is not valid %s err:%+v", countRow, eachrecord[3], err)
			}
			DataBankCSV = append(DataBankCSV, model.DataBankCSV{
				UniqueIdentifier: eachrecord[0],
				Amount:           amount,
				DateString:       eachrecord[2],
				Date:             transactionDate,
			})
			countRow++
		}
		resp[v.BankName] = DataBankCSV

	}
	return resp, nil
}

func (r *repoReadFile) openFile(filepath string) (header []string, records [][]string, err error) {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}

	// Closes the file
	defer file.Close()

	reader := csv.NewReader(file)
	// get header csv
	header, err = reader.Read()
	if err != nil {
		return nil, nil, err
	}
	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err = reader.ReadAll()
	// Checks for the error
	if err != nil {
		return nil, nil, err
	}
	return header, records, nil
}
