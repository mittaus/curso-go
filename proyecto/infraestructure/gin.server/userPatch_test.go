package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/mittaus/ddd-example/infraestructure/gin.server"
	"example.com/mittaus/ddd-example/infraestructure/jwt.authHandler"
	logger "example.com/mittaus/ddd-example/infraestructure/logrus.logger"
	"example.com/mittaus/ddd-example/infraestructure/application.mock"
	"example.com/mittaus/ddd-example/testData"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/baloo.v3"
)

var userPutPath = "/api/user"

func TestUserPut_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	jane := testData.User("jane")
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().
		UserEdit(testData.User("rick").Name, gomock.Any()).
		Return(&jane, "token", nil).
		Times(1)

	gE := gin.Default()
	jwtHandler := jwt.New("mySalt")
	router := server.NewRouter(ucHandler, jwtHandler)
	router.Logger = logger.SimpleLogger{}
	router.SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	token, err := jwtHandler.GenUserToken(testData.User("rick").Name)
	assert.NoError(t, err)
	baloo.New(ts.URL).
		Put(userPutPath).
		AddHeader("Authorization", testData.TokenPrefix+token).
		BodyString(`{
  			"user": {
				"bio": "` + testData.User("rick").Email + `",
				"name": "` + testData.User("rick").Name + `"
  			}
		}`).
		Expect(t).
		Status(http.StatusOK).
		JSONSchema(testData.UserRespDefinition).
		Done()
}
