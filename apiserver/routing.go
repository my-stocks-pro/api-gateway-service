package apiserver

import (
	"github.com/kataras/iris"
)


func (s *TypeServer) Routs() {

	s.API.Get("/welcome", func(ctx iris.Context) {
		ctx.Writef("Hello from MY_STOCKS_PRO SERVER")
	})

	s.API.PartyFunc("/room", func(r iris.Party) {
		r.Get("/welcome", func(ctx iris.Context) {
			ctx.Writef("Hello from ROOM")
		})
	})


	s.API.PartyFunc("/data", func(r iris.Party) {
		//r.Get("/psql/{service:string}", project.GetAll)
		r.Post("/redis/approved", ServiceRedis.PostALL)
	})

}