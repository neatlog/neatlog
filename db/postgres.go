package db

type Postgres struct {
	ConnString string
}

func (db *Postgres) Connect() error {
	return nil
}
