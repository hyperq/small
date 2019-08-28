package main

import (
	"log"
	"small/router"

	"github.com/micro/go-micro/web"
	_ "small/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title small api
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @basePath /
// @host localhost:8080
func main() {
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	service.Init()
	r := router.GetRouter()
	url := ginSwagger.URL("http://localhost:8080/user/swagger/doc.json")
	r.GET("/user/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,url))

	service.Handle("/", r)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
