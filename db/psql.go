package db

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type TypePSQL struct {
	PGHOST string
	PGPORT string
	PGNAME string
	PGUSER string
	PGPASS string
}

func (p *TypePSQL) MakeMigrations(connection *gorm.DB) {
	migrate := os.Getenv("MIGRATE")

	if migrate == "1" {
		fmt.Println("Migrate")

		//TODO: ADD CONSTRAINTS
		connection.AutoMigrate()

		fmt.Println("Migrations done")
	}
}

func (p *TypePSQL) NewConn() *gorm.DB  {
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		p.PGHOST, p.PGPORT, p.PGNAME, p.PGUSER, p.PGPASS)

	connection, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = connection.DB().Ping()
	if err != nil {
		panic(err)
	}

	//p.MakeMigrations(connection)

	return connection
}

func NewPSQL() *TypePSQL {
	return &TypePSQL{
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGNAME"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASS"),
	}
}
