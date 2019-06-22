package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/zhanfangzxc/web/handler"
	"github.com/zhanfangzxc/web/model"
	"github.com/zhanfangzxc/web/pkg/errno"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
