package handlers

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shipherman/gophermart/mock"
)

func TestHandleBalance(t *testing.T) {
	type mockBehavior func(r *mock.MockDBClientInt, u string)

	const contentType = "applicaiton/json"

	// Init mock controller
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name         string
		user         string
		mockBehavior mockBehavior
		want         want
	}{
		{
			name: "Test_check_balance",
			user: "user",
			mockBehavior: func(r *mock.MockDBClientInt, u string) {
				r.EXPECT().SelectBalance(u).Return(1)
			},
			want: want{
				contentType: contentType,
				statusCode:  http.StatusOK,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			dbcln := mock.NewMockDBClientInt(c)
			tc.mockBehavior(dbcln, tc.user)

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
