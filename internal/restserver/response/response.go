package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
)

type dataRespBodyWrapper struct {
	Data interface{} `json:"data"`
}

type errRespBodyWrapper struct {
	Err errInfo `json:"err"`
}

type errInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		agent := pkgerr.GetRESTAgentByError(err)
		c.JSON(
			agent.HTTPStatus(),
			errRespBodyWrapper{
				errInfo{Code: int(agent.Code()), Msg: agent.Msg()},
			},
		)
		return
	}
	if data != nil {
		c.JSON(http.StatusOK, dataRespBodyWrapper{data})
		return
	}
	c.JSON(http.StatusOK, nil)
}
