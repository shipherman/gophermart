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

func TestHandler_HandleGetWithdrawals(t *testing.T) {
	type mockBehavior func(r *mock.MockDBClientInt, wr []models.WithdrawResponse)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		withdrawals          []models.WithdrawResponse
		user                 string
		inputBody            string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Test_withdrawals_for_existing_user",
			mockBehavior: func(r *mock.MockDBClientInt, wr []models.WithdrawResponse) {
				r.EXPECT().SelectWithdrawals("user").Return(wr, nil).Times(1)
			},
			withdrawals: []models.WithdrawResponse{
				{
					OrderNum: "2673220062063",
					Sum:      10.1,
				},
				{
					OrderNum: "84481607050888",
					Sum:      20.2,
				},
			},
			user:                 "user",
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"order":"2673220062063","sum":10.1,"processed_at":"0001-01-01T00:00:00Z"},{"order":"84481607050888","sum":20.2,"processed_at":"0001-01-01T00:00:00Z"}]` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbclient := mock.NewMockDBClientInt(ctrl)
			tt.mockBehavior(dbclient, tt.withdrawals)

			h := &Handler{
				Client: dbclient,
			}

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/user/balance/withdraw",
				bytes.NewBufferString(tt.inputBody))

			req = req.WithContext(context.WithValue(req.Context(), models.UserCtxKey{}, tt.user))

			// Make Request
			h.HandleGetWithdrawals(w, req)
			result := w.Result()

			// tt.expectedResponseBody, _ = json.Marshal(tt.withdrawals)
			assert.Equal(t, tt.expectedStatusCode, result.StatusCode)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())

			result.Body.Close()
		})
	}
}
