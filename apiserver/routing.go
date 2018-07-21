package apiserver

import (
	"github.com/kataras/iris"
)

func (s *TypeServer) Routs() {

	s.API.Get("/welcome", func(ctx iris.Context) {
		ctx.Writef("Hello from MY_STOCKS_PRO SERVER")
	})

	s.API.PartyFunc("/api", func(r iris.Party) {
		r.Get("/redis", ServiceApprovedRDS.GetALL)
	})


	s.API.PartyFunc("/data", func(r iris.Party) {
		//r.Get("/psql/{service:string}", project.GetAll)
		r.Post("/redis/approved", ServiceApprovedRDS.PostALL)
		r.Post("/psql/approved", ServiceApprovedPSQL.PostALL)
		r.Post("/psql/earnings", ServiceEarningsPSQL.PostALL)
	})


	s.API.PartyFunc("/history/", func(r iris.Party){
		r.Get("/approved", ServiceApprovedPSQL.GetHistory)
		r.Get("/earnings", ServiceEarningsPSQL.GetHistory)
	})
}