package handlers

import (
	"bytes"
	"context"
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
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Test_withdrawals_for_existing_user",
			mockBehavior: func(r *mock.MockDBClientInt, wr models.WithdrawResponse) {
				r.EXPECT().InsertWithdraw("user", "10").Return(nil).Times(1)
			},
			inputBody:          "",
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbclient := mock.NewMockDBClientInt(ctrl)
			tt.mockBehavior(dbclient, models.WithdrawResponse{OrderNum: "2673220062063", Sum: 10})

			h := &Handler{
				Client: dbclient,
			}

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/user/balance/withdraw",
				bytes.NewBufferString(tt.inputBody))

			req = req.WithContext(context.WithValue(req.Context(), models.UserCtxKey{}, "user"))

			// Make Request
			h.HandlePostWithdraw(w, req)
			result := w.Result()

			assert.Equal(t, result.StatusCode, tt.expectedStatusCode)
			assert.Equal(t, w.Body.String(), tt.expectedResponseBody)

			result.Body.Close()
		})
	}

}
