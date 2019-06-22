package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanfangzxc/web/handler"
	"github.com/zhanfangzxc/web/pkg/errno"
	"github.com/zhanfangzxc/web/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
