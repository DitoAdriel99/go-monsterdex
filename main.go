package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/router"
	_ "github.com/DitoAdriel99/go-monsterdex/docs/echosimple"
)

// @title Monsterdex Enpoints
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://35.188.107.108
// @BasePath /
// @schemes http
func main() {
	e := router.New()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		Handler: e,
	}

	// Start the server
	e.Logger.Fatal(e.Start(server.Addr))

}
