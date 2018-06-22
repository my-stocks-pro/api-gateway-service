package approvedServ

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/models"
	"github.com/my-stocks-pro/api-server/crud/approvedCrud"
)


type Approved struct {
	crud *approvedCrud.Crud
}

func New(crud *approvedCrud.Crud) *Approved {
	return &Approved{
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

func (m *Approved) PostALL(ctx iris.Context) {
	data := models.Approve{}

	if err := ctx.ReadJSON(&data); err != nil {
		panic(err.Error())
	}

	m.crud.Save(data)
}