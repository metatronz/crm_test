package handlers

import "github.com/jmoiron/sqlx"

type Database struct {
	BillDB *sqlx.DB
	AppDB  *sqlx.DB
}
