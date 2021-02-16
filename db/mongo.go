package db

import (
	"gopkg.in/mgo.v2"
)

// MongoSession wraps both a MongoDatabase and a mgo.Session
type MongoSession struct {
	session  *mgo.Session
	database *MongoDatabase
}

func NewMongoSession(cfg Config) (Session, error) {
	session, err := mgo.Dial(cfg.Uri())
	if err != nil {
		return nil, err
	}
	return &MongoSession{session: session, database: &MongoDatabase{session.DB(cfg.Name())}}, nil
}

func (c *MongoSession) Close() {
	c.session.Close()
}

func (c *MongoSession) DB() Database {
	return c.database
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d *MongoDatabase) C(name string) Collection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

func (c *MongoCollection) Find(query interface{}) Query {
	return &MongoQuery{Query: c.Collection.Find(query)}
}

func (c *MongoCollection) FindId(id interface{}) Query {
	return &MongoQuery{Query: c.Collection.FindId(id)}
}

type MongoQuery struct {
	*mgo.Query
}
