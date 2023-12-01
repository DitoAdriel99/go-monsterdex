package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/router"
	_ "github.com/DitoAdriel99/go-monsterdex/docs/echosimple"
)

func main() {
	e := router.New()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		Handler: e,
	}

	// Start the server
	e.Logger.Fatal(e.Start(server.Addr))

}
