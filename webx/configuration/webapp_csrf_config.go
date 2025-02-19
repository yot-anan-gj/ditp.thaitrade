package configuration

import "fmt"

type WebAppCSRFConfig struct {
	CookieName string
	CookiePath string
	//in second
	CookieMaxAge   int
	CookieSecure   bool
	CookieHTTPOnly bool
	Skip bool
}

func (wcf WebAppCSRFConfig) String() string {
	return fmt.Sprintf("CookieName: %s, CookiePath: %s, CookieMaxAge: %d, CookieSecure: %t, CookieHTTPOnly: %t,  Skip: %t",
		wcf.CookieName,
		wcf.CookiePath ,
		wcf.CookieMaxAge,
		wcf.CookieSecure, wcf.CookieHTTPOnly, wcf.Skip)
}


