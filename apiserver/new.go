package apiserver

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/db"
	"github.com/my-stocks-pro/api-server/config"
)


type TypeServer struct {
	Config *config.TypeConfig
	PSQL *db.TypePSQL
	RDS  *db.TypeRedis
	API  *iris.Application
}


func NewServer() *TypeServer {
	return &TypeServer{
		Config: config.LoadConfig(),
		PSQL:   db.NewPSQL(),
		RDS:    db.NewRDS(),
		API:    iris.New(),
	}
}