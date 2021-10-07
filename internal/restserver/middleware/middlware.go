// Package middleware provides middlewares for REST server.
package middleware

import "github.com/gin-gonic/gin"

var middlewares = map[string]gin.HandlerFunc{
	"recovery":  gin.Recovery(),
	"logger":    gin.Logger(),
	"requestid": requestID(),
}

// Get gets a middleware value by the given name.
func Get(name string) (mw gin.HandlerFunc, ok bool) {
	mw, ok = middlewares[name]
	return
}
