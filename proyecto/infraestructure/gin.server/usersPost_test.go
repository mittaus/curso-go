package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock "example.com/mittaus/ddd-example/infraestructure/application.mock"
	server "example.com/mittaus/ddd-example/infraestructure/gin.server"
	jwt "example.com/mittaus/ddd-example/infraestructure/jwt.authHandler"
	"example.com/mittaus/ddd-example/testData"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
)

var userPostPath = "/api/users"

func TestUserPost_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	jane := testData.User("jane")
	ucHandler := mock.NewMockHandler(mockCtrl)

	ucHandler.EXPECT().
		UserCreate(jane.Name, jane.Email, jane.Password).
		Return(&jane, "authToken", nil).
		Times(1)

	gE := gin.Default()
	server.NewRouter(ucHandler, jwt.New("mySalt")).SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	baloo.New(ts.URL).
		Post(userPostPath).
		BodyString(`
		{
  			"user": {
    			"username": "` + testData.User("jane").Name + `",
				"email": "` + testData.User("jane").Email + `",
    			"password": "` + testData.User("jane").Password + `"
  			}
		}`).
		Expect(t).
		JSONSchema(testData.UserRespDefinition).
		Status(http.StatusCreated).
		Done()
}
