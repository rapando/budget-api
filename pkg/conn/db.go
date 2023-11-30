package conn

import "database/sql"

func DBConn(uri string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
