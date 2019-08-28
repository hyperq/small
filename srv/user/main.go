package main

import (
	"context"
	"time"

	dao "small/srv/user/dao"
	user "small/srv/user/protoc/user"

	"small/lib/log"

	jsoniter "github.com/json-iterator/go"

	"github.com/micro/go-micro"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// User rpc stuct
type User struct {
	d *dao.Dao
}

// NewUser new a user
func NewUser() *User {

	return &User{
		d: dao.New(),
	}
}

// GetUser 重写GetUser
func (s *User) GetUser(ctx context.Context, req *user.Request, rsp *user.Response) error {
	d := dao.New()
	res, err := d.QueryUserByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	jres, err := json.Marshal(res)
	if err != nil {
		log.Error(err)
		return err
	}
	err = json.Unmarshal(jres, rsp)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// GetUserList 重写GetUserList
func (s *User) GetUserList(ctx context.Context, req *user.Requestlist, rsp *user.Responselist) error {
	d := dao.New()
	res, err := d.QueryUserList(req.Page, req.PromotionId, req.Ob)
	if err != nil {
		log.Error(err)
		return err
	}
	jres, err := json.Marshal(res)
	if err != nil {
		log.Error(err)
		return err
	}
	err = json.Unmarshal(jres, &rsp.Userlist)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()
	User := NewUser()
	// Register Handlers
	user.RegisterUserHandler(service.Server(), User)
	// Run server
	if err := service.Run(); err != nil {
		log.Error(err)
	}
}
