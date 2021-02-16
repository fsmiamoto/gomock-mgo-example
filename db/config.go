package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	// Uri for dialing
	Uri() string
	//Name of the database
	Name() string
}

type config struct {
	user     string
	password string
	host     string
	port     int
	name     string
	uri      string
}

func NewConfig() Config {
	c := &config{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		host:     os.Getenv("DB_HOST"),
		name:     os.Getenv("DB_NAME"),
	}

	var err error
	c.port, err = strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatalln(err)
	}

	c.uri = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", c.user, c.password, c.host, c.port, c.name)
	return c
}

func (c *config) Uri() string {
	return c.uri
}

func (c *config) Name() string {
	return c.name
}
