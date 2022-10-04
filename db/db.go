package db

type Database interface {
	Connect() error
}
