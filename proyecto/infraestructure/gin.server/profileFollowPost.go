package server

import (
	"net/http"

	"example.com/mittaus/ddd-example/infraestructure/json.formatter"
	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) profileFollowPost(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	user, err := rH.ucHandler.ProfileUpdateFollow(rH.getUserName(c), c.Param("username"), true)
	if err != nil {
		log(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": formatter.NewProfileFromDomain(*user, false)})
}
