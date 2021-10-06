package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func requestID() gin.HandlerFunc {
	return requestid.New(requestid.Config{
		Generator: func() string {
			return uuid.NewString()
		},
	})
}
