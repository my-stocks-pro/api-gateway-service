package apiserver

import (
	"github.com/my-stocks-pro/api-server/crud/approvedCrud"
	"github.com/my-stocks-pro/api-server/crud/earningsCrud"
	"github.com/my-stocks-pro/api-server/service/approvedServ"
	"github.com/my-stocks-pro/api-server/service/earningsServ"
	"github.com/kataras/iris/hero"
	"github.com/my-stocks-pro/api-server/service/redisServ"
	"github.com/my-stocks-pro/api-server/crud/redisCrud"
)

var ServiceApproved *approvedServ.Approved
var ServiceEarnings *earningsServ.Earnings
var ServiceRedis *redisServ.Redis


func (s *TypeServer) InitMiddleware() {

	crudApproved := approvedCrud.New(s.PSQL.NewConn())
	crudEarnings := earningsCrud.New(s.PSQL.NewConn())
	crudRedis := redisCrud.New(s.RDS.NewConn("1"))

	ServiceApproved = approvedServ.New(crudApproved)
	ServiceEarnings = earningsServ.New(crudEarnings)
	ServiceRedis = redisServ.New(crudRedis)

	hero.Register(ServiceApproved)
	hero.Register(ServiceEarnings)
	hero.Register(ServiceRedis)

}