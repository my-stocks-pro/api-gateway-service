package approvedPSQLServ

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/crud/approvedCrud"
	"fmt"
	"time"
	"github.com/my-stocks-pro/api-server/models"
	//"github.com/my-stocks-pro/api-server/utils"
	//"github.com/dyninc/qstring"
)

type ImageFormatType struct {
	DisplayName  string `json:"display_name"`
	DPI          int    `json:"dpi"`
	FileSize     int    `json:"file_size"`
	Format       string `json:"format"`
	Height       int    `json:"height"`
	IsLicensable bool   `json:"is_licensable"`
	Width        int    `json:"width"`
}

type ImageLinksType struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type DataImageType struct {
	ID        string  `json:"id"`
	AddedDate string  `json:"added_date"`
	Aspect    float64 `json:"aspect"`
	Assets struct {
		SmallJpg      ImageFormatType `json:"small_jpg"`
		MediumJpg     ImageFormatType `json:"medium_jpg"`
		HugeJpg       ImageFormatType `json:"huge_jpg"`
		SupersizeJpg  ImageFormatType `json:"supersize_jpg"`
		HugeTiff      ImageFormatType `json:"huge_tiff"`
		SupersizeTiff ImageFormatType `json:"supersize_tiff"`
		Preview       ImageLinksType  `json:"preview"`
		SmallThumb    ImageLinksType  `json:"small_thumb"`
		LargeThumb    ImageLinksType  `json:"large_thumb"`
		HugeThumb     ImageLinksType  `json:"huge_thumb"`
	}
	Categories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	Contributor struct {
		ID string `json:"id"`
	}
	Description        string   `json:"description"`
	ImageType          string   `json:"image_type"`
	IsAdult            bool     `json:"is_adult"`
	IsIllustration     bool     `json:"is_illustration"`
	HasPropertyRelease bool     `json:"has_property_release"`
	Keywords           []string `json:"keywords"`
	MediaType          string   `json:"media_type"`
}

type Date struct {
	Start string `form:"start"`
	End   string `form:"end"`
}


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
	var data DataImageType

	if err := ctx.ReadJSON(&data); err != nil {
		panic(err.Error())
	}

	t, errParse := time.Parse("2006-01-02", data.AddedDate)
	if errParse != nil {
		fmt.Println(errParse)
	}

	image := models.Approve{
		Timestamp:   t.Unix(),
		IDI:         data.ID,
		AddedDate:   data.AddedDate,
		Link:        data.Assets.SmallThumb.URL,
		Description: data.Description,
	}

	m.crud.Save(image)
}

func (m *Approved) GetHistory(ctx iris.Context) {

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

