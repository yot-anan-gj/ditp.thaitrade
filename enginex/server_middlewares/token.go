package server_middlewares

import (
	"github.com/labstack/echo"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/server_constant"
)

func UserProfileContext(devMode bool, sellerCode string, accountID string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if devMode {
				c.Set(server_constant.UserAccountID, accountID)
				c.Set(server_constant.SellerCode, sellerCode)
			}

			return next(c)
		}
	}
}
