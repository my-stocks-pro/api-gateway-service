package earningsCrud

import (
	"github.com/jinzhu/gorm"
	"github.com/my-stocks-pro/api-server/models"
	"fmt"
)

type Crud struct {
	db *gorm.DB
}

func New(connection *gorm.DB) *Crud {
	return &Crud{
		db: connection,
	}
}

func (i *Crud) Save(earnings models.Earnings) {
	i.db.Where("id_i=?", earnings.IDI).Omit("country", "city").Assign(earnings).FirstOrCreate(&earnings)
	fmt.Println(earnings)
}