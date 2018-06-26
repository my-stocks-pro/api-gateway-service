package apiserver

import (
	"github.com/my-stocks-pro/api-server/crud/approvedCrud"
	"github.com/my-stocks-pro/api-server/crud/earningsCrud"
	"github.com/my-stocks-pro/api-server/service/approvedPSQLServ"
	"github.com/my-stocks-pro/api-server/service/earningsPSQLServ"
	"github.com/kataras/iris/hero"
	"github.com/my-stocks-pro/api-server/service/earningsRDSServ"
	"github.com/my-stocks-pro/api-server/crud/redisCrud"
	"github.com/my-stocks-pro/api-server/service/approvedRDSServ"
)

var ServiceApprovedPSQL *approvedPSQLServ.Approved
var ServiceEarningsPSQL *earningsPSQLServ.Earnings
var ServiceApprovedRDS *approvedRDSServ.Approved
var ServiceEarningsRDS *earningsRDSServ.Earnings


func (s *TypeServer) InitMiddleware() {

	psqlApproved := approvedCrud.New(s.PSQL.NewConn())
	psqlEarnings := earningsCrud.New(s.PSQL.NewConn())
	rdsApproved := redisCrud.New(s.RDS.NewConn("1"))
	rdsEarnings := redisCrud.New(s.RDS.NewConn("2"))

	ServiceApprovedPSQL = approvedPSQLServ.New(psqlApproved)
	ServiceEarningsPSQL = earningsPSQLServ.New(psqlEarnings)
	ServiceApprovedRDS = approvedRDSServ.New(rdsApproved)
	ServiceEarningsRDS = earningsRDSServ.New(rdsEarnings)

	hero.Register(ServiceApprovedPSQL)
	hero.Register(ServiceEarningsPSQL)
	hero.Register(ServiceApprovedRDS)
	hero.Register(ServiceEarningsRDS)

}