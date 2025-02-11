package database

import (
	"database/sql"
	"github.com/pkg/errors"
)

//collect database connection
type Connections map[string]*sql.DB

var (

	ErrDBURLRequire = errors.New("database url is require")
	ErrDBUserRequire = errors.New("database user is require")
	ErrDBPasswordRequire = errors.New("database password is require")
	ErrDBNameRequire = errors.New("database name is require")
)
