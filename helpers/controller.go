package helpers

import (
	"github.com/gorilla/context"
	"github.com/gorilla/csrf"
	"net/http"
	"strconv"
)

//DefaultData returns common to all pages template data
func DefaultData(r *http.Request) map[string]interface{} {
	return map[string]interface{}{
		"ActiveUser":     context.Get(r, "user"),           //signed in models.User
		"Active":         "",                               //active uri shortening for menu item highlight
		"Title":          "",                               //page title:w
		"SignupEnabled":  context.Get(r, "signup_enabled"), //signup route is enabled (otherwise everyone can signup ;)
		csrf.TemplateTag: csrf.TemplateField(r),
	}
}

//ErrorData returns template data for error
func ErrorData(err error) map[string]interface{} {
	return map[string]interface{}{
		"Title": err.Error(),
		"Error": err.Error(),
	}
}

//Atoi64 converts string to int64, returns 0 if error
func Atoi64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

//Atob converts string to bool
func Atob(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}
