package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/mittaus/ddd-example/infraestructure/gin.server"
	jwt "example.com/mittaus/ddd-example/infraestructure/jwt.authHandler"
	"example.com/mittaus/ddd-example/infraestructure/application.mock"
	"example.com/mittaus/ddd-example/testData"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/baloo.v3"
)

var userGetPath = "/api/user"

func TestUserGet_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	jane := testData.User("jane")
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().
		UserGet(jane.Name).
		Return(&jane, "authToken", nil).
		Times(1)

	jwtHandler := jwt.New("mySalt")
	gE := gin.Default()
	server.NewRouter(ucHandler, jwtHandler).SetRoutes(gE)
	authToken, err := jwtHandler.GenUserToken(testData.User("jane").Name)
	assert.NoError(t, err)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	baloo.New(ts.URL).
		Get(userGetPath).
		AddHeader("Authorization", testData.TokenPrefix+authToken).
		Expect(t).
		Status(http.StatusOK).
		JSONSchema(testData.UserRespDefinition).
		Done()
}
