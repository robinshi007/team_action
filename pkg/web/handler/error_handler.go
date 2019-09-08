package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"team_action/pkg/web/types"
)

// NotFoundResponse - 404
func NotFoundResponse(c *gin.Context) {
	HandleErrorCodeRepsonse("1102", c)
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
					HandleErrorCodeCustomRepsonse("1102", []string{"Oops! Internal error with XMLHttpRequest, please try again."}, c)
					return
				}
				HandleErrorCodeRepsonse("1102", c)
			}
		}(c)

		c.Next()
	}
}

// HandleErrorRepsonse -
func HandleErrorRepsonse(err error, ctx *gin.Context) {
	if err != nil {
		ge, ok := errors.Cause(err).(types.GeneralError)
		if ok {
			HandleErrorCodeCustomRepsonse(string(ge.Code()), ge.Messages(), ctx)
		}
		ie, ok := errors.Cause(err).(types.InternalError)
		if ok && ie.Internal() {
			// log the info
		} else {
			// log the info
		}
		HandleErrorCodeCustomRepsonse("1103", []string{err.Error()}, ctx)
	}
}
func HandleBadRequestRepsonse(err error, ctx *gin.Context) {
	if err != nil {
		ge, ok := errors.Cause(err).(types.GeneralError)
		if ok {
			HandleErrorCodeCustomRepsonse(string(ge.Code()), ge.Messages(), ctx)
		}
		ie, ok := errors.Cause(err).(types.InternalError)
		if ok && ie.Internal() {
			// log the info
		} else {
			// log the info
		}
		HandleErrorCodeCustomRepsonse("1101", []string{err.Error()}, ctx)
	}
}

// HandleErrorCodeRepsonse -
func HandleErrorCodeRepsonse(codeStr string, ctx *gin.Context) {
	var code types.ErrorCode = types.ErrorCode(codeStr)
	ctx.JSON(types.GetHTTPStatus(code), &types.ErrorResponse{
		Code:   code,
		Errors: []string{types.GetErrorMessage(code)},
	})
}

// HandleErrorCodeCustomRepsonse -
func HandleErrorCodeCustomRepsonse(codeStr string, messages []string, ctx *gin.Context) {
	var code types.ErrorCode = types.ErrorCode(codeStr)
	if len(messages) >= 1 {
		strs := strings.Split(messages[0], "\n")
		if len(strs) > 1 {
			var res []string
			for _, v := range strs {
				res = append(res, v)
			}
			ctx.JSON(types.GetHTTPStatus(code), &types.ErrorResponse{
				Code:   code,
				Errors: res,
			})
			return
		}
	}
	ctx.JSON(types.GetHTTPStatus(code), &types.ErrorResponse{
		Code:   code,
		Errors: messages,
	})
}
