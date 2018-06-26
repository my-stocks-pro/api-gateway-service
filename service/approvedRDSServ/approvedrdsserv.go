package approvedRDSServ

import (
	"github.com/my-stocks-pro/api-server/crud/redisCrud"
	"fmt"
	"github.com/kataras/iris"
)

type TypeDataForRedis struct {
	Date      string
	ListIDS   []string
}

type Approved struct {
	crud *redisCrud.Crud
}

func New(crud *redisCrud.Crud) *Approved {
	return &Approved{
		crud: crud,
	}
}

func (m *Approved) PostALL(ctx iris.Context) {
	data := TypeDataForRedis{}

	if err := ctx.ReadJSON(&data); err != nil {
		panic(err.Error())
	}

	fmt.Println(data)

	//m.crud.Save(data)
}


func (m *Approved) GetALL(ctx iris.Context) {
	_, err := ctx.JSON(m.crud.GetApprovedData())
	if err != nil {
		fmt.Println(err)
	}
}
