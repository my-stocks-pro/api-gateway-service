package earningsPSQLServ

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/models"
	"github.com/my-stocks-pro/api-server/crud/earningsCrud"
	"fmt"
	"time"
	"github.com/my-stocks-pro/api-server/utils"
	"encoding/json"
)

type DataEarningsType struct {
	IDI      int    `json:"idi"`
	Date     string `json:"date"`
	Download int    `json:"download"`
	Category string `json:"category"`
	Country  string `json:"country"`
	City     string `json:"city"`
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

	t, errParse := time.Parse("2006-01-02", data.Date)
	if errParse != nil {
		fmt.Println(errParse)
	}

	image := models.Earnings{
		Timestamp: t.Unix(),
		Date:      data.Date,
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

type Date struct {
	Start string `form:"start"`
	End   string `form:"end"`
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

	date := Date{
		Start: "2018-06-01",
		End:   "2018-07-01",
	}

	b, err := json.Marshal(date)
	if b != nil {
		fmt.Println(err)
	}

	b, e := utils.NewRequest("http://127.0.0.1:8003/history/earnings", b)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(b))
}
