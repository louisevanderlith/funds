package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong/middle"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	/*
		credCtrl := &handles.Credits{}
			reqCtrl := &handles.Requisitions{}
			accCtrl := &handles.Accounts{}
			e.JoinBundle("/", roletype.Owner, mix.JSON, credCtrl, reqCtrl, accCtrl)
	*/

	r := mux.NewRouter()

	lst, err := middle.Whitelist(http.DefaultClient, secureUrl, "funds.accounts.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
