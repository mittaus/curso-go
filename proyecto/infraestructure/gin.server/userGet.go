package server

import (
	"net/http"

	"example.com/mittaus/ddd-example/infraestructure/json.formatter"
	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userGet(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	user, token, err := rH.ucHandler.UserGet(rH.getUserName(c))
	if err != nil {
		log(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": formatter.NewUserResp(*user, token)})
}
