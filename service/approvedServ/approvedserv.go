package approvedServ

import (
	"github.com/kataras/iris"
	"github.com/my-stocks-pro/api-server/crud/approvedCrud"
	"fmt"
)

type ImageFormatType struct {
	Display_name  string
	DPI           int
	File_size     int
	Format        string
	Height        int
	Is_licensable bool
	Width         int
}

type ImageLinksType struct {
	Height int
	URL    string
	Width  int
}

type DataImageType struct {
	ID         string
	Added_date string
	Aspect     float64
	Assets struct {
		Small_jpg      ImageFormatType
		Medium_jpg     ImageFormatType
		Huge_jpg       ImageFormatType
		Supersize_jpg  ImageFormatType
		Huge_tiff      ImageFormatType
		Supersize_tiff ImageFormatType
		Preview        ImageLinksType
		Small_thumb    ImageLinksType
		Large_thumb    ImageLinksType
		Huge_thumb     ImageLinksType
	}
	Categories []struct {
		ID   string
		Name string
	}
	Contributor struct {
		ID string
	}
	Description          string
	Image_type           string
	Is_adult             bool
	Is_illustration      bool
	Has_property_release bool
	Keywords             []string
	media_type           string
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
	var data interface{}

	if err := ctx.ReadJSON(&data); err != nil {
		panic(err.Error())
	}

	fmt.Println(data)

	//m.crud.Save(data)
}