package application_test

import (
	"testing"

	mock "example.com/mittaus/ddd-example/infraestructure/application.mock"
	"example.com/mittaus/ddd-example/testData"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authToken := "token"
	rick := testData.User("rick")
	i := mock.NewMockedInteractor(mockCtrl)
	i.UserRW.EXPECT().GetByEmailAndPassword(rick.Email, rick.Password).Return(&rick, nil).Times(1)
	i.AuthHandler.EXPECT().GenUserToken(rick.Name).Return(authToken, nil)
	retUser, token, err := i.GetUCHandler().UserLogin(rick.Email, rick.Password)

	assert.NoError(t, err)
	assert.Equal(t, authToken, token)
	assert.Equal(t, rick, *retUser)

}
