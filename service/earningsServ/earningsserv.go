package earningsServ

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/models"
	"github.com/my-stocks-pro/api-server/crud/earningsCrud"
	"fmt"
)

type Earnings struct {
	crud *earningsCrud.Crud
}

func New(crud *earningsCrud.Crud) *Earnings {
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
	data := models.Earnings{}

	if err := ctx.ReadJSON(&data); err != nil {
		panic(err.Error())
	}

	fmt.Println(data)

	//m.crud.Save(data)
}
