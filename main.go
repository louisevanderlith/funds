package main

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/funds/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	_ "github.com/louisevanderlith/mango/funds"
)

func main() {
	mode := os.Getenv("RUNMODE")

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
