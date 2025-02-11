package configuration

import "fmt"

type WebAPICSRFConfig struct {
	CookieName string
	CookiePath string
	//in second
	CookieMaxAge   int
	CookieSecure   bool
	CookieHTTPOnly bool
	Skip bool
}

func (apicsrf WebAPICSRFConfig) String() string {
	return fmt.Sprintf("CookieName: %s, CookiePath: %s, CookieMaxAge: %d, CookieSecure: %t, CookieHTTPOnly: %t, Skip: %t",
		apicsrf.CookieName,
		apicsrf.CookiePath ,
		apicsrf.CookieMaxAge,
		apicsrf.CookieSecure, apicsrf.CookieHTTPOnly, apicsrf.Skip)
}
