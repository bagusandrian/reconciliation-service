package impl

import (
	"context"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

func (u *usecase) ReconciliationComparition(ctx context.Context, req model.ReconciliationRequest) (resp model.ReconciliationResponse, err error) {
	// get csv file system
	systemTransaction, err := u.readFile.GetSystemReconciliationCSV(req)
	if err != nil {
		return resp, err
	}
	bankTransaction, err := u.readFile.GetBankReconciliationCSV(req)
	if err != nil {
		return resp, err
	}
	// processing of matched transactions
	resp = u.reconciliationData(&systemTransaction, &bankTransaction)
	// processing of unmatched transactions
	return resp, nil
}

func (u *usecase) reconciliationData(systemTransaction *model.DataSystem, bankTransaction *model.DataBank) model.ReconciliationResponse {
	var (
		result                   model.ReconciliationResponse
		unMatchTransactionSystem []model.DetailTransaction
		totalUnmatchSystem       int64
		unMatchTransactionsBanks map[string]map[string]model.DetailTransaction
		// match  model.DetailMatchedTransaction
	)
	unMatchTransactionsBanks = make(map[string]map[string]model.DetailTransaction)
	result.DetailOfMatchedTransactions = make(map[string]model.DetailMatchedTransaction)
	for i := 0; i < len(systemTransaction.DataSystemCSV); i++ {
		if systemTransaction.DataSystemCSV[i].MatchTransaction {
			continue
		}
		for bankName := range bankTransaction.DataBankCSV {
			if len(bankTransaction.DataBankCSV[bankName]) == 0 {
				continue
			}
			if _, ok := unMatchTransactionsBanks[bankName]; !ok {
				unMatchTransactionsBanks[bankName] = make(map[string]model.DetailTransaction)
			}
			for j := 0; j < len(bankTransaction.DataBankCSV[bankName]); j++ {
				if systemTransaction.DataSystemCSV[i].MatchTransaction {
					break
				}
				if bankTransaction.DataBankCSV[bankName][j].MatchTransaction {
					continue
				}
				if systemTransaction.DataSystemCSV[i].TransactionTime.Format("2006-01-02") == bankTransaction.DataBankCSV[bankName][j].Date.Format("2006-01-02") &&
					systemTransaction.DataSystemCSV[i].Amount == bankTransaction.DataBankCSV[bankName][j].Amount &&
					systemTransaction.DataSystemCSV[i].Type == bankTransaction.DataBankCSV[bankName][j].Type {
					// transaction match
					systemTransaction.DataSystemCSV[i].MatchTransaction = true
					systemTransaction.DataSystemCSV[i].BankID = bankTransaction.DataBankCSV[bankName][j].UniqueIdentifier
					bankTransaction.DataBankCSV[bankName][j].MatchTransaction = true
					bankTransaction.DataBankCSV[bankName][j].TrxID = systemTransaction.DataSystemCSV[i].TrxID
					v := result.DetailOfMatchedTransactions[bankName]
					v.TotalNumberMatchedTransactions += 1
					v.DetailTransaction = append(v.DetailTransaction, model.DetailTransaction{
						TrxID:            systemTransaction.DataSystemCSV[i].TrxID,
						UniqueIdentifier: bankTransaction.DataBankCSV[bankName][j].UniqueIdentifier,
						Amount:           bankTransaction.DataBankCSV[bankName][j].Amount,
						Date:             bankTransaction.DataBankCSV[bankName][j].DateString,
						TransactionTime:  systemTransaction.DataSystemCSV[i].TransactionTimeString,
					})
					result.DetailOfMatchedTransactions[bankName] = v
					delete(unMatchTransactionsBanks[bankName], bankTransaction.DataBankCSV[bankName][j].UniqueIdentifier)
					break
				} else {
					if _, ok := unMatchTransactionsBanks[bankName][bankTransaction.DataBankCSV[bankName][j].UniqueIdentifier]; !ok {
						unMatchTransactionsBanks[bankName][bankTransaction.DataBankCSV[bankName][j].UniqueIdentifier] = model.DetailTransaction{}
					}
					unMatchTransactionsBanks[bankName][bankTransaction.DataBankCSV[bankName][j].UniqueIdentifier] = model.DetailTransaction{
						UniqueIdentifier: bankTransaction.DataBankCSV[bankName][j].UniqueIdentifier,
						Amount:           bankTransaction.DataBankCSV[bankName][j].Amount,
						Date:             bankTransaction.DataBankCSV[bankName][j].DateString,
						Type:             bankTransaction.DataBankCSV[bankName][j].Type,
					}
				}

			}
		}
		if !systemTransaction.DataSystemCSV[i].MatchTransaction {
			totalUnmatchSystem++
			unMatchTransactionSystem = append(unMatchTransactionSystem, model.DetailTransaction{
				TrxID:           systemTransaction.DataSystemCSV[i].TrxID,
				Amount:          systemTransaction.DataSystemCSV[i].Amount,
				Type:            systemTransaction.DataSystemCSV[i].Type,
				TransactionTime: systemTransaction.DataSystemCSV[i].TransactionTimeString,
			})
		}
	}
	result.DetailOfUnmatchedTransactions = make(map[string]model.DetailUnmatchedTransaction)
	// processing unmatch transaction system
	if len(unMatchTransactionSystem) > 0 {
		unmatchSystem := model.DetailUnmatchedTransaction{
			Info:              "system transaction not found on any bank statement",
			DetailTransaction: unMatchTransactionSystem,
		}
		result.DetailOfUnmatchedTransactions["system"] = unmatchSystem
		result.TotalNumberUnmatchedTransactions += totalUnmatchSystem
	}
	// procesing unmatch transaction bank
	for bankName, unMatchTransactionsBank := range unMatchTransactionsBanks {
		if len(unMatchTransactionsBank) > 0 {
			temp := []model.DetailTransaction{}
			var count int64
			if len(unMatchTransactionsBank) > 0 {
				for _, unMatchTransactionBank := range unMatchTransactionsBank {
					temp = append(temp, unMatchTransactionBank)
					count++

				}
			}
			unmatchBank := model.DetailUnmatchedTransaction{
				Info:              "bank statement not found on any system transaction",
				DetailTransaction: temp,
			}
			result.DetailOfUnmatchedTransactions[bankName] = unmatchBank
			result.TotalNumberUnmatchedTransactions += count
		}
	}
	result.TotalTranscationsProcessed = bankTransaction.TotalData + systemTransaction.TotalData
	result.TotalNumberMatchedTransactions = func(r model.ReconciliationResponse) int64 {
		var result int64
		for _, v := range r.DetailOfMatchedTransactions {
			result += v.TotalNumberMatchedTransactions
		}
		return result * 2
	}(result)
	return result
}
