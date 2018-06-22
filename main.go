package main

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/apiserver"
)


func main(){

	server := apiserver.NewServer()
	server.API.Logger().SetLevel("debug")
	server.PSQL.MakeMigrations(server.PSQL.NewConn())

	server.InitMiddleware()

	server.Routs()

	errServer := server.API.Run(
		// Start the web serverHTTP at localhost:80
		iris.Addr(":8008"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err serverHTTP closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)

	if errServer != nil {
		panic(errServer)
	}

}
