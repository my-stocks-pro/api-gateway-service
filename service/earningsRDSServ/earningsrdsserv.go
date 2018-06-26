package earningsRDSServ

import (
	"github.com/kataras/iris"
	"fmt"
	"github.com/my-stocks-pro/api-server/crud/redisCrud"
)

type TypeDataForRedis struct {
	Date      string
	ListIDS   []string
}

type Earnings struct {
	crud *redisCrud.Crud
}

func New(crud *redisCrud.Crud) *Earnings {
	return &Earnings{
		crud: crud,
	}
}

//func (p *ProjectFinance) GetAll(ctx iris.Context) {
//	ctx.JSON(p.crud.Select())
//}
//
//func (p *ProjectFinance) GetById(ctx iris.Context, id string) {
//	ctx.JSON(p.crud.Find(id))
//}

func (m *Earnings) PostALL(ctx iris.Context) {
	data := TypeDataForRedis{}

	if err := ctx.ReadJSON(&data); err != nil {
		panic(err.Error())
	}

	//m.crud.Save(data)
}


func (m *Earnings) GetALL(ctx iris.Context) {
	_, err := ctx.JSON(m.crud.GetApprovedData())
	if err != nil {
		fmt.Println(err)
	}
}