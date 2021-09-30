package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
)

type DataRespBodyWrapper struct {
	Data interface{} `json:"data"`
}

type ErrRespBodyWrapper struct {
	Err ErrInfo `json:"err"`
}

type ErrInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		agent := pkgerr.GetRESTAgentByError(err)
		c.JSON(
			agent.HTTPStatus(),
			ErrRespBodyWrapper{
				ErrInfo{Code: int(agent.Code()), Msg: agent.Msg()},
			},
		)
		return
	}
	c.JSON(http.StatusOK, DataRespBodyWrapper{data})
}
