package api

import (
	"context"
	"net/http"
	"small/lib/log"
	go_micro_srv_user "small/srv/user/protoc/user"
	"strconv"

	"github.com/swaggo/swag/example/celler/httputil"

	"github.com/gin-gonic/gin"
)

// User rpc stuct
type User struct{}

// UserRPC UserRPC
var UserRPC go_micro_srv_user.UserService

// GetUser godoc
// @Description 获取用户信息
// @Tags Member
// @Summary 用户信息
// @Accept  json
// @Produce  json
// @Param id query int true "Account ID"
// @Success 200 {object} go_micro_srv_user.Response "用户信息"
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /user?id={id} [get]
func (s *User) GetUser(c *gin.Context) {
	id := c.Query("id")
	idi, _ := strconv.ParseInt(id, 10, 64)
	resp, err := UserRPC.GetUser(context.TODO(), &go_micro_srv_user.Request{
		Id: idi,
	})
	if err != nil {
		log.Error(err)
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200, resp)
}

// GetUserList godoc
// @Description 获取用户信息
// @Tags Member
// @Summary 用户信息
// @Accept  json
// @Produce  json
// @Param ob query string true "Order by"
// @Param page query int true "Page"
// @Param promotion_id query int true "User ID"
// @Success 200 {object} go_micro_srv_user.Responselist "用户信息"
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /user/list?page={page}&ob={ob}&promotion_id={promotion_id} [get]
func (s *User) GetUserList(c *gin.Context) {
	page := c.Query("page")
	promotionID := c.Query("promotion_id")
	pagei, _ := strconv.ParseInt(page, 10, 64)
	promotionIDi, _ := strconv.ParseInt(promotionID, 10, 64)
	resp, err := UserRPC.GetUserList(context.TODO(), &go_micro_srv_user.Requestlist{
		Page:        pagei,
		Ob:          "",
		PromotionId: promotionIDi,
	})
	if err != nil {
		log.Error(err)
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200, resp)
}
