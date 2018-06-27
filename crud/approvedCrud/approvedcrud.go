package approvedCrud

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

func (i *Crud) Save(approved models.Approve) {
	i.db.Where("id_i=?", approved.IDI).Assign(approved).FirstOrCreate(&approved)
	fmt.Println(approved)
}

