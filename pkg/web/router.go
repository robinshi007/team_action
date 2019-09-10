package web

import (
	"team_action/pkg/note"
	note_handler "team_action/pkg/note/web/handler"
	"team_action/pkg/user"
	user_handler "team_action/pkg/user/web/handler"
	"team_action/pkg/web/handler"
	mw "team_action/pkg/web/middleware"

	"github.com/gin-gonic/gin"
)

func (ds *dserver) InitRoutes() {
	ds.globalRoutes(ds.router)

	apiV1 := ds.router.Group("api/v1")
	ds.healthRoutes(apiV1)
	ds.userRoutes(apiV1)
	ds.noteRoutes(apiV1)
}

func (ds *dserver) globalRoutes(gr *gin.Engine) {
	jwtMW, err := mw.NewJWT("test zone", "secret key")
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
	jwtMW, err := mw.NewJWT("test zone", "secret key")
	if err != nil {
		ds.logger.Info("JWT Error:" + err.Error())
	}
	userRoutes := api.Group("/users")
	{
		var userSvc user.Service
		ds.cont.Invoke(func(u user.Service) {
			userSvc = u
		})

		usr := user_handler.NewUserCtrl(ds.logger, userSvc)

		userRoutes.GET("/", usr.GetAll)
		userRoutes.GET("/:id", usr.GetByID)
		userRoutes.Use(jwtMW.MiddlewareFunc())
		{
			userRoutes.POST("/", usr.Store)
			userRoutes.PUT("/:id", usr.Update)
			userRoutes.DELETE("/:id", usr.Delete)
		}
	}
}
func (ds *dserver) noteRoutes(api *gin.RouterGroup) {
	//	jwtMW, err := mw.NewJWT("test zone", "secret key")
	//	if err != nil {
	//		ds.logger.Info("JWT Error:" + err.Error())
	//	}
	noteRoutes := api.Group("/notes")
	{
		var noteSvc note.Service
		ds.cont.Invoke(func(u note.Service) {
			noteSvc = u
		})

		n := note_handler.NewNoteCtrl(ds.logger, noteSvc)

		noteRoutes.GET("/", n.GetAll)
		noteRoutes.GET("/:id", n.GetByID)
		//		noteRoutes.Use(jwtMW.MiddlewareFunc())
		//		{
		noteRoutes.POST("/", n.Store)
		noteRoutes.PUT("/:id", n.Update)
		noteRoutes.DELETE("/:id", n.Delete)
		//		}
	}
}
