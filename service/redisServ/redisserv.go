package redisServ

import (
	"github.com/kataras/iris"
	"fmt"
	"github.com/my-stocks-pro/api-server/crud/redisCrud"
)

type TypeDataForRedis struct {
	Date      string
	ListIDS   []string
}

type Redis struct {
	crud *redisCrud.Crud
}

func New(crud *redisCrud.Crud) *Redis {
	return &Redis{
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

func (m *Redis) PostALL(ctx iris.Context) {
	data := TypeDataForRedis{}

	if err := ctx.ReadJSON(&data); err != nil {
		panic(err.Error())
	}

	fmt.Println(data)

	//m.crud.Save(data)
}