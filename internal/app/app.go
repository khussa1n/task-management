package app

import (
	"github.com/khussa1n/task-management/internal/config"
	"github.com/khussa1n/task-management/internal/handler"
	"github.com/khussa1n/task-management/internal/repository/pgrepo"
	"github.com/khussa1n/task-management/internal/service"
	"github.com/khussa1n/task-management/pkg/httpserver"
	"github.com/khussa1n/task-management/pkg/jwttoken"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	db, err := pgrepo.New(
		pgrepo.WithHost(cfg.DB.Host),
		pgrepo.WithPort(cfg.DB.Port),
		pgrepo.WithDBName(cfg.DB.DBName),
		pgrepo.WithUsername(cfg.DB.Username),
		pgrepo.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	log.Println("connection success")

	migration := pgrepo.NewMigrate(cfg)
	err = migration.MigrateToVersion(cfg.DB.MigrationVersion)
	if err != nil {
		log.Printf("from migration")
		return err
	}
	log.Println("migration success")

	token := jwttoken.New(cfg.Token.SecretKey)
	srvs := service.New(db, token, cfg)
	hndlr := handler.New(srvs)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
