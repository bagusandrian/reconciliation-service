package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

func (u *usecase) GetDummy(ctx context.Context, request model.GetDummyRequest) (response model.GetDummyResponse, err error) {
	now := time.Now()
	if request.Wait > 0 {
		if request.Wait > 5 {
			request.Wait = 5
		}
		for i := 0; i < request.Wait; i++ {
			time.Sleep(time.Second)
		}
	}
	response.Text = request.Text
	response.ProcessingTIme = time.Since(now).String()
	if request.StatusCode >= 400 && request.StatusCode < 600 {
		return response, fmt.Errorf("custom error by status-code:%v", request.StatusCode)
	}

	return response, nil
}
