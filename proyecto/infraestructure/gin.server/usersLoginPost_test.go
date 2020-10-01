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
	"gopkg.in/h2non/baloo.v3"
)

var userLoginPostPath = "/api/users/login"

func TestUserLoginPost_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	jane := testData.User("jane")
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().
		UserLogin(jane.Email, jane.Password).
		Return(&jane, "authToken", nil).
		Times(1)

	jwtHandler := jwt.New("mySalt")
	gE := gin.Default()
	server.NewRouter(ucHandler, jwtHandler).SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	baloo.New(ts.URL).
		Post(userLoginPostPath).
		BodyString(`
		{
  			"user": {
				"email": "` + testData.User("jane").Email + `",
    			"password": "` + testData.User("jane").Password + `"
  			}
		}`).
		Expect(t).
		Status(http.StatusOK).
		JSONSchema(testData.UserRespDefinition).
		Done()
}
