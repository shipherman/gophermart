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

func TestHandler_HandleBalance(t *testing.T) {
	type mockBehavior func(r *mock.MockDBClientInt, user models.User)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		inputBody            string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Test_withdrawals_for_existing_user",
			mockBehavior: func(r *mock.MockDBClientInt, user models.User) {
				r.EXPECT().SelectBalance("user").Return(models.BalanceResponse{
					Current:   0,
					Withdrawn: 0,
				}, nil).Times(1)
			},
			inputBody:            "",
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"current":0,"withdrawn":0}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbclient := mock.NewMockDBClientInt(ctrl)
			tt.mockBehavior(dbclient, models.User{Login: "user", Password: "pass"})

			h := &Handler{
				Client: dbclient,
			}

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/user/balance",
				bytes.NewBufferString(tt.inputBody))

			req = req.WithContext(context.WithValue(req.Context(), models.UserCtxKey{}, "user"))

			// Make Request
			h.HandleBalance(w, req)
			result := w.Result()

			assert.Equal(t, result.StatusCode, tt.expectedStatusCode)
			assert.Equal(t, w.Body.String(), tt.expectedResponseBody)

		})
	}
}
