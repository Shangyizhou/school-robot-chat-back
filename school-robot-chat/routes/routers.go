package routers

import (
	"net/http"
	"school-robot-chat/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r := gin.Default()
	routerGroup := r.Group("/api/v1")
	routerGroup.POST("/login", controller.LoginHandler)
	routerGroup.POST("/signup", controller.SignUpHandler)
	routerGroup.GET("/refresh_token", controller.RefreshTokenHandler)

	routerGroup.Use(controller.JWTAuthMiddleware())
	{
		routerGroup.GET("/community", controller.CommunityHandler)
		routerGroup.GET("/community/:id", controller.CommunityDetailHandler)

		routerGroup.POST("/post", controller.CreatePostHandler)
		routerGroup.GET("/post/:id", controller.PostDetailHandler)
		routerGroup.GET("/post", controller.PostListHandler)

		routerGroup.GET("/post2", controller.PostList2Handler)

		routerGroup.POST("/vote", controller.VoteHandler)

		routerGroup.POST("/comment", controller.CommentHandler)
		routerGroup.GET("/comment", controller.CommentListHandler)

		routerGroup.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
