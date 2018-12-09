package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/gateway"
)

type Message struct {
	Service string
	Data    []byte
}

type Engine struct {
	config     infrastructure.Config
	Engine     *gin.Engine
	logger     infrastructure.Logger
	consul     infrastructure.Consul
	httpClient *http.Client
}

func New(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, httpClient *http.Client) Engine {
	return Engine{
		config:     config,
		Engine:     gin.New(),
		logger:     logger,
		consul:     consul,
		httpClient: httpClient,
	}
}

func (e *Engine) InitMux() {

	//e.Engine.GET("/health", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"startTime":    s.StartTime,
			//"currDate": time.Now().Format("2006-01-02 15:04"),
			//"version":  "1.0",
			//"service":  g.config.Name,
		//})
		//fmt.Println("OK")
	//})


	e.Engine.POST("/scheduler", gin.WrapH(e.getHandler("scheduler")))

	//api := e.Engine.Group("/gateway")
	//{
	//	api.POST("/scheduler", gin.WrapH(e.getHandler("")))
	//
	//	api.GET("/redis", g.HandleRedis)
	//	api.POST("/redis", g.HandleRedis)
	//
	//	api.GET("/postres", func(c *gin.Context) {})
	//	api.POST("/postres", func(c *gin.Context) {})
	//}

}

func (e Engine) getHandler(s string) http.Handler {
	return gateway.SchedulerType{}
}