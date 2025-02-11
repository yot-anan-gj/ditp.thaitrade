package server_middlewares

import (
	"errors"
	"github.com/labstack/echo"
	"gitlab.com/ditp.thaitrade/enginex/session"
	"gitlab.com/ditp.thaitrade/enginex/util/uuid"
)

var (
	ErrUUIDGeneratorServerReq = errors.New("uuid generator middleware is require web server")
)

func UUIDGenerator(stores session.Stores) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for sessionName := range stores {
				session := session.GET(sessionName, c)
				if session != nil {
					v := session.Get("uuid")
					id, ok := v.(string)
					if !ok {
						//regenerate uuid
						id = uuid.UUIDv4()
						//save to session
						session.Set("uuid", id)
						session.Save()
					}else{
						//for retain session
						session.Set("uuid", id)
						session.Save()
					}
				}
			}
			return next(c)
		}
	}
}
