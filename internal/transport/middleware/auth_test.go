package middleware

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shipherman/gophermart/mock"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticator_Auth(t *testing.T) {
	type mockBehavior func(r *mock.MockDBClientInt)

	tests := []struct {
		name          string
		mockBehavior  mockBehavior
		user          string
		password      string
		expectedJWT   string
		expectedError error
	}{
		{
			name: "Test_existing_user",
			mockBehavior: func(r *mock.MockDBClientInt) {
				r.EXPECT().SelectUserExistence("user", "pass").Return(true, nil).Times(1)
			},
			user:          "user",
			password:      "pass",
			expectedJWT:   ".*", // some JWT string should be returned
			expectedError: nil,
		},
		{
			name: "Test_non-existent_user",
			mockBehavior: func(r *mock.MockDBClientInt) {
				r.EXPECT().SelectUserExistence("user", "pass").Return(false, nil).Times(1)
			},
			user:          "user",
			password:      "pass",
			expectedJWT:   "", // empty JWT string should be returned
			expectedError: ErrUserDoesNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbclient := mock.NewMockDBClientInt(ctrl)
			tt.mockBehavior(dbclient)

			a := NewAuthenticator(dbclient)
			response, err := a.Auth(tt.user, tt.password)

			assert.Regexp(t, tt.expectedJWT, response)
			assert.ErrorIs(t, err, tt.expectedError)

		})
	}
}
