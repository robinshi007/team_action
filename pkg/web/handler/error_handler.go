package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"team_action/pkg/cerrors"
)

// NotFoundResponse - 404 error
func NotFoundResponse(c *gin.Context) {
	HandlePublicError(&cerrors.CustomError{
		ErrorCode: "1102",
		Errors:    []string{cerrors.GetErrorMessage("1102")},
	}, c)
}

// XHR check if is XMLHttpRequest
func XHR(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

// ErrorRecover -  50x for un catched painc error
func ErrorRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(c *gin.Context) {
			if rec := recover(); rec != nil {
				// check if is XHR
				if XHR(c) {
					HandlePublicError(&cerrors.CustomError{
						ErrorCode: "1103",
						Errors:    []string{"Oops! something went wrong with XMLHttpRequest, please contact system admin."},
					}, c)
					return
				}
				HandlePublicError(&cerrors.CustomError{
					ErrorCode: "1103",
					Errors:    []string{"Oops! something went wrong, please contact system admin."},
				}, c)
			}
		}(c)
		c.Next()
	}
}

// HandlePublicError -
func HandlePublicError(err cerrors.GeneralError, ctx *gin.Context) {
	HTTPCode := cerrors.GetHTTPStatus(err.Code())
	ctx.JSON(HTTPCode, err)
}

// HandleInternalError -
func HandleInternalError(err cerrors.InternalError, ctx *gin.Context) {
	// log the message
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"errors": []string{cerrors.GetErrorMessage("1103")},
	})
}

// ErrorHandling - middeware to handling errors
func ErrorHandling() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()
		if err == nil {
			return
		}
		// public error
		if gError, ok := err.Err.(cerrors.GeneralError); ok {
			HandlePublicError(gError, ctx)
			return
		}
		// internal error
		if iError, ok := err.Err.(cerrors.InternalError); ok {
			HandleInternalError(iError, ctx)
			return
		}

		// logger the info
		//fmt.Println("Oops, unkown error occurs")
		HandlePublicError(&cerrors.ParamError{
			ErrorCode: "1103",
			Errors:    []string{"Oops! something went wrong, please contact system admin."},
		}, ctx)
	}
}
