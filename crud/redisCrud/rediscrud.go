package redisCrud

import (
	"github.com/my-stocks-pro/api-server/models"
	"github.com/go-redis/redis"
	"encoding/json"
	"fmt"
)

type TypeDataForRedis struct {
	Date      string
	ListIDS   []string
}

type Crud struct {
	db *redis.Client
}

func New(connection *redis.Client) *Crud {
	return &Crud{
		db: connection,
	}
}

func (i *Crud) Save(r models.Earnings) {
	//i.db.Where("ico_id=?", finance.IcoId).Assign(finance).FirstOrCreate(&finance)
}


func (i *Crud) Get(key string) []byte {
	b, errGet := i.db.Get(key).Bytes()
	if errGet == redis.Nil {
		return nil
	}

	return b
}


func (i *Crud) GetApprovedData() TypeDataForRedis {

	dateByte := i.Get("date")
	listidsByte := i.Get("listids")

	var date string
	err1 := json.Unmarshal(dateByte, &date)
	if err1 != nil {
		fmt.Println(err1)
	}

	var listids []string
	err2 := json.Unmarshal(listidsByte, &listids)
	if err2 != nil {
		fmt.Println(err2)
	}

	res := TypeDataForRedis{
		date,
		listids,
	}

	return res
}