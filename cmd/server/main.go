package main

import (
	"fmt"
	"os"

	"team_action/di"
	"team_action/pkg/logger"
	"team_action/pkg/web"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {
	g := gin.New()
	d := di.BuildContainer()

	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})

	svr := web.NewServer(g, d, l)

	svr.InitMiddleware()
	svr.InitRoutes()

	if err := svr.InitDB(); err != nil {
		return err
	}
	return svr.Start()
}
