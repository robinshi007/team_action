package web

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"team_action/pkg/user"
	user_handler "team_action/pkg/user/delivery/handler"
	"team_action/pkg/web/handler"
	mw "team_action/pkg/web/middleware"
)

func (ds *dserver) InitRoutes() {
	ds.globalRoutes(ds.router)

	apiV1 := ds.router.Group("api/v1")
	ds.healthRoutes(apiV1)
	ds.userRoutes(apiV1)
}

func (ds *dserver) globalRoutes(gr *gin.Engine) {
	jwtMW, err := mw.NewJWT()
	if err != nil {
		ds.logger.Info("JWT Error:" + err.Error())
	}
	a := handler.NewHelloCtrl()
	gr.POST("/login", jwtMW.LoginHandler)
	gr.NoRoute(jwtMW.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		ds.logger.Infof("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := gr.Group("/auth")
	auth.GET("/refresh_token", jwtMW.RefreshHandler)
	auth.Use(jwtMW.MiddlewareFunc())
	{
		auth.GET("/hello", a.SayHi)
	}
}

func (ds *dserver) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")
	{
		h := handler.NewHealthCtrl()
		healthRoutes.GET("/", h.Ping)
	}
}
func (ds *dserver) userRoutes(api *gin.RouterGroup) {
	userRoutes := api.Group("/users")
	{
		var userSvc user.Service
		ds.cont.Invoke(func(u user.Service) {
			userSvc = u
		})

		usr := user_handler.NewUserCtrl(ds.logger, userSvc)

		userRoutes.GET("/", usr.GetAll)
		userRoutes.POST("/", usr.Store)
		userRoutes.GET("/:id", usr.GetByID)
		userRoutes.PUT("/:id", usr.Update)
		userRoutes.DELETE("/:id", usr.Delete)
	}
}
