package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/funds/controllers/account"
	"github.com/louisevanderlith/funds/controllers/credit"
	"github.com/louisevanderlith/funds/controllers/requisition"
	"github.com/louisevanderlith/funds/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()

	accounts := r.Group("/account")
	// accounts.POST("", account.Create)
	// accounts.PUT("/:key", account.Update)
	// accounts.DELETE("/:key", account.Delete)

	r.GET("/accounts", account.Get)
	// r.GET("/accounts/:pagesize/*hash", account.Search)

	credits := r.Group("/credit")
	// credits.POST("", credit.Create)
	// credits.PUT("/:key", credit.Update)
	// credits.DELETE("/:key", credit.Delete)

	r.GET("/credits", credit.Get)
	// r.GET("/credits/:pagesize/*hash", credit.Search)

	requisitions := r.Group("/requisition")
	// requisitions.POST("", requisition.Create)
	// requisitions.PUT("/:key", requisition.Update)
	// requisitions.DELETE("/:key", requisition.Delete)

	r.GET("/requisitions", requisition.Get)
	// r.GET("/requisitions/:pagesize/*hash", requisition.Search)

	err := r.Run(":8092")

	if err != nil {
		panic(err)
	}
}

// func main() {
// 	keyPath := os.Getenv("KEYPATH")
// 	pubName := os.Getenv("PUBLICKEY")
// 	host := os.Getenv("HOST")
// 	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
// 	appName := os.Getenv("APPNAME")
// 	pubPath := path.Join(keyPath, pubName)

// 	// Register with router
// 	srv := bodies.NewService(appName, "", pubPath, host, httpport, servicetype.API)

// 	routr, err := do.GetServiceURL("", "Router.API", false)

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = srv.Register(routr)

// 	if err != nil {
// 		panic(err)
// 	}

// 	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, "none"))
// 	routers.Setup(poxy)
// 	poxy.EnableCORS(host)

// 	core.CreateContext()
// 	defer core.Shutdown()

// 	err = droxolite.Boot(poxy)

// 	if err != nil {
// 		panic(err)
// 	}
// }
