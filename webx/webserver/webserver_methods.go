package webserver

import (
	"github.com/labstack/echo"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/database"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/session"
)

func (ws *WebServer) Engine() *echo.Echo {
	return ws.engine
}

func (ws *WebServer) DBConnections() database.Connections {
	return ws.dbConnections

}

func (ws *WebServer) SessionStores() session.Stores {
	return ws.sessionStores
}
