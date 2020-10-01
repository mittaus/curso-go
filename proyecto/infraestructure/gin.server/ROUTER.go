package server

import (
	"fmt"
	"net/http"

	"strings"

	"errors"

	"example.com/mittaus/ddd-example/application"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type RouterHandler struct {
	ucHandler   application.Handler
	authHandler application.AuthHandler
	Logger      application.Logger
}

func NewRouter(i application.Handler, auth application.AuthHandler) RouterHandler {
	return RouterHandler{
		ucHandler:   i,
		authHandler: auth,
	}
}

func NewRouterWithLogger(i application.Handler, auth application.AuthHandler, logger application.Logger) RouterHandler {
	return RouterHandler{
		ucHandler:   i,
		authHandler: auth,
		Logger:      logger,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.Use(rH.errorCatcher())

	rH.profileRoutes(api)
	rH.usersRoutes(api)
	rH.articlesRoutes(api)

	api.GET("/tags", rH.tagsGet) //tags
}

func (rH RouterHandler) profileRoutes(api *gin.RouterGroup) {
	profiles := api.Group("/profiles")
	profiles.GET("/:username", rH.profileGet)                                        // Get a profile of a user of the system. Auth is optional
	profiles.POST("/:username/follow", rH.jwtMiddleware(), rH.profileFollowPost)     // Follow a user by username
	profiles.DELETE("/:username/follow", rH.jwtMiddleware(), rH.profileFollowDelete) // Unfollow a user by username
}

func (rH RouterHandler) usersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	users.POST("", rH.userPost)            // Register a new user
	users.POST("/login", rH.userLoginPost) // Login for existing user

	user := api.Group("/user")
	user.GET("", rH.jwtMiddleware(), rH.userGet)     // Gets the currently logged-in user
	user.PUT("", rH.jwtMiddleware(), rH.userPatch)   // WARNING : it's a in fact a PATCH request in the API contract !!!
	user.PATCH("", rH.jwtMiddleware(), rH.userPatch) // just in case it's fixed one day....
}

func (rH RouterHandler) articlesRoutes(api *gin.RouterGroup) {
	articles := api.Group("/articles")

	articles.GET("", rH.articlesFilteredGet)
	articles.POST("", rH.jwtMiddleware(), rH.articlePost)
	articles.GET("/:slug", func(c *gin.Context) { // ugly api contract !
		if c.Param("slug") == "feed" {
			rH.jwtMiddleware()(c)
			rH.articlesFeedGet(c)
			return
		}

		rH.articleGet(c)
	})

	articles.PUT("/:slug", rH.jwtMiddleware(), rH.articlePut)
	articles.DELETE("/:slug", rH.jwtMiddleware(), rH.articleDelete)
	articles.POST("/:slug/favorite", rH.jwtMiddleware(), rH.updateFavorite)
	articles.DELETE("/:slug/favorite", rH.jwtMiddleware(), rH.updateFavorite)

	articles.GET("/:slug/comments", rH.commentsGet)
	articles.POST("/:slug/comments", rH.jwtMiddleware(), rH.commentPost)
	articles.DELETE("/:slug/comments/:id", rH.jwtMiddleware(), rH.commentDelete)
}

const userNameKey = "userNameKey"

func (rH RouterHandler) jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := getJWT(c.GetHeader("Authorization"))
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		userName, err := rH.authHandler.GetUserName(jwt)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.SetAccepted()
		c.Set(userNameKey, userName)
		c.Next()
	}
}

func (rH RouterHandler) getUserNameFromToken(c *gin.Context) string {
	jwt, err := getJWT(c.GetHeader("Authorization"))
	if err != nil {
		return ""
	}

	userName, err := rH.authHandler.GetUserName(jwt)
	if err != nil {
		return ""
	}

	return userName
}

func getJWT(authHeader string) (string, error) {
	splitted := strings.Split(authHeader, "Token ")
	if len(splitted) != 2 {
		return "", errors.New("malformed header")
	}
	return splitted[1], nil
}

func (rH RouterHandler) errorCatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() > 399 {
			c.Render(
				c.Writer.Status(),
				render.Data{
					ContentType: "application/json; charset=utf-8",
					Data:        []byte(`{"errors": {"body": ["wooops, something went wrong !"]}}`),
				},
			)
		}
	}
}

func (RouterHandler) getUserName(c *gin.Context) string {
	if userName, ok := c.Keys[userNameKey].(string); ok {
		return userName
	}
	return ""
}

// log is used to "partially apply" the title to the rH.logger.Log function
// so we can see in the logs from which route the log comes from
func (rH RouterHandler) log(title string) func(...interface{}) {
	return func(logs ...interface{}) {
		rH.Logger.Log(title, logs)
	}
}

func (RouterHandler) MethodAndPath(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
}
