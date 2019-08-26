package api

import (
	"context"
	"fmt"
	user "small/srv/user/protoc/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User rpc stuct
type User struct{}

// UserRPC UserRPC
var UserRPC user.UserService

// GetUser get user info by id
func (s *User) GetUser(c *gin.Context) {
	id := c.Query("id")
	idi, _ := strconv.ParseInt(id, 10, 64)
	resp, err := UserRPC.GetUser(context.TODO(), &user.Request{
		Id: idi,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, resp)
}

// GetUserList get user list by id
func (s *User) GetUserList(c *gin.Context) {
	page := c.Query("page")
	promotionID := c.Query("promotion_id")
	pagei, _ := strconv.ParseInt(page, 10, 64)
	promotionIDi, _ := strconv.ParseInt(promotionID, 10, 64)
	resp, err := UserRPC.GetUserList(context.TODO(), &user.Requestlist{
		Page:        pagei,
		Ob:          "",
		PromotionId: promotionIDi,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, resp)
}
