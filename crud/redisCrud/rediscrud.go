package redisCrud

import (
	"github.com/my-stocks-pro/api-server/models"
	"github.com/go-redis/redis"
)

type Crud struct {
	db *redis.Client
}

func New(connection *redis.Client) *Crud {
	return &Crud{
		db: connection,
	}
}

func (i *Crud) Save(finance models.Earnings) {
	//i.db.Where("ico_id=?", finance.IcoId).Assign(finance).FirstOrCreate(&finance)
}
