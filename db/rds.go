package db

import (
	"os"
	"strconv"
	"github.com/go-redis/redis"
	"fmt"
)

type TypeRedis struct {
	RDSHOST string
	RDSPORT string
}

func (r *TypeRedis) NewConn(name string) *redis.Client {
	dbname, err := strconv.Atoi(name)
	if err != nil {
		panic(err)
	}
	connection := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.RDSHOST, r.RDSPORT),
		Password: "", // no password set
		DB:       dbname,
	})

	_, errConn := connection.Ping().Result()
	if errConn != nil {
		panic(errConn)
	}

	return connection
}


func NewRDS() *TypeRedis {
	return &TypeRedis{
		os.Getenv("RDSHOST"),
		os.Getenv("RDSPORT"),
	}
}