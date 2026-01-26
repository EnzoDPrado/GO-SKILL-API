package database

import (
	"fmt"
	"log"
	"rest-api/internal/domain"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Connection struct {
	Db       *gorm.DB
	host     string
	port     int64
	user     string
	password string
	dbname   string
}

func NewConnection(
	host string,
	port int64,
	user string,
	password string,
	dbname string,
) *Connection {
	return &Connection{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

func (c *Connection) Connect() (*gorm.DB, error) {
	var err error

	db, err := gorm.Open(postgres.Open(c.makeDsn()), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	c.Db = db

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("Falha ao rodar migrations:", err)
	}

	return db, nil
}

func (c *Connection) makeDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.host, c.port, c.user, c.password, c.dbname)

}
