package earningsCrud

import (
	"github.com/jinzhu/gorm"
	"github.com/my-stocks-pro/api-server/models"
)

type Crud struct {
	db *gorm.DB
}

func New(connection *gorm.DB) *Crud {
	return &Crud{
		db: connection,
	}
}

func (i *Crud) Save(finance models.Earnings) {
	//i.db.Where("ico_id=?", finance.IcoId).Assign(finance).FirstOrCreate(&finance)
}