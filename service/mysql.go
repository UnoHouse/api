package service

import (
	mySql "github.com/UnoHouse/api/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SqlService interface {
}

type SqlServiceImpl struct {
	db mySql.DB
}

func NewMySqlService(db mySql.DB) SqlServiceImpl {
	return SqlServiceImpl{db}
}
