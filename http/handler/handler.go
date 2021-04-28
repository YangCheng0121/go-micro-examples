package handler

import (
	"context"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	helloworld "go-micro-examples/helloworld/proto"
)

//demo
type demo struct{}

func NewDemo() *demo {
	return &demo{}
}

func (a *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", a.demo)
}

func (a *demo) demo(c *gin.Context) {
	// create a service
	service := micro.NewService()
	service.Init()

	client := helloworld.NewHelloworldService("go.micro.srv.HelloWorld", service.Client())

	rsp, err := client.Call(context.Background(), &helloworld.Request{
		Name: "world!",
	})
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": rsp.Msg})
}
