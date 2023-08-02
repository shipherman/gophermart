package handlers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shipherman/gophermart/internal/models"
	"github.com/shipherman/gophermart/mock"
)

func TestHandler_HandleBalance(t *testing.T) {
	type mockBehavior func(r *mock.MockDBClientInt, user models.User)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			mockBehavior: func(r *mock.MockDBClientInt, user models.User) {
				r.EXPECT().SelectBalance("user").Return(0)
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbclient := mock.NewMockDBClientInt(ctrl)
			tt.mockBehavior(dbclient, models.User{Login: "user", Password: "pass"})

			h := &Handler{
				Client: tt.fields.Client,
			}
			h.HandleBalance(tt.args.w, tt.args.r)
		})
	}
}
