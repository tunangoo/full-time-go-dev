package main

import (
	"log"

	"github.com/tunangoo/full-time-go-dev/internal/config"
)

func main() {
	log.Println(config.SvcCfg.Server.Addr)
	log.Println(config.SvcCfg.Database.Host)
	log.Println(config.SvcCfg.Database.Port)
	log.Println(config.SvcCfg.Database.User)
	log.Println(config.SvcCfg.Database.Password)
	log.Println(config.SvcCfg.Database.Database)
	log.Println(config.SvcCfg.Database.Schema)
	log.Println(config.SvcCfg.Database.SSLMode)
	log.Println(config.SvcCfg.Environment)
}
