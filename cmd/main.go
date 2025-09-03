package main

import (
	"fmt"

	_ "github.com/tunangoo/full-time-go-dev/docs"
	"github.com/tunangoo/full-time-go-dev/internal/config"
)

// @title Full Time Go Dev API
// @version 1.0
// @description Full Time Go Dev API
// @contact.name API Support
// @contact.email support@full-time-go-dev.com
// @host localhost:7000
// @BasePath /api
// @schemes http
func main() {
	server := config.NewAPIServer(config.SvcCfg.Server.Name, config.SvcCfg.Server.Addr)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.SvcCfg.Database.User, config.SvcCfg.Database.Password, config.SvcCfg.Database.Host, config.SvcCfg.Database.Port, config.SvcCfg.Database.Database, config.SvcCfg.Database.SSLMode)
	db := config.NewPostgres(dsn)

	handler, err := wireApp(db)
	if err != nil {
		panic(err)
	}

	handler.RegisterRoutes(server.Router)

	server.Run()
}
