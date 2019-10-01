package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"team_action/core/logger"
	"team_action/core/web/server"
	"team_action/di"
)

func main() {
	g := gin.New()
	d := di.BuildContainer()

	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})

	svr := server.NewServer(g, d, l)

	if err := svr.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}
