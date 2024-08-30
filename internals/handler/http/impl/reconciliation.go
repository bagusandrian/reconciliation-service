package impl

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/bagusandrian/reconciliation-service/internals/model"
	"github.com/gofiber/fiber/v2"
)

// handler for reconciliation
func (h *handler) Reconciliation(c *fiber.Ctx) error {
	now := time.Now()
	req := model.ReconciliationRequest{}
	// parsing body into data ReconciliationRequest
	if err := c.BodyParser(&req); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(&fiber.Map{
			"header": model.HeaderResponse{
				Status:         fiber.StatusBadRequest,
				Error:          err.Error(),
				ProcessingTime: time.Since(now).String(),
			},
			"data": nil,
		})
	}
	// validation request parameters
	err := validationReconciliationRequest(&req)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(&fiber.Map{
			"header": model.HeaderResponse{
				Status:         fiber.StatusBadRequest,
				Error:          err.Error(),
				ProcessingTime: time.Since(now).String(),
			},
			"data": nil,
		})
	}
	resp, err := h.usecase.ReconciliationComparition(c.Context(), req)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(&fiber.Map{
			"header": model.HeaderResponse{
				Status:         fiber.StatusInternalServerError,
				Error:          err.Error(),
				ProcessingTime: time.Since(now).String(),
			},
			"data": nil,
		})
	}
	c.SendStatus(fiber.StatusOK)
	return c.JSON(&fiber.Map{
		"header": model.HeaderResponse{
			Status:         fiber.StatusOK,
			ProcessingTime: time.Since(now).String(),
		},
		"data": resp,
	})
}

func validationReconciliationRequest(req *model.ReconciliationRequest) error {
	var err error
	// validation reconciliation start date
	req.ReconciliationStartDate, err = time.Parse("2006-01-02", req.ReconciliationStartDateString)
	if err != nil {
		return fmt.Errorf("failed parsing reconciliaton_start_date: %+v", err)
	}
	req.ReconciliationStartDate = time.Date(req.ReconciliationStartDate.Year(), req.ReconciliationEndDate.Month(), req.ReconciliationStartDate.Day(), 0, 0, 0, 0, time.Local)
	// validation reconciliation end date
	req.ReconciliationEndDate, err = time.Parse("2006-01-02", req.ReconciliationEndDateString)
	if err != nil {
		return fmt.Errorf("failed parsing reconciliaton_end_date: %+v", err)
	}
	req.ReconciliationEndDate = time.Date(req.ReconciliationStartDate.Year(), req.ReconciliationEndDate.Month(), req.ReconciliationStartDate.Day(), 23, 59, 59, 999999999, time.Local)
	// validation start date must lower than start date
	if req.ReconciliationEndDate.Compare(req.ReconciliationStartDate) < 0 {
		return fmt.Errorf("start date: %s must lower than end date: %s", req.ReconciliationStartDateString, req.ReconciliationEndDateString)
	}
	// validate ext of file
	if ext := filepath.Ext(req.SystemTransactionCSVFilePath); ext != ".csv" {
		return fmt.Errorf("system transaction ext is not csv but: %s", ext)
	}
	// validation file path system transaction
	if _, err := os.Stat(req.SystemTransactionCSVFilePath); err != nil {
		return fmt.Errorf("system transaction csv file path not valid: %+v", err)
	}
	// validation file path bank statements
	if len(req.BankStatements) == 0 {
		return fmt.Errorf("empty of bank statement")
	}
	for _, v := range req.BankStatements {
		// validate bank name
		if v.BankName == "" {
			return fmt.Errorf("bank name cannot empty")
		}
		// validate ext file
		if ext := filepath.Ext(v.CSVFilePath); ext != ".csv" {
			return fmt.Errorf("%s bank statement ext is not csv but: %s", v.BankName, ext)
		}
		// validate file path bank
		if _, err := os.Stat(v.CSVFilePath); err != nil {
			return fmt.Errorf("bank statement csv of %s file path not valid: %+v", v.BankName, err)
		}
	}

	return nil
}
