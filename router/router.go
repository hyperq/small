package router

import (
	api "small/api/user"

	user "small/srv/user/protoc/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

// GetRouter get router config
func GetRouter() *gin.Engine {
	r := gin.Default()
	users := new(api.User)
	// setup Greeter Server Client
	api.UserRPC = user.NewUserService("go.micro.srv.user", client.DefaultClient)

	r.GET("/user", users.GetUser)
	r.GET("/user/list", users.GetUserList)

	// v1 := r.Group("/api/v1")
	// {
	// }
	return r
}
