package main

import (
	"fmt"
	"net/http"
	"os"

	unoHttp "github.com/UnoHouse/api/http"
	"github.com/UnoHouse/api/service"
	"github.com/UnoHouse/api/sql"
	"github.com/UnoHouse/api/utils/logger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	env := os.Getenv("APP_ENV")

	db := sql.NewDB()
	s := service.NewMySqlService(db)

	handler := unoHttp.NewHandler(&s)
	router := unoHttp.NewRouter(handler)
	server := unoHttp.NewServer(router)
	logger.Log(fmt.Sprintf("Server started - listen on address %s \n", server.Addr), logger.Info)
	listen(server, env)

}

func listen(server *http.Server, env string) {
	var err error

	if env == "production" || env == "staging" {
		err = server.ListenAndServeTLS("/etc/ssl/private/unohouse.com.pl.crt", "/etc/ssl/private/unohouse.com.pl.key")
	} else {
		err = server.ListenAndServe()
	}

	if err != nil {
		logger.Log(err.Error(), logger.Error)
	}
}
