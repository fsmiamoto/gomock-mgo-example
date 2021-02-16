package db

// Session is an interface for a database session
type Session interface {
	Close()
	DB() Database
}

// Collection is an interface to access to the collection struct.
type Collection interface {
	Find(query interface{}) Query
	FindId(id interface{}) Query
	Count() (n int, err error)
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	RemoveId(id interface{}) error
	Update(selector interface{}, update interface{}) error
	UpdateId(id interface{}, update interface{}) error
}

// Database is an interface for fetching a Collection
type Database interface {
	C(name string) Collection
}

type Query interface {
	One(result interface{}) error
	All(result interface{}) error
	Count() (n int, err error)
}
