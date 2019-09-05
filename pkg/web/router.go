package web

import (
	"github.com/gin-gonic/gin"

	"team_action/pkg/user"
	user_handler "team_action/pkg/user/web/handler"
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

	auth := gr.Group("/auth")
	auth.GET("/refresh_token", jwtMW.RefreshHandler)
	auth.Use(jwtMW.MiddlewareFunc())
	{
		auth.GET("/hello", a.SayHi)
	}

	gr.GET("/crash", a.Crash)
	gr.NoRoute(handler.NotFoundResponse)
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
