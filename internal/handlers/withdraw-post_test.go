package handlers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shipherman/gophermart/internal/models"
	"github.com/shipherman/gophermart/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_HandlePostWithdraw(t *testing.T) {
	type mockBehavior func(r *mock.MockDBClientInt, wr models.WithdrawResponse)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		inputBody            string
		order                models.WithdrawResponse
		user                 string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Test_withdrawals_for_existing_user",
			mockBehavior: func(r *mock.MockDBClientInt, wr models.WithdrawResponse) {
				r.EXPECT().InsertWithdraw("user", wr).Return(nil).Times(1)
				r.EXPECT().UpdateWithdraw("user", wr.Sum).Return(nil).Times(1)
			},
			inputBody: `{"order":"2673220062063","sum":10.1}`,
			order: models.WithdrawResponse{
				OrderNum: "2673220062063",
				Sum:      10.1,
			},
			user:               "user",
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Test_withdrawals_not_enough_balance",
			mockBehavior: func(r *mock.MockDBClientInt, wr models.WithdrawResponse) {
				r.EXPECT().InsertWithdraw("user", wr).Return(fmt.Errorf("not anough bonuses to withdraw")).Times(1)
				r.EXPECT().UpdateWithdraw("user", wr.Sum).Return(nil).Times(1)
			},
			inputBody: `{"order":"2673220062063","sum":10.1}`,
			order: models.WithdrawResponse{
				OrderNum: "2673220062063",
				Sum:      10.1,
			},
			user:                 "user",
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "not anough bonuses to withdraw\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbclient := mock.NewMockDBClientInt(ctrl)
			tt.mockBehavior(dbclient, models.WithdrawResponse{OrderNum: tt.order.OrderNum, Sum: tt.order.Sum})

			h := &Handler{
				Client: dbclient,
			}

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/user/balance/withdraw",
				bytes.NewBufferString(tt.inputBody))

			req = req.WithContext(context.WithValue(req.Context(), models.UserCtxKey{}, tt.user))

			// Make Request
			h.HandlePostWithdraw(w, req)
			result := w.Result()

			assert.Equal(t, tt.expectedStatusCode, result.StatusCode)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())

			result.Body.Close()
		})
	}
}
