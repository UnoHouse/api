package main

import (
	"fmt"
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
	usersDb := sql.NewUsersDB()
	s := service.NewMySqlService(db)
	us := service.NewMySqlService(usersDb)
	handler := unoHttp.NewHandler(&s, &us)
	router := unoHttp.NewRouter(handler)
	server := unoHttp.NewServer(router)
	logger.Log(fmt.Sprintf("Server started - listen on address %s \n", server.Addr), logger.Info)
	unoHttp.SetupServer(server, env)
}
