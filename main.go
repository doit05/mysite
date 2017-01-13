package main

import (
	"github.com/astaxie/beego"
	_ "mysite/routers"
	_ "mysite/utils"
	// "github.com/astaxie/beego/plugins/auth"
)

func main() {
	// authenticate every request
	// authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
	// beego.InsertFilter("*", beego.BeforeRouter,authPlugin)
	beego.Run()
}

// func SecretAuth(username, password string) bool {
// 	return username == "astaxie" && password == "helloBeego"
// }
