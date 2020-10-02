package main_test

import (
	"testing"

	. "example.com/username/api-gin"
	"example.com/username/api-gin/Mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestUsuario_GetUsers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	MockInterface := Mocks.NewMockIUsuario(controller)
	var ctx *gin.Context
	MockInterface.EXPECT().GetUsers(ctx).Return(true)
	// Usuario{}.GetUsers(ctx)
	// result := Usuario{}.GetUsers(ctx)
	// assert.Equal(t, true, result)
}
