package handler

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/context"
)


type SchedulerType struct {

}

func (g Gateway) HandleScheduler(c *gin.Context) {
	msg := &Message{}
	if err := c.BindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	servicePath := g.consul.DiscoveryService(msg.Service)
	fmt.Println(servicePath)

	if err := g.proxy.POST(servicePath, nil); err != nil {
		fmt.Println(err)
	}
}

func (s SchedulerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}