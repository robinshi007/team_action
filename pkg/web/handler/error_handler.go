package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// NotFoundResponse - 404
func NotFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    ErrNotFound.Error(),
		"message": "404 Page not found!!!",
	})
}

// XHR check if is XMLHttpRequest
func XHR(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

// InternalServerErrRecover - 50x, please use as gin middeware
func InternalServerErrRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(c *gin.Context) {
			if rec := recover(); rec != nil {
				// check if is XHR
				if XHR(c) {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    ErrInternalServerError.Error(),
						"message": "Oops with xmlhttprequest! please retry.",
					})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    ErrInternalServerError.Error(),
					"message": "Oops! please retry.",
				})
			}
		}(c)

		c.Next()
	}
}
