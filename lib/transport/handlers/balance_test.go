package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleBalance(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	mockClient := mock_db.NewMockDBClientInt()
	mockClient.EXPECT().Get()

	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name       string
		request    string
		httpMethod string
		want       want
	}{
		{
			name:       "Test_check_balance",
			request:    "/api/user/balance",
			httpMethod: http.MethodGet,
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusOK,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.httpMethod, tc.request, nil)
			w := httptest.NewRecorder()

			h := NewHandler(mockClient)
			h.HandleBalance(w, req)

			result := w.Result()
			assert.Equal(t, tc.want.contentType, result.Header.Get("Content-Type"))
			assert.Equal(t, tc.want.statusCode, result.StatusCode)

			err := result.Body.Close()
			require.NoError(t, err)

		})
	}

}
