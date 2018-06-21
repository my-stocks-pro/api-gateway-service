package apiserver

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/db"
)

type TypeServer struct {
	PSQL *db.TypePSQL
	RDS  *db.TypeRedis
	API  *iris.Application
}


func NewServer() *TypeServer {
	return &TypeServer{
		db.NewPSQL(),
		db.NewRDS(),
		iris.New(),
	}
}