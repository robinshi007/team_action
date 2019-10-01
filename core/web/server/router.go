package server

import (
	"github.com/gin-gonic/gin"

	note_handler "team_action/apps/note/web/handler"
	"team_action/core/user"
	user_handler "team_action/core/user/web/handler"
	"team_action/core/web/handler"
	mw "team_action/core/web/middleware"

	"team_action/apps/note"
)

// init Routes -
func (ds *DServer) initRoutes() {
	ds.globalRoutes(ds.router)

	apiV1 := ds.router.Group("api/v1")
	ds.healthRoutes(apiV1)
	ds.authRoutes(apiV1)
	ds.userRoutes(apiV1)

	noteAppV1 := apiV1.Group("/noteapp")
	ds.noteAppRoutes(noteAppV1)
}

func (ds *DServer) globalRoutes(gr *gin.Engine) {
	a := handler.NewHelloCtrl()
	gr.GET("/crash", a.Crash)
	gr.NoRoute(handler.NotFoundResponse)
}

func (ds *DServer) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")
	{
		h := handler.NewHealthCtrl()
		healthRoutes.GET("/", h.Ping)
	}
}
func (ds *DServer) authRoutes(api *gin.RouterGroup) {
	jwtMW, err := mw.NewJWT("test zone", "secret key")
	if err != nil {
		ds.logger.Info("JWT Error:" + err.Error())
	}
	a := handler.NewHelloCtrl()

	api.POST("/login", jwtMW.LoginHandler)
	auth := api.Group("/auth")
	auth.GET("/refresh_token", jwtMW.RefreshHandler)
	auth.Use(jwtMW.MiddlewareFunc())
	{
		auth.GET("/hello", a.SayHi)
	}
}
func (ds *DServer) userRoutes(api *gin.RouterGroup) {
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

		userRoutes.Use(jwtMW.MiddlewareFunc())
		{
			userRoutes.GET("/", usr.GetAll)
			userRoutes.GET("/:id", usr.GetByID)
			userRoutes.POST("/", usr.Store)
			userRoutes.PUT("/:id", usr.Update)
			userRoutes.PUT("/:id/update_password", usr.UpdatePassword)
			userRoutes.PUT("/:id/update_last_login", usr.UpdateLastLogin)
			userRoutes.DELETE("/:id", usr.Delete)
		}
	}
}
func (ds *DServer) noteAppRoutes(app *gin.RouterGroup) {
	jwtMW, err := mw.NewJWT("test zone", "secret key")
	if err != nil {
		ds.logger.Info("JWT Error:" + err.Error())
	}
	noteRoutes := app.Group("/notes")
	{
		var noteSvc note.NoteService
		ds.cont.Invoke(func(u note.NoteService) {
			noteSvc = u
		})
		nh := note_handler.NewNoteCtrl(ds.logger, noteSvc)

		noteRoutes.GET("/", nh.GetAll)
		noteRoutes.GET("/:id", nh.GetByID)
		noteRoutes.GET("/:id/search", nh.Search)
		noteRoutes.Use(jwtMW.MiddlewareFunc())
		{
			noteRoutes.POST("/", nh.Store)
			noteRoutes.PUT("/:id", nh.Update)
			noteRoutes.DELETE("/:id", nh.Delete)
		}
	}
	categoryRoutes := app.Group("/categories")
	{
		var categorySvc note.CategoryService
		ds.cont.Invoke(func(u note.CategoryService) {
			categorySvc = u
		})
		ch := note_handler.NewCategoryCtrl(ds.logger, categorySvc)

		categoryRoutes.GET("/", ch.GetAll)
		categoryRoutes.GET("/:id", ch.GetByID)
		categoryRoutes.Use(jwtMW.MiddlewareFunc())
		{
			categoryRoutes.POST("/", ch.Store)
			categoryRoutes.PUT("/:id", ch.Update)
			categoryRoutes.DELETE("/:id", ch.Delete)
		}
	}
}
