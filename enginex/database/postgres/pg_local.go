package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/ditp.thaitrade/enginex/database"
	"gitlab.com/ditp.thaitrade/enginex/util/stringutil"
)

func Open(
	url string,
	user string,
	password string,
	dbName string) (*sql.DB, error){
		if stringutil.IsEmptyString(url) {
			return nil, database.ErrDBURLRequire
		}
		if stringutil.IsEmptyString(user){
			return nil, database.ErrDBUserRequire
		}
		if stringutil.IsEmptyString(password){
			return nil, database.ErrDBPasswordRequire
		}
		if stringutil.IsEmptyString(dbName){
			return nil, database.ErrDBNameRequire
		}

	urlStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user,password,url,dbName)
	db, err := sql.Open("postgres", urlStr)
	if err != nil {
		return nil, err
	}
	return db, nil

}
