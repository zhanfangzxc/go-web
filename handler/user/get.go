package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/zhanfangzxc/web/handler"
	"github.com/zhanfangzxc/web/model"
	"github.com/zhanfangzxc/web/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}
	SendResponse(c, nil, user)
}
