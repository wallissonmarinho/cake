package main

import (
	"fmt"
	"log"

	"github.com/wallissonmarinho/cake/bootstrap"
	"github.com/wallissonmarinho/cake/config"
)

func main() {
	app := bootstrap.NewApplication()
	cfg := config.Get()
	addr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)

	if cfg.DevMode {
		log.Fatal(app.Listen(addr))
	} else {
		log.Fatal(app.ListenTLS(addr, cfg.CertFilePath, cfg.KeyFilePath))
	}
}
