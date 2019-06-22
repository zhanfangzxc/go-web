package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	. "github.com/zhanfangzxc/web/handler"
	"github.com/zhanfangzxc/web/model"
	"github.com/zhanfangzxc/web/pkg/errno"
	"github.com/zhanfangzxc/web/util"
)

func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	SendResponse(c, nil, rsp)
	// var r struct {
	// 	Username string `json:"username"`
	// 	Password string `json:"password"`
	// }

	//第二版老代码
	// var r CreateRequest

	// if err := c.Bind(&r); err != nil {
	// 	SendResponse(c, errno.ErrBind, nil)
	// 	return
	// }

	// admin2 := c.Param("username")
	// log.Infof("URL username: %s", admin2)

	// desc := c.Query("desc")
	// log.Infof("URL key param desc: %s", desc)

	// contentType := c.GetHeader("Content-Type")
	// log.Infof("Header Content-Type: %s", contentType)

	// log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	// if r.Username == "" {
	// 	SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
	// }

	// if r.Password == "" {
	// 	SendResponse(c, fmt.Errorf("password is empty"), nil)
	// }

	// rsp := CreateResponse{
	// 	Username: r.Username,
	// }

	// SendResponse(c, nil, rsp)
	//第一版老代码
	// var err error

	// if err := c.Bind(&r); err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
	// 	return
	// }

	// log.Debugf("username is :[%s], password is [%s]", r.Username, r.Password)
	// if r.Username == "" {
	// 	err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message。")
	// 	log.Errorf(err, "Get an error")
	// }

	// if errno.IsErrUserNotFound(err) {
	// 	log.Debug("err type is ErrUserNotFound")
	// }

	// if r.Password == "" {
	// 	err = fmt.Errorf("password is empty")
	// }

	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
