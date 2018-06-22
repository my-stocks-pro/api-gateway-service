package apiserver

import (
	"github.com/my-stocks-pro/api-server/crud/approvedCrud"
	"github.com/my-stocks-pro/api-server/crud/earningsCrud"
	"github.com/my-stocks-pro/api-server/service/approvedServ"
	"github.com/my-stocks-pro/api-server/service/earningsServ"
	"github.com/kataras/iris/hero"
)

func (s *TypeServer) InitMiddleware() {

	crudApproved := approvedCrud.New(s.PSQL.NewConn())
	crudEarnings := earningsCrud.New(s.PSQL.NewConn())

	serviceApproved := approvedServ.New(crudApproved)
	serviceEarnings := earningsServ.New(crudEarnings)

	hero.Register(serviceApproved)
	hero.Register(serviceEarnings)

}