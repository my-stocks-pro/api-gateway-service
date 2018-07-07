package earningsPSQLServ

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/models"
	"github.com/my-stocks-pro/api-server/crud/earningsCrud"
	"fmt"
)

type DataEarningsType struct {
	IDI      int
	Download int
	Category string
	Country  string
	City     string
}

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
	var data DataEarningsType

	if err := ctx.ReadJSON(&data); err != nil {
		panic(err.Error())
	}

	//t, errParse := time.Parse("2006-01-02", data.AddedDate)
	//if errParse != nil {
	//	fmt.Println(errParse)
	//}

	image := models.Earnings{
		//Timestamp: t.Unix(),
		IDI:       data.IDI,
		Download:  data.Download,
		Category:  data.Category,
		Country:   data.Country,
		City:      data.City,
	}

	fmt.Println(image)

	//m.crud.Save(image)

	//m.crud.Save(data)
}

func (m *Earnings) GetHistory(ctx iris.Context) {

	//date := Date{}
	//err := ctx.ReadForm(&date)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//query, errQ := qstring.MarshalString(&date)
	//if errQ != nil {
	//	fmt.Println(errQ)
	//}
	//
	//b, e := utils.NewRequest(fmt.Sprintf("%s?%s", "http://127.0.0.1:8002/history/approved", query))
	//if e != nil {
	//	fmt.Println(e)
	//}
	//fmt.Println(string(b))
}

